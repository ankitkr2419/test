package service

import (
	"context"
	"encoding/json"
	"mylab/cpagent/config"
	"mylab/cpagent/db"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

type WellResult struct {
	ID           uuid.UUID          `json:"id"`
	Position     int32              `json:"position" validate:"required"`
	ExperimentID uuid.UUID          `json:"experiment_id" validate:"required"`
	SampleID     uuid.UUID          `json:"sample_id" validate:"required"`
	Task         string             `json:"task" validate:"required"`
	ColorCode    string             `json:"color_code"`
	Targets      []WellTargetResult `json:"targets" validate:"required"`
	SampleName   string             `db:"sample_name" json:"sample_name"`
}
type WellTargetResult struct {
	ExperimentID uuid.UUID `json:"experiment_id"`
	WellPosition int32     `json:"well_position"`
	TargetID     uuid.UUID `json:"target_id" validate:"required"`
	TargetName   string    `json:"target_name"`
	CT           string    `json:"ct"`
	Selected     bool      `json:"selected"`
	Threshold    float32   `json:"threshold"`
}

func listWellsHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		expID, err := parseUUID(vars["experiment_id"])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		wells, err := deps.Store.ListWells(req.Context(), expID)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error fetching data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		if len(wells) > 0 {

			welltargets, err := deps.Store.ListWellTargets(req.Context(), expID)
			if err != nil {
				logger.WithField("err", err.Error()).Error("Error fetching data")
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}

			for _, t := range welltargets {
				for i, w := range wells {
					if w.Position == t.WellPosition {
						wells[i].Targets = append(w.Targets, t)
					}
				}
			}
		}
		wellResult, err := makeWellResultData(deps, wells, expID)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error making wells data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		respBytes, err := json.Marshal(wellResult)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling Wells data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(respBytes)
	})
}
func makeWellResultData(deps Dependencies, wells []db.Well, expID uuid.UUID) (wellResult []WellResult, err error) {
	for _, w := range wells {

		var wt []WellTargetResult
		for _, t := range w.Targets {
			thresholdR, err := deps.Store.GetTargetThreshold(context.Background(), expID, t.TargetID)
			logger.Infoln(thresholdR)
			if err != nil {
				logger.WithField("err", err.Error()).Error("Error getting threshold data")
				return nil, err
			}
			wtr := WellTargetResult{
				ExperimentID: t.ExperimentID,
				WellPosition: t.WellPosition,
				TargetID:     t.TargetID,
				TargetName:   t.TargetName,
				CT:           t.CT,
				Selected:     t.Selected,
				Threshold:    scaleThreshold(thresholdR.Threshold),
			}
			wt = append(wt, wtr)
		}

		wr := WellResult{
			ID:           w.ID,
			Position:     w.Position,
			ExperimentID: w.ExperimentID,
			SampleID:     w.SampleID,
			Task:         w.Task,
			ColorCode:    w.ColorCode,
			Targets:      wt,
			SampleName:   w.SampleName,
		}
		wellResult = append(wellResult, wr)
	}
	return
}
func upsertWellHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		expID, err := parseUUID(vars["experiment_id"])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		var wc db.WellConfig
		err = json.NewDecoder(req.Body).Decode(&wc)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			logger.WithField("err", err.Error()).Error("Error while decoding Well data")
			return
		}

		valid, respBytes := validate(wc)
		if !valid {
			responseBadRequest(rw, respBytes)
			return
		}

		valid, message := ValidateDyeTarget(wc, deps)
		if !valid {
			logger.Errorln("INVALID REQUEST", message)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: message})
			return
		}

		// create sample if sample_id not present
		if !isvalidID(wc.Sample.ID) {
			wc.Sample, err = deps.Store.CreateSample(req.Context(), wc.Sample)
		}

		// create wells
		var wells []db.Well
		for _, p := range wc.Position {
			w := db.Well{
				Position:     p,
				ExperimentID: expID,
				SampleID:     wc.Sample.ID,
				Task:         wc.Task,
				ColorCode:    green, //initially all wells will be green
			}
			wells = append(wells, w)
		}

		var createdWell []db.Well
		createdWell, err = deps.Store.UpsertWells(req.Context(), wells, expID)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error upsert wells")
			return
		}

		// create well targets
		var targets []db.WellTarget
		for w := 0; w < len(createdWell); w++ {
			for t := 0; t < len(wc.Targets); t++ {
				t := db.WellTarget{
					WellPosition: createdWell[w].Position,
					TargetID:     wc.Targets[t],
				}
				targets = append(targets, t)
			}
		}

		createdTargets, err := deps.Store.UpsertWellTargets(req.Context(), targets, expID, true)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error upsert wells")
			return
		}

		for _, t := range createdTargets {
			for i, w := range createdWell {
				if w.Position == t.WellPosition {
					createdWell[i].Targets = append(createdWell[i].Targets, t)
				}
			}
		}
		respBytes, err = json.Marshal(createdWell)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling wells data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusCreated)
		rw.Write(respBytes)
		rw.Header().Add("Content-Type", "application/json")
	})
}

func showWellHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		id, err := parseUUID(vars["id"])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		var latestT db.Well

		latestT, err = deps.Store.ShowWell(req.Context(), id)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error show Well")
			return
		}

		latestT.Targets, err = deps.Store.GetWellTarget(req.Context(), latestT.Position, latestT.ExperimentID)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error show Well")
			return
		}

		respBytes, err := json.Marshal(latestT)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling Well data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusOK)
		rw.Write(respBytes)
		rw.Header().Add("Content-Type", "application/json")
	})
}

func deleteWellHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		err = deps.Store.DeleteWell(req.Context(), id)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error while deleting Well")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusOK)
		rw.Header().Add("Content-Type", "application/json")
		rw.Write([]byte(`{"msg":"Well deleted successfully"}`))
	})
}

func listActiveWellsHandler() http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		activeWells := config.ActiveWells("activeWells")

		if len(activeWells) == 0 {
			logger.WithField("err", "active wells not set in config").Error("Error marshaling Wells data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		respBytes, err := json.Marshal(activeWells)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling active Wells")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(respBytes)
	})
}
