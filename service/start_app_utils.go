package service

import (
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"os"
	"os/signal"
	"syscall"
	"time"

	"mylab/cpagent/plc/compact32"
	"mylab/cpagent/responses"

	logger "github.com/sirupsen/logrus"
)

const (
	C32 = "compact32"
	SIM = "simulator"
)

const (
	utilsPath = "$HOME/cpagent/utils/"

	logsPath         = utilsPath + "logs"
	tecPath          = utilsPath + "tec"
	expOutputPath    = utilsPath + "output"
	reportOutputPath = utilsPath + "reports"
)

func monitorForPLCTimeout(deps *Dependencies, exit chan error) {
	for {
		select {
		case err := <-deps.ExitCh:
			logger.Errorln(err)
			driverDeckA, handler := compact32.NewCompact32DeckDriverA(deps.WsMsgCh, deps.WsErrCh, exit, false)
			driverDeckB := compact32.NewCompact32DeckDriverB(deps.WsMsgCh, deps.WsErrCh, exit, false, handler)
			plcDeckMap := map[string]plc.Extraction{
				plc.DeckA: driverDeckA,
				plc.DeckB: driverDeckB,
			}
			deps.PlcDeck = plcDeckMap
		default:
			time.Sleep(5 * time.Second)
		}
	}
}

// M44 for White Light
// Magnet M45
// Flap M46 LH, M47 RH -> Pause and show warning with cancel, abort option
// Handle Flap Open in Discard

func monitorFlapSensor(deps *Dependencies) {
	for {
		// TODO: monitor Flap RH, LH -> Both Decks
		// deps.PlcDeck[plc.DeckA].ReadFlapSensor()
		// deps.PlcDeck[plc.DeckB].ReadFlapSensor()
		time.Sleep(5 * time.Second)
	}
}

func SendHeaterDataToEng(deps Dependencies) {
	go deps.PlcDeck[plc.DeckA].HeaterData()
	go deps.PlcDeck[plc.DeckB].HeaterData()
}

func WaitForGracefulShutdown(deps Dependencies, idleConnsClosed chan struct{}) {

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	<-signals

	ShutDownGracefully(deps)
}

func SetLoggersAndFiles() (err error) {

	logger.SetFormatter(&logger.TextFormatter{
		FullTimestamp:   true,
		ForceColors:     true,
		TimestampFormat: "02-01-2006 15:04:05",
	})

	if _, err = os.Stat(os.ExpandEnv(logsPath)); os.IsNotExist(err) {
		os.MkdirAll(os.ExpandEnv(logsPath), 0755)
		// ignore error and try creating log output file
	}
	if _, err = os.Stat(os.ExpandEnv(tecPath)); os.IsNotExist(err) {
		os.MkdirAll(os.ExpandEnv(tecPath), 0755)
	}
	if _, err = os.Stat(os.ExpandEnv(reportOutputPath)); os.IsNotExist(err) {
		os.MkdirAll(os.ExpandEnv(reportOutputPath), 0755)
	}
	if _, err = os.Stat(os.ExpandEnv(expOutputPath)); os.IsNotExist(err) {
		os.MkdirAll(os.ExpandEnv(expOutputPath), 0755)
	}
	//NOTE: if we don't return nil below program will panic for the first time
	logger.SetOutput(os.Stdout)
	return nil
}

func LoadAllSetups(store db.Storer) (err error) {
	err = LoadAllServiceFuncs(store)
	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.ServiceAllLoadError)
		return
	}

	err = db.LoadAllDBSetups(store)
	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.DBAllSetupError)
		return
	}

	err = plc.LoadAllPLCFuncs(store)
	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.PLCAllLoadError)
	}
	return
}
