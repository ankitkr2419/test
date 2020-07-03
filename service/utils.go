package service

import (
	"encoding/json"
	"fmt"
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"net/http"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
)

var (
	plcStage     plc.Stage
	experimentID uuid.UUID   //set experimentID which is currently running
	green        = "#3FC13A" // All CT values for the well are below threshold,
	red          = "#F06666" //Even a single value crosses threshold for target
	orange       = "#F3811F" // If the CT values are close to threshold (delta)
)

func validate(i interface{}) (valid bool, respBytes []byte) {

	fieldErrors := make(map[string]string)

	v := validator.New()
	err := v.Struct(i)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			fieldErrors[e.Namespace()] = e.Tag()

			logger.WithFields(logger.Fields{
				"field": e.Namespace(),
				"tag":   e.Tag(),
				"error": "invalid value",
			}).Error("Validation Error")
		}

		respBytes, err = json.Marshal(fieldErrors)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling data")
			return
		}
		return
	}

	valid = true
	return
}

func parseUUID(s string) (validUUID uuid.UUID, err error) {

	validUUID, err = uuid.Parse(s)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error id key invalid")
		return
	}

	return
}

func isvalidID(id uuid.UUID) bool {
	// 00000000-0000-0000-0000-000000000000 is default value of uuid
	if id.String() == "00000000-0000-0000-0000-000000000000" {
		return false
	}
	return true
}

func responseBadRequest(rw http.ResponseWriter, respBytes []byte) {
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(http.StatusBadRequest)
	rw.Write(respBytes)
	return
}

// makePLCStage return plc.Stage from stagesteps
func makePLCStage(ss []db.StageStep) plc.Stage {
	// var plcStage plc.Stage

	for _, s := range ss {
		var step plc.Step
		step.RampUpTemp = s.RampRate
		step.TargetTemp = s.TargetTemperature
		step.HoldTime = s.HoldTime

		switch s.Type {
		case "hold":
			plcStage.Holding = append(plcStage.Holding, step)
		case "cycle":
			plcStage.Cycle = append(plcStage.Cycle, step)
			plcStage.CycleCount = s.RepeatCount
		default:
			logger.WithField("Unknown stage type", s.Type).Error("Error in configuring plc stages")
			return plcStage
		}
	}
	return plcStage
}

// makeResult return result from plc.scan
func makeResult(aw []int32, scan plc.Scan, td []db.TargetDetails, experimentID uuid.UUID) (result []db.Result) {
	for _, w := range aw {
		var r db.Result
		r.WellPosition = w
		r.ExperimentID = experimentID
		r.Cycle = scan.Cycle
		for _, t := range td {
			t.DyePosition = t.DyePosition - 1   // -1 dye position starts with 1
			r.TargetID = t.TargetID
			r.FValue = scan.Wells[w][t.DyePosition]
//			fmt.Println("FValue",scan.Wells[w][t.DyePosition])
			result = append(result, r)
		}
	}
	return
}

func analyseResult(activeWells []int32,td []db.TargetDetails,result []db.Result,TotalCycles uint16) (finalResult []db.FinalResult) {
	fmt.Printf("IN analyseResult: \n")
	for _, aw := range activeWells {
		var wellResult db.FinalResult
		wellResult.WellPosition = aw

			for _, t := range td {
				for _, r := range result {
				if r.WellPosition == wellResult.WellPosition && r.TargetID == t.TargetID {
					wellResult.ExperimentID = r.ExperimentID
					wellResult.TargetID = r.TargetID
					wellResult.Threshold = r.Threshold
					wellResult.TotalCycles = TotalCycles
					wellResult.Cycle = append(wellResult.Cycle,r.Cycle)
					wellResult.FValue = append(wellResult.FValue,r.FValue)
			}

			}
	finalResult = append(finalResult, wellResult)
	fmt.Printf("wellResult : %+v \n", wellResult)
break
	}

	}
	fmt.Printf("finalResult : %+v \n", finalResult)
	return
}
