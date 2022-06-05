package plc

import (
	"context"
	"errors"
	"mylab/cpagent/db"
	"time"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

const (
	ErrorExtractionMonitor = "ErrorExtractionMonitor"
	ErrorLidPIDTuning      = "PID Error"
	ErrorOperationAborted  = "ErrorOperationAborted"
)

type Status int32

const (
	OK      Status = iota // all good
	ERROR                 // something is wrong
	RESTART               // cycle was interrupted, can restart
)

type Step struct {
	TargetTemp  float32 `json:"target_temp"` // holding temperature for step
	RampUpTemp  float32 `json:"ramp_rate"`   // ramp-up temperature for step
	HoldTime    int32   `json:"hold_time"`   // hold time for step
	DataCapture bool    `json:"data_capture"`
}

// We can have at most 4 Holding steps and 6 Cycling steps.
type Stage struct {
	Holding      []Step // 4 steps
	Cycle        []Step // 6 steps
	CycleCount   uint16 // number of cycles to run
	IdealLidTemp uint16 // ideal lid temp
}

type Emissions [6]uint16

type Scan struct {
	Cycle         uint16 // current running cycle
	Wells         [16]Emissions
	Temp          float32
	LidTemp       float32
	CycleComplete bool
}

type Driver interface {
	SelfTest() Status               // Check if Homing or any other errors during bootup of PLC {ERROR | RESTART | OK}
	HeartBeat()                     // Attempt to write heartbeat 3 times else fail
	ConfigureRun(Stage) error       // Configure the various holding and cycling stages
	Start() error                   // trigger the start of the cycling process
	Stop() error                    // Stop the cycle, Status: ABORT (if pre-emptive) OK: All Cycles have completed
	Monitor(uint16) (Scan, error)   // Monitor periodically. If Status=CYCLE_COMPLETE, the Scan will be populated
	Calibrate() error               // TBD
	HomingRTPCR() error             // Homing of RTPCR
	Reset() error                   // reseting the values
	Cycle() error                   // start the cycle
	SetLidTemp(uint16) error        // set Lid Temperature
	SwitchOffLidTemp() error        // Lid will return to room temp
	LidPIDCalibration() error       // LID PID Tuning
	SetScanSpeedAndScanTime() error // Scan speed and Scan Time is settable by user
	CalculateOpticalResult(dye db.Dye, kitID string, knownValue, cycleCount int64) (opticalResult []db.DyeWellTolerance, err error)
}

type HeaterData struct {
	Shaker1Temp float64 `json:"shaker_1_temp"`
	Shaker2Temp float64 `json:"shaker_2_temp"`
	Deck        string  `json:"deck"`
	HeaterOn    bool    `json:"heater_on"`
}

type WSData struct {
	Progress         float64          `json:"progress"`
	Deck             string           `json:"deck"`
	Status           string           `json:"status"`
	OperationDetails OperationDetails `json:"operation_details"`
}

type WSError struct {
	Message string `json:"message"`
	Deck    string `json:"deck"`
}

type OperationDetails struct {
	Message        string         `json:"message"`
	CurrentStep    int64          `json:"current_step,omitempty"`
	TotalProcesses int64          `json:"total_processes,omitempty"`
	RecipeID       uuid.UUID      `json:"recipe_id,omitempty"`
	RemainingTime  *TimeHMS       `json:"remaining_time,omitempty"`
	TotalTime      *TimeHMS       `json:"total_time,omitempty"`
	ProcessName    string         `json:"process_name,omitempty"`
	ProcessType    db.ProcessType `json:"process_type,omitempty"`
	Progress       *int64         `json:"progress,omitempty"`
	TotalCycles    int64          `json:"total_cycles,omitempty"`
}

type TimeHMS struct {
	Hours   uint8 `json:"hours"`
	Minutes uint8 `json:"minutes"`
	Seconds uint8 `json:"seconds"`
}

type Compact32Deck struct {
	name       string
	ExitCh     chan error
	WsErrCh    chan error
	WsMsgCh    chan string
	DeckDriver Compact32Driver
}

// Internal Interface to ensure sync'ing and testing modbus interfaces
type Compact32Driver interface {
	WriteSingleRegister(address, value uint16) (results []byte, err error)
	WriteMultipleRegisters(address, quantity uint16, value []byte) (results []byte, err error)
	ReadCoils(address, quantity uint16) (results []byte, err error)
	ReadSingleCoil(address uint16) (value uint16, err error)
	ReadHoldingRegisters(address, quantity uint16) (results []byte, err error)
	ReadSingleRegister(address uint16) (value uint16, err error)
	WriteSingleCoil(address, value uint16) (err error)
}

type Extraction interface {
	AspireDispense(ad db.AspireDispense, cartridgeID int64) (response string, err error)
	AttachDetach(ad db.AttachDetach) (response string, err error)
	DiscardBoxCleanup() (response string, err error)
	RestoreDeck() (response string, err error)
	UVLight(uvTime int64) (response string, err error)
	AddDelay(delay db.Delay, runRecipe bool) (response string, err error)
	DiscardTipAndHome(discard bool) (response string, err error)
	Heating(ht db.Heating, live bool) (response string, err error)
	Homing() (response string, err error)
	ManualMovement(motorNum, direction uint16, distance float32) (response string, err error)
	Resume() (response string, err error)
	Pause() (response string, err error)
	Abort() (response string, err error)
	Piercing(pi db.Piercing, cartridgeID int64) (response string, err error)
	Shaking(shakerData db.Shaker, live bool) (response string, err error)
	TipDocking(td db.TipDock, cartridgeID int64) (response string, err error)
	SetRunInProgress()
	SetPaused()
	ResetPaused()
	ResetAborted()
	ResetRunInProgress()
	IsMachineHomed() bool
	IsRunInProgress() bool
	IsHeaterInProgress() bool
	IsShakerInProgress() bool
	TipOperation(to db.TipOperation) (response string, err error)
	RunRecipeWebsocketData(recipe db.Recipe, processes []db.Process) (err error)
	SetCurrentProcessNumber(step int64)
	SwitchOffAllCoils() (response string, err error)
	PIDCalibration(context.Context) error
	SetEngineerOrAdminLogged(value bool)
	HeaterData() error
	StartRecipeTimeCounter()
	UpdateRecipeEstimatedTime(ctx context.Context, recipeID uuid.UUID, s db.Storer, err *error) error
	ReadFlapSensor() error
	ShutDown() error
}

func SetDeckName(C32 *Compact32Deck, deck string) {
	C32.name = deck
}

func HoldSleep(sleepTime int32) (err error) {

	var elaspedTime int32
	for {
		logger.Infoln("plc.ExperimentRunning && elaspedTime < sleepTime ", ExperimentRunning, elaspedTime, sleepTime)
		if ExperimentRunning && elaspedTime < sleepTime {
			time.Sleep(time.Second * 1)
			logger.Infoln("sleeping in holdsleep")
		} else {
			if !ExperimentRunning {
				logger.Errorln("experiment has stoped running")
				return errors.New("experiment has stoped running")
			}
			return nil
		}
		elaspedTime = elaspedTime + 1
	}
}
