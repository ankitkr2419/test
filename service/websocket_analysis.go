package service

import (
	"fmt"
	"math"
	"mylab/cpagent/config"
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

func setExperimentValues(aw []int32, t uuid.UUID, TargetDetails []db.TargetDetails, ExperimentID uuid.UUID, stage plc.Stage) {
	experimentValues = experimentResultValues{
		experimentID: ExperimentID,
		activeWells:  aw,
		targets:      TargetDetails,
		plcStage:     stage,
		icTargetID:   t,
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
		step.DataCapture = s.DataCapture

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
func makeResult(scan plc.Scan, file *excelize.File) (result []db.Result) {
	var wellFval []uint16
	for _, w := range experimentValues.activeWells {
		var r db.Result
		r.WellPosition = w
		r.ExperimentID = experimentValues.experimentID
		r.Cycle = scan.Cycle
		for _, t := range experimentValues.targets {
			t.DyePosition = t.DyePosition - 1 // -1 dye position starts with 1 and Emission starts from 0
			r.TargetID = t.TargetID
			r.FValue = scan.Wells[w-1][t.DyePosition] // for 5th well & target 2 = scanWells[5][1] //w-1 as emissions starts from 0
			wellFval = append(wellFval, r.FValue)
			result = append(result, r)
		}

	}
	row := []interface{}{fmt.Sprintf("cycle %d", scan.Cycle)}
	for _, v := range wellFval {
		row = append(row, v)
	}
	db.AddRowToExcel(file, db.RTPCRSheet, row)

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
					if r.WellPosition == w.Position && r.TargetID == t.TargetID && t.WellPosition == w.Position && r.TargetID != experimentValues.icTargetID {

						switch {
						case redlowerlimit <= currentCycle && currentCycle < redupperlimit && scaleThreshold(float32(r.FValue)) >= r.Threshold && t.CT == "" && DBWells[i].ColorCode == green:
							// mark red
							DBWells[i].ColorCode = red
							DBWellTargets[j].CT = strconv.Itoa(int(r.FValue))

						case orangelowerlimit <= currentCycle && scaleThreshold(float32(r.FValue)) >= r.Threshold && t.CT == "" && DBWells[i].ColorCode == green:
							// mark orange
							DBWells[i].ColorCode = orange
							DBWellTargets[j].CT = strconv.Itoa(int(r.FValue))

						case redlowerlimit > currentCycle && scaleThreshold(float32(r.FValue)) >= r.Threshold && t.CT == "" && DBWells[i].ColorCode == green:

							DBWellTargets[j].CT = strconv.Itoa(int(r.FValue))
							DBWells[i].ColorCode = red // here, we do detemine color as cycle is 1 to lowerLimitOfRed also crosses threshold

						case scaleThreshold(float32(r.FValue)) <= r.Threshold && t.CT != "":
							DBWellTargets[j].CT = undetermine // undertermine is marked when second time graph cuts threshold line
							DBWells[i].ColorCode = red

						case t.CT != "" && DBWells[i].ColorCode == green: //if earlier CT value is updated when well was not configured then change only color of the well
							DBWells[i].ColorCode = red

						case scaleThreshold(float32(r.FValue)) >= r.Threshold && t.CT == "" && DBWells[i].ColorCode != green: //when color is already marked we have only update CT Value
							DBWellTargets[j].CT = strconv.Itoa(int(r.FValue))

						}

					} else if r.WellPosition == w.Position && r.TargetID == t.TargetID && t.WellPosition == w.Position && r.TargetID == experimentValues.icTargetID {
						// for IC: internal control target only update CT Values, do not update color
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

func analyseResult(result []db.Result, wells []int32, targets []db.TargetDetails, cycles uint16) (finalResult []graph) {

	// ex: for 8 active wells * 6 targets * no of cycle
	for _, aw := range wells {
		var wellResult graph
		wellResult.WellPosition = aw

		for _, t := range targets {
			wellResult.TargetID = t.TargetID
			for _, r := range result {
				if r.WellPosition == wellResult.WellPosition && r.TargetID == wellResult.TargetID {
					wellResult.ExperimentID = r.ExperimentID
					wellResult.TargetID = r.TargetID
					wellResult.Threshold = r.Threshold
					wellResult.TotalCycles = cycles

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

func analyseResultForThreshold(result []db.Result, threshold float32, DBWells []db.Well, wellTargets []db.WellTarget) (DBWellTargets []db.WellTarget) {
	for _, r := range result {
		for i, t := range wellTargets {
			for _, w := range DBWells {
				if r.WellPosition == w.Position && r.TargetID == t.TargetID {
					logger.Infoln(r.FValue, t.CT, w.Position, t.TargetName)
					if t.CT == "" && threshold <= float32(r.FValue) {
						DBWellTargets[i].CT = strconv.Itoa(int(r.FValue))
					} else if t.CT != "" && t.CT != undetermine && r.Threshold >= float32(r.FValue) {
						// if ct value again crosses threshold then only set it as undertermine
						DBWellTargets[i].CT = undetermine
					}
				}
			}
		}
	}
	// ex: for 8 active wells * 6 targets * no of cycle

	return
}

func getAutoThreshold(result []db.Result, wells []int32, targets []db.TargetDetails, cycles uint16) (thresholdLine map[db.TargetDetails]float32) {

	formulaWellTarget := make(map[TargetCycleWell]float32, cycles)
	var finalSum float32
	thresholdLine = make(map[db.TargetDetails]float32, len(targets))
	// ex: for 8 active wells * 6 targets * no of cycle

	for _, t := range targets {

		var wellResult graph
		var key TargetCycleWell

		wellResult.TargetID = t.TargetID
		for _, aw := range wells {
			var sum uint16
			var avg, std float32
			wellResult.WellPosition = aw
			for _, r := range result {
				if r.WellPosition == wellResult.WellPosition && r.TargetID == wellResult.TargetID {
					sum = sum + r.FValue
					wellResult.FValue = append(wellResult.FValue, float32(r.FValue))
				}
			}
			key = TargetCycleWell{
				Target: wellResult.TargetID,
				Well:   wellResult.WellPosition,
			}
			avg = float32(sum / cycles)
			std = calculateStandardDeviation(wellResult.FValue, avg)
			formulaWellTarget[key] = avg + 10*std
			wellResult.FValue = []float32{}
		}

		for _, v := range formulaWellTarget {
			finalSum = finalSum + v
		}

		thresholdLine[t] = finalSum / float32(len(formulaWellTarget))
	}

	return
}

func calculateStandardDeviation(array []float32, average float32) (deviation float32) {

	var deviatedSum float32
	for _, v := range array {
		value := v - average
		deviatedSum = deviatedSum + value*value
	}

	deviatedAverage := deviatedSum / float32(len(array)-1)

	deviation = float32(math.Sqrt(float64(deviatedAverage)))
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
