package service

import (
	"fmt"
	"io"
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
	logsPath = "./utils/logs"
	tecPath  = "./utils/tec"
)

func monitorForPLCTimeout(deps *Dependencies, exit chan error) {
	for {
		select {
		case err := <-deps.ExitCh:
			logger.Errorln(err)
			driverDeckA, handler := compact32.NewCompact32DeckDriverA(deps.WsMsgCh, deps.WsErrCh, exit, false)
			driverDeckB := compact32.NewCompact32DeckDriverB(deps.WsMsgCh, exit, false, handler)
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
		deps.PlcDeck[plc.DeckA].ReadFlapSensor()
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

	if _, err = os.Stat(logsPath); os.IsNotExist(err) {
		os.MkdirAll(logsPath, 0755)
		// ignore error and try creating log output file
	}
	if _, err = os.Stat(tecPath); os.IsNotExist(err) {
		os.MkdirAll(tecPath, 0755)
	}
	if _, err = os.Stat(ReportOutputPath); os.IsNotExist(err) {
		os.MkdirAll(ReportOutputPath, 0755)
	}
	if _, err = os.Stat(ExpOutputPath); os.IsNotExist(err) {
		os.MkdirAll(ExpOutputPath, 0755)
	}

	// All terminal logs will be noted in below file
	filename := fmt.Sprintf("%v/output_%v.log", logsPath, time.Now().Unix())
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		logger.Errorln(responses.WriteToFileError)
	}
	// logging output to file and console
	mw := io.MultiWriter(os.Stdout, f)
	logger.SetOutput(mw)
	return
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
