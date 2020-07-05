package service

import (
	"encoding/json"
	"mylab/cpagent/config"
	"mylab/cpagent/db"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

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
			wellID := make([]uuid.UUID, 0)
			for _, w := range wells {
				wellID = append(wellID, w.ID)
			}

			welltargets, err := deps.Store.ListWellTargets(req.Context(), wellID)
			if err != nil {
				logger.WithField("err", err.Error()).Error("Error fetching data")
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}

			for i, w := range wells {
				for _, t := range welltargets {
					if w.ID == t.WellID {
						wells[i].Targets = append(w.Targets, t)
					}
				}
			}
		}
		respBytes, err := json.Marshal(wells)
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
					WellID:   createdWell[w].ID,
					TargetID: wc.Targets[t],
				}
				targets = append(targets, t)
				createdWell[w].Targets = append(createdWell[w].Targets, t)
			}
		}

		err = deps.Store.UpsertWellTargets(req.Context(), targets)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error upsert wells")
			return
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

		latestT.Targets, err = deps.Store.GetWellTarget(req.Context(), id)
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
