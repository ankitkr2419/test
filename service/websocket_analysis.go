package service

import (
	"mylab/cpagent/config"
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

	redlowerlimit = config.GetColorLimits("redlowerlimit")
	redupperlimit = config.GetColorLimits("redupperlimit")
	orangelowerlimit = config.GetColorLimits("orangelowerlimit")

}

// scale threshold
func scaleThreshold(val float32) float32 {

	return ((val-pcrMin)/(pcrMax-pcrMin))*(graphMax-graphMin) + graphMin
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
			r.FValue = scan.Wells[w-1][t.DyePosition] // for 5th well & target 2 = scanWells[5][1] //w-1 as emissions starts from 0

			result = append(result, r)
		}
	}

	return
}

func wellColorAnalysis(Result []db.Result, DBWellTargets []db.WellTarget, DBWells []db.Well, currentCycle uint16) ([]db.WellTarget, []db.Well) {
	//if no well configured
	if len(DBWellTargets) > 0 && len(DBWells) == 0 { //when only targets added in prev cycle
		for _, r := range Result {
			for i, t := range DBWellTargets {
				if r.WellPosition == t.WellPosition && r.TargetID == t.TargetID {
					if t.CT == "" && r.Threshold <= scaleThreshold(float32(r.FValue)) {

						// add ct
						DBWellTargets[i].CT = strconv.Itoa(int(r.FValue))
					} else if t.CT != "" && t.CT != undetermine && r.Threshold >= scaleThreshold(float32(r.FValue)) {

						// if ct value again crosses threshold then only set it as undertermine
						DBWellTargets[i].CT = undetermine
					}
				}
			}
		}
		return DBWellTargets, DBWells
	} else {
		wellsConfigured := make([]uint16, len(DBWells))
		for _, w := range DBWells {
			wellsConfigured = append(wellsConfigured, uint16(w.Position))
		}
		for _, r := range Result {
			for j, t := range DBWellTargets {
				for i, w := range DBWells {
					// determine color
					if r.WellPosition == w.Position && r.TargetID == t.TargetID && t.WellPosition == w.Position {

						switch {
						case redlowerlimit <= currentCycle && currentCycle < redupperlimit && scaleThreshold(float32(r.FValue)) >= r.Threshold && t.CT == "":
							// mark red
							DBWells[i].ColorCode = red
							DBWellTargets[j].CT = strconv.Itoa(int(r.FValue))

						case orangelowerlimit <= currentCycle && scaleThreshold(float32(r.FValue)) >= r.Threshold && t.CT == "":
							// mark orange
							DBWells[i].ColorCode = orange
							DBWellTargets[j].CT = strconv.Itoa(int(r.FValue))

						case redlowerlimit > currentCycle && scaleThreshold(float32(r.FValue)) >= r.Threshold && t.CT == "":

							DBWellTargets[j].CT = strconv.Itoa(int(r.FValue))
							DBWells[i].ColorCode = red // here, we do detemine color as cycle is 1 to lowerLimitOfRed also crosses threshold

						case scaleThreshold(float32(r.FValue)) <= r.Threshold && t.CT != "":
							DBWellTargets[j].CT = undetermine // undertermine is marked when second time graph cuts threshold line
							DBWells[i].ColorCode = red

						case t.CT != "" && DBWells[i].ColorCode == green: //if earlier CT value is updated when well was not configured then change only color of the well
							DBWells[i].ColorCode = red

						}

					}
				}
				if r.WellPosition == t.WellPosition && r.TargetID == t.TargetID && !found(uint16(t.WellPosition), wellsConfigured) {

					// if well is not configured for any target we should not miss the CT update
					if t.CT == "" && r.Threshold <= scaleThreshold(float32(r.FValue)) {

						// add ct
						DBWellTargets[j].CT = strconv.Itoa(int(r.FValue))
					} else if t.CT != "" && t.CT != undetermine && r.Threshold >= scaleThreshold(float32(r.FValue)) {

						// if ct value again crosses threshold then only set it as undertermine
						DBWellTargets[j].CT = undetermine
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
						wellResult.FValue = append(wellResult.FValue, scaleThreshold(float32(r.FValue)))
					}

				}

			}
			finalResult = append(finalResult, wellResult)
			wellResult.Cycle = []uint16{}
			wellResult.FValue = []float32{}
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

// initializeWellTargets adds all well targets at start of experiment
func initializeWellTargets() (WTs []db.WellTarget) {
	for _, w := range experimentValues.activeWells {

		for _, t := range experimentValues.targets {

			var wt db.WellTarget

			wt.ExperimentID = experimentValues.experimentID
			wt.TargetID = t.TargetID
			wt.WellPosition = w
			wt.CT = ""

			WTs = append(WTs, wt)

		}
	}
	return
}
