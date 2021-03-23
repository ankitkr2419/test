package service

import (
	"mylab/cpagent/db"
	"mylab/cpagent/plc"

	"github.com/google/uuid"
)

var (
	green             = "#3FC13A" // All CT values for the well are below threshold,
	red               = "#F06666" //Even a single value crosses threshold for target
	orange            = "#F3811F" // If the CT values are close to threshold (delta)
	experimentRunning = false     // In case of pre-emptive stop we need to send signal to monitor through this flag
	experimentValues  experimentResultValues
	redlowerlimit     uint16
	redupperlimit     uint16
	orangelowerlimit  uint16
	pcrMin            float32 = 0 //actual scale of emission values
	pcrMax            float32 = 32000
	graphMin          float32 = 0 // scale for graph
	graphMax          float32 = 10
	undetermine               = "UNDETERMINE"
	
)

type experimentResultValues struct {
	plcStage     plc.Stage
	experimentID uuid.UUID
	activeWells  []int32
	targets      []db.TargetDetails
	icTargetID   uuid.UUID
}

type resultGraph struct {
	Type string  `json:"type"`
	Data []graph `json:"data"`
}

type graph struct {
	WellPosition int32     `db:"well_position" json:"well_position"`
	TargetID     uuid.UUID `db:"target_id" json:"target_id"`
	ExperimentID uuid.UUID `db:"experiment_id" json:"experiment_id"`
	TotalCycles  uint16    `db:"total_cycles" json:"total_cycles"`
	Cycle        []uint16  `db:"cycle" json:"cycle"`
	FValue       []float32 `db:"f_value" json:"f_value"`
	Threshold    float32   `db:"threshold" json:"threshold"`
}
type resultWells struct {
	Type string    `json:"type"`
	Data []db.Well `json:"data"`
}

type resultOnSuccess struct {
	Type string        `json:"type"`
	Data db.Experiment `json:"data"`
}

type resultOnFail struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type experimentTemperature struct {
	Type string                     `json:"type"`
	Data []db.ExperimentTemperature `json:"data"`
}

type resultsOnHoming struct{
	InProgress bool
 Success bool
 err bool
}