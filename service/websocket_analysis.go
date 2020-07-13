package service

import (
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"strconv"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

func setExperimentValues(aw []int32, TargetDetails []db.TargetDetails, ExperimentID uuid.UUID, stage plc.Stage) {
	experimentValues = experimentResultValues{
		experimentID: ExperimentID,
		activeWells:  aw,
		targets:      TargetDetails,
		plcStage:     stage,
	}

}

// makePLCStage return plc.Stage from stagesteps
func makePLCStage(ss []db.StageStep) plc.Stage {
	var plcStage plc.Stage

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
func makeResult(scan plc.Scan) (result []db.Result) {

	for _, w := range experimentValues.activeWells {
		var r db.Result
		r.WellPosition = w
		r.ExperimentID = experimentValues.experimentID
		r.Cycle = scan.Cycle
		for _, t := range experimentValues.targets {
			t.DyePosition = t.DyePosition - 1 // -1 dye position starts with 1 and Emission starts from 0
			r.TargetID = t.TargetID
			r.FValue = scan.Wells[w][t.DyePosition] // for 5th well & target 2 = scanWells[5][1]

			result = append(result, r)
		}
	}

	return
}

func wellColorAnalysis(Result []db.Result, DBWellTargets []db.WellTarget, DBWells []db.Well, currentCycle uint16) ([]db.WellTarget, []db.Well) {
	//if no well configured
	if len(DBWells) == 0 && len(DBWellTargets) == 0 {
		for _, r := range Result {
			var wt db.WellTarget
			wt.WellPosition = r.WellPosition
			wt.TargetID = r.TargetID

			wt.ExperimentID = r.ExperimentID

			if r.Threshold < float32(r.FValue) {
				// add ct value
				wt.CT = strconv.Itoa(int(r.FValue))
			} else {
				wt.CT = ""
			}

			DBWellTargets = append(DBWellTargets, wt)
		}
		return DBWellTargets, DBWells
	} else if len(DBWellTargets) > 0 && len(DBWells) == 0 { //when only targets added in prev cycle
		for _, r := range Result {
			for i, t := range DBWellTargets {
				if r.WellPosition == t.WellPosition && r.TargetID == t.TargetID {
					if t.CT == "" && r.Threshold < float32(r.FValue) {

						// add ct
						DBWellTargets[i].CT = strconv.Itoa(int(r.FValue))
					}
				}
			}
		}
		return DBWellTargets, DBWells
	} else { // determine color
		for _, r := range Result {
			for i, w := range DBWells {
				for j, t := range DBWellTargets {
					if r.WellPosition == w.Position && r.TargetID == t.TargetID && t.WellPosition == w.Position {

						switch {
						case 5 >= currentCycle && currentCycle <= 25 && float32(r.FValue) > r.Threshold && t.CT == "":
							// mark red
							DBWells[i].ColorCode = red
							DBWellTargets[j].CT = strconv.Itoa(int(r.FValue))

						case 25 <= currentCycle && float32(r.FValue) > r.Threshold && t.CT == "":
							// mark orange
							DBWells[i].ColorCode = orange
							DBWellTargets[j].CT = strconv.Itoa(int(r.FValue))

						case float32(r.FValue) > r.Threshold && t.CT == "":
							// only update ct
							DBWellTargets[j].CT = strconv.Itoa(int(r.FValue))
						case float32(r.FValue) > r.Threshold && t.CT != "":
							// only update ct
							DBWellTargets[j].CT = "UNDETERMINE"
							DBWells[i].ColorCode = red
						}

					}
				}
			}
		}
		return DBWellTargets, DBWells
	}
}

func analyseResult(result []db.Result) (finalResult []graph) {

	// ex: for 8 active wells * 6 targets * no of cycle
	for _, aw := range experimentValues.activeWells {
		var wellResult graph
		wellResult.WellPosition = aw

		for _, t := range experimentValues.targets {
			wellResult.TargetID = t.TargetID
			for _, r := range result {
				if r.WellPosition == wellResult.WellPosition && r.TargetID == wellResult.TargetID {
					wellResult.ExperimentID = r.ExperimentID
					wellResult.TargetID = r.TargetID
					wellResult.Threshold = r.Threshold
					wellResult.TotalCycles = experimentValues.plcStage.CycleCount

					// if cycle found do not add again!
					if !found(r.Cycle, wellResult.Cycle) {
						wellResult.Cycle = append(wellResult.Cycle, r.Cycle)
						wellResult.FValue = append(wellResult.FValue, r.FValue)
					}

				}

			}
			finalResult = append(finalResult, wellResult)
			wellResult.Cycle = []uint16{}
			wellResult.FValue = []uint16{}
		}

	}
	return
}

func found(key uint16, search []uint16) (found bool) {
	for _, v := range search {
		if v == key {
			found = true
			return
		}
	}
	return
}
