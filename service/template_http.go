package service

import (
	"mylab/cpagent/tec"
	"encoding/json"
	"mylab/cpagent/db"
	"mylab/cpagent/config"
	"net/http"
	"context"
	"math"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func listTemplateHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		t, err := deps.Store.ListTemplates(req.Context())
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error fetching data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		respBytes, err := json.Marshal(t)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling templates data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(respBytes)
	})
}

// @Title createTemplateHandler
// @Description Create createTemplateHandler
// @Router /template [post]
// @Accept  json
// @Success 200 {object}
// @Failure 400 {object}
func createTemplateHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var t db.Template
		err := json.NewDecoder(req.Body).Decode(&t)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			logger.WithField("err", err.Error()).Error("Error while decoding template data")
			return
		}

		valid, respBytes := validate(t)
		if !valid {
			responseBadRequest(rw, respBytes)
			return
		}

		var createdTemp db.Template
		createdTemp, err = deps.Store.CreateTemplate(req.Context(), t)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error create target")
			return
		}

		// Initialize stages 1 holding & 1 cycle stage

		stgs := []db.Stage{
			{
				Type:        hold,
				TemplateID:  createdTemp.ID,
				RepeatCount: 0,
			},
			{
				Type:        cycle,
				TemplateID:  createdTemp.ID,
				RepeatCount: 0,
			},
		}

		createdTemp.Stages, err = deps.Store.CreateStages(req.Context(), stgs)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error create target")
			return
		}

		respBytes, err = json.Marshal(createdTemp)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling targets data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusCreated)
		rw.Write(respBytes)
		rw.Header().Add("Content-Type", "application/json")
	})
}

func updateTemplateHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		var t db.Template

		err = json.NewDecoder(req.Body).Decode(&t)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			logger.WithField("err", err.Error()).Error("Error while decoding template data")
			return
		}

		valid, respBytes := validate(t)
		if !valid {
			responseBadRequest(rw, respBytes)
			return
		}

		t.ID = id

		err = deps.Store.UpdateTemplate(req.Context(), t)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error update template")
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: err.Error()} )
			return
		}

		rw.WriteHeader(http.StatusOK)
		rw.Header().Add("Content-Type", "application/json")
		rw.Write([]byte(`{"msg":"template updated successfully"}`))
	})
}

func showTemplateHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		id, err := parseUUID(vars["id"])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		var latestT db.Template

		latestT, err = deps.Store.ShowTemplate(req.Context(), id)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error show template")
			return
		}

		respBytes, err := json.Marshal(latestT)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling template data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusOK)
		rw.Write(respBytes)
		rw.Header().Add("Content-Type", "application/json")
	})
}

func deleteTemplateHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		err = deps.Store.DeleteTemplate(req.Context(), id)
		if err != nil {
			if err.Error() == "Violates foreign key constraint" {
				// cannot delete template as it is used in experiments
				rw.WriteHeader(http.StatusForbidden)
				rw.Header().Add("Content-Type", "application/json")
				rw.Write([]byte(`{"err":"This template is used in experiments so cannot be deleted"}`))
				return

			}
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusOK)
		rw.Header().Add("Content-Type", "application/json")
		rw.Write([]byte(`{"msg":"template deleted successfully"}`))
	})
}

func publishTemplateHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		id, err := parseUUID(vars["id"])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		t, err := deps.Store.ListTemplateTargets(req.Context(), id)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error fetching data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		ss, err := deps.Store.ListStageSteps(req.Context(), id)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error fetching data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		// validate template
		errorResp, valid := db.ValidateTemplate(t, ss)

		if !valid {
			respBytes, err := json.Marshal(errorResp)
			if err != nil {
				logger.WithField("err", err.Error()).Error("Error marshaling template data")
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(respBytes)
			rw.Header().Add("Content-Type", "application/json")
			return
		}

		err = deps.Store.PublishTemplate(req.Context(), id)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error fetching data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp, err := deps.Store.CheckIfICTargetAdded(req.Context(), id)
		if err != nil {
			if err.Error() == "Record Not Found" {
				// no Internal control added
				respBytes, err := json.Marshal(resp)
				if err != nil {
					logger.WithField("err", err.Error()).Error("Error marshaling template data")
					rw.WriteHeader(http.StatusInternalServerError)
					return
				}
				rw.Header().Add("Content-Type", "application/json")
				rw.WriteHeader(http.StatusAccepted)
				rw.Write(respBytes)
				return
			}
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(`{"msg":"template published successfully"}`))

		rw.Header().Add("Content-Type", "application/json")
		return
	})
}

func listPublishedTemplateHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		t, err := deps.Store.ListPublishedTemplates(req.Context())
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error fetching data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		respBytes, err := json.Marshal(t)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling templates data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(respBytes)
	})
}


// ALGORITHM
// 1. Get Stage ID from Step
// 2. Fetch Template ID from this Stage ID
// 3. Fetch All Stages and Steps from Template ID
// 4. Iterate over the stages and steps and calculate time accordingly
// 5. Update time in DB
func updateEstimatedTime(ctx context.Context, s db.Storer, step db.Step) (err error) {

	var currentTemp, estimatedTime, roomTemp, temp float64

	// Get stage
	stage, err := s.ShowStage(ctx, step.StageID)
	if err != nil{
		logger.WithField("err", err.Error()).Error("Error fetching Stage Data")
		return
	}

	ss, err := s.ListStageSteps(ctx, stage.TemplateID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error fetching Stage Steps Data")
		return
	}
	plcStage := makePLCStage(ss)
	logger.WithField("Calculating Time for :", plcStage).Infoln("Calculating for every Step")

	roomTemp = config.GetRoomTemp()
	currentTemp = roomTemp
	logger.Infoln(currentTemp)

	// Calculate Homing Time as its included in Experiment Time
	estimatedTime += float64(config.GetHomingTime())
	logger.Infoln("Estimated Time for Homing RTPCR: ", config.GetHomingTime())

	// Calculate Lid Temp Time
	// NOTE: This is where most variance exists for estimated time
	// TODO: Handle this in a better and accurate way
	// here 0.5 is the rate of heating/ cooling per sec 
	estimatedTime += math.Abs(float64(plcStage.IdealLidTemp) - roomTemp)/ 0.5
	logger.Infoln("Estimated Time for Lid Temp Reaching: ", math.Abs(float64(plcStage.IdealLidTemp) - roomTemp)/ 0.5)


	// Calculate Hold Stage Time
	for _, ho := range plcStage.Holding {
		estimatedTime += math.Abs(currentTemp - float64(ho.TargetTemp))/float64(ho.RampUpTemp)
		estimatedTime += float64(ho.HoldTime)
		currentTemp = float64(ho.TargetTemp)
	}
	logger.Infoln("Estimated Time after Holding Stage: ", estimatedTime)

	// Calculate Cycle Stage Time
	temp = 0
	for _, co := range plcStage.Cycle {
		temp += math.Abs(currentTemp - float64(co.TargetTemp))/float64(co.RampUpTemp)
		temp += float64(co.HoldTime)
		currentTemp = float64(co.TargetTemp)
	}

	estimatedTime += temp * float64(plcStage.CycleCount)
	logger.Infoln("Estimated Time after Cycling Stage: ", estimatedTime)


	// Last step to go back to Room Temp
	estimatedTime += math.Abs(currentTemp - roomTemp)/tec.RoomTempRamp

	err = s.UpdateEstimatedTime(ctx, stage.TemplateID, int64(estimatedTime))
	return
}