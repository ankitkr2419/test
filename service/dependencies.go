package service

import (
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"mylab/cpagent/tec"

	"mylab/cpagent/plc/compact32"
	"mylab/cpagent/plc/simulator"
	"mylab/cpagent/responses"
	tecSim "mylab/cpagent/tec/simulator"
	"mylab/cpagent/tec/tec_1089"

	"github.com/goburrow/modbus"

	logger "github.com/sirupsen/logrus"
)

type Dependencies struct {
	Store   db.Storer
	Tec     tec.Driver
	Plc     plc.Driver
	PlcDeck map[string]plc.Extraction
	ExitCh  <-chan error
	WsErrCh chan error
	WsMsgCh chan string
	// define other service dependencies
}

func GetAllDependencies(plcName string, test, noRTPCR, noExtraction bool) (deps *Dependencies, err error) {
	var store db.Storer
	var driver plc.Driver
	var tecDriver tec.Driver
	var handler *modbus.RTUClientHandler
	var driverDeckA, driverDeckB plc.Extraction

	if plcName != SIM && plcName != C32 {
		logger.Errorln(responses.UnsupportedPLCError)
		return nil, responses.UnsupportedPLCError
	}

	exit := make(chan error)
	websocketMsg := make(chan string)
	websocketErr := make(chan error)

	defer func() {
		if err == nil {
			// NOTE: monitorForPLCTimeout uses the same exit channel that is why it is to be here
			go monitorForPLCTimeout(deps, exit)
		}
	}()

	switch {
	case noExtraction && noRTPCR:
		logger.Infoln("application neither supports extraction nor rtpcr")
		Application = None
	case noExtraction && plcName == C32:
		driver = compact32.NewCompact32Driver(websocketMsg, websocketErr, exit, test)
		tecDriver = tec_1089.NewTEC1089Driver(websocketMsg, websocketErr, exit, test, driver)
		Application = RTPCR
	case noExtraction && plcName == SIM:
		driver = simulator.NewSimulator(exit)
		tecDriver = tecSim.NewSimulatorDriver(websocketMsg, websocketErr, exit, test)
		Application = RTPCR
	case noRTPCR && plcName == C32:
		driverDeckA, handler = compact32.NewCompact32DeckDriverA(websocketMsg, websocketErr, exit, test)
		driverDeckB = compact32.NewCompact32DeckDriverB(websocketMsg, exit, test, handler)
		Application = Extraction
	case noRTPCR && plcName == SIM:
		driverDeckA = simulator.NewExtractionSimulator(websocketMsg, websocketErr, exit, plc.DeckA)
		driverDeckB = simulator.NewExtractionSimulator(websocketMsg, websocketErr, exit, plc.DeckB)
		Application = Extraction
		// Only cases that remain are of combined RTPCR and Extraction
	case plcName == C32:
		driver = compact32.NewCompact32Driver(websocketMsg, websocketErr, exit, test)
		driverDeckA, handler = compact32.NewCompact32DeckDriverA(websocketMsg, websocketErr, exit, test)
		driverDeckB = compact32.NewCompact32DeckDriverB(websocketMsg, exit, test, handler)
		tecDriver = tec_1089.NewTEC1089Driver(websocketMsg, websocketErr, exit, test, driver)
		Application = Combined
	case plcName == SIM:
		driver = simulator.NewSimulator(exit)
		driverDeckA = simulator.NewExtractionSimulator(websocketMsg, websocketErr, exit, plc.DeckA)
		driverDeckB = simulator.NewExtractionSimulator(websocketMsg, websocketErr, exit, plc.DeckB)
		tecDriver = tecSim.NewSimulatorDriver(websocketMsg, websocketErr, exit, test)
		Application = Combined
	default:
		logger.Errorln(responses.UnknownCase)
		return nil, responses.UnknownCase
	}

	// PLC work in a completely separate go-routine!

	store, err = db.Init()
	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.DatabaseInitError)
		return
	}

	plcDeckMap := map[string]plc.Extraction{
		plc.DeckA: driverDeckA,
		plc.DeckB: driverDeckB,
	}

	return &Dependencies{
		Store:   store,
		Tec:     tecDriver,
		Plc:     driver,
		PlcDeck: plcDeckMap,
		ExitCh:  exit,
		WsErrCh: websocketErr,
		WsMsgCh: websocketMsg,
	}, nil
}
