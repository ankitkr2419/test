package main

import (
	"io"

	rice "github.com/GeertJohan/go.rice"

	"flag"
	"fmt"
	"mylab/cpagent/config"
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"mylab/cpagent/plc/compact32"
	"mylab/cpagent/plc/simulator"
	"mylab/cpagent/responses"
	"mylab/cpagent/service"
	"mylab/cpagent/tec"
	tecSim "mylab/cpagent/tec/simulator"
	"mylab/cpagent/tec/tec_1089"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/goburrow/modbus"

	"github.com/rs/cors"

	logger "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/urfave/negroni"
)

const (
	C32 = "compact32"
	SIM = "simulator"
)

const(
	logsPath = "./utils/logs"
	tecPath = "./utils/tec"
	configPath = "./conf"
)

func main() {
	logger.SetFormatter(&logger.TextFormatter{
		FullTimestamp:   true,
		ForceColors:     true,
		TimestampFormat: "02-01-2006 15:04:05",
	})

	// logging output to file and console
	if _, err := os.Stat(logsPath); os.IsNotExist(err) {
		os.MkdirAll(logsPath, 0755)
		// ignore error and try creating log output file
	}
	if _, err := os.Stat(tecPath); os.IsNotExist(err) {
		os.MkdirAll(tecPath, 0755)
		// ignore error and try creating log output file
	}

	filename := fmt.Sprintf("%v/output_%v.log", logsPath, time.Now().Unix())
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		logger.Errorln(responses.WriteToFileError)
	}
	mw := io.MultiWriter(os.Stdout, f)
	logger.SetOutput(mw)

	config.LoadAllConfs()

	cliApp := cli.NewApp()
	cliApp.Name = config.AppName()
	cliApp.Version = "1.0.0"
	cliApp.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "start server [--plc {simulator|compact32}] [--test] [--no-extraction] [--no-rtpcr] [--delay range:(0,100] ]",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "plc",
					Value: "compact32",
					Usage: "Choose the PLC. simulator|compact32",
				},
				&cli.BoolFlag{
					Name:  "test",
					Usage: "Run in test mode!",
				},
				&cli.BoolFlag{
					Name:  "no-extraction",
					Usage: "Run without extraction",
				},
				&cli.BoolFlag{
					Name:  "no-rtpcr",
					Usage: "Run withour rtpcr",
				},
				&cli.IntFlag{
					Name:  "delay",
					Value: 50,
					Usage: "Input a delay in range (0, 100]",
				},
			},
			Action: func(c *cli.Context) error {
				if c.String("plc") != SIM && c.Int("delay") != 50 {
					return responses.SimulatorReservedDelayError
				}
				err := simulator.UpdateDelay(c.Int("delay"))
				if err != nil {
					logger.Error("Re-check delay argument")
					return err
				}
				return startApp(c.String("plc"), c.Bool("test"), c.Bool("no-rtpcr"), c.Bool("no-extraction"))
			},
		},
		{
			Name:  "create_migration",
			Usage: "create migration file",
			Action: func(c *cli.Context) error {
				logger.Infoln("Creating migration -->", c.Args().Get(0))
				return db.CreateMigrationFile(c.Args().Get(0))
			},
		},
		{
			Name:  "migrate",
			Usage: "run db migrations",
			Action: func(c *cli.Context) error {
				logger.Infoln("Running migrations")
				return db.RunMigrations()
			},
		},
		{
			Name:  "rollback",
			Usage: "rollback migrations",
			Action: func(c *cli.Context) error {
				logger.Infoln("Rolling back migrations by ", c.Args().Get(0), " steps")
				return db.RollbackMigrations(c.Args().Get(0))
			},
		},
		{
			Name:  "import",
			Usage: "import --csv CSV_ABSOLUTE_PATH ",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "csv",
					Value: "recipe.csv",
					Usage: "put recipe's csv complete file path",
				},
			},
			Action: func(c *cli.Context) error {
				logger.Infoln("Importing CSV named -->", c.String("csv"))
				return db.ImportCSV(c.String("csv"))
			},
		},
		{
			Name:  "version",
			Usage: "version",
			Action: func(c *cli.Context) {
				logger.Infoln("Printing Version Information")
				service.PrintBinaryInfo()
			},
		},
	}

	if err := cliApp.Run(os.Args); err != nil {
		panic(err)
	}
}

func startApp(plcName string, test, noRTPCR, noExtraction bool) (err error) {
	logger.Println("run in test mode --->", test)
	var store db.Storer
	var driver plc.Driver
	var tecDriver tec.Driver
	var handler *modbus.RTUClientHandler
	var driverDeckA plc.Extraction
	var driverDeckB plc.Extraction

	if plcName != SIM && plcName != C32 {
		logger.Errorln(responses.UnsupportedPLCError)
		return
	}

	exit := make(chan error)
	websocketMsg := make(chan string)
	websocketErr := make(chan error)

	switch {
	case noExtraction && noRTPCR:
		logger.Infoln("application neither supports extraction nor rtpcr")
		service.Application = service.None
	case noExtraction && plcName == C32:
		driver = compact32.NewCompact32Driver(websocketMsg, websocketErr, exit, test)
		tecDriver = tec_1089.NewTEC1089Driver(websocketMsg, websocketErr, exit, test, driver)
		service.Application = service.RTPCR
	case noExtraction && plcName == SIM:
		driver = simulator.NewSimulator(exit)
		tecDriver = tecSim.NewSimulatorDriver(websocketMsg, websocketErr, exit, test)
		service.Application = service.RTPCR
	case noRTPCR && plcName == C32:
		driverDeckA, handler = compact32.NewCompact32DeckDriverA(websocketMsg, websocketErr, exit, test)
		driverDeckB = compact32.NewCompact32DeckDriverB(websocketMsg, exit, test, handler)
		service.Application = service.Extraction
	case noRTPCR && plcName == SIM:
		driverDeckA = simulator.NewExtractionSimulator(websocketMsg, websocketErr, exit, plc.DeckA)
		driverDeckB = simulator.NewExtractionSimulator(websocketMsg, websocketErr, exit, plc.DeckB)
		service.Application = service.Extraction
		// Only cases that remain are of combined RTPCR and Extraction
	case plcName == C32:
		driver = compact32.NewCompact32Driver(websocketMsg, websocketErr, exit, test)
		driverDeckA, handler = compact32.NewCompact32DeckDriverA(websocketMsg, websocketErr, exit, test)
		driverDeckB = compact32.NewCompact32DeckDriverB(websocketMsg, exit, test, handler)
		tecDriver = tec_1089.NewTEC1089Driver(websocketMsg, websocketErr, exit, test, driver)
		service.Application = service.Combined
	case plcName == SIM:
		driver = simulator.NewSimulator(exit)
		driverDeckA = simulator.NewExtractionSimulator(websocketMsg, websocketErr, exit, plc.DeckA)
		driverDeckB = simulator.NewExtractionSimulator(websocketMsg, websocketErr, exit, plc.DeckB)
		tecDriver = tecSim.NewSimulatorDriver(websocketMsg, websocketErr, exit, test)
		service.Application = service.Combined
	default:
		logger.Errorln(responses.UnknownCase)
		return responses.UnknownCase
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

	deps := service.Dependencies{
		Store:   store,
		Tec:     tecDriver,
		Plc:     driver,
		PlcDeck: plcDeckMap,
		ExitCh:  exit,
		WsErrCh: websocketErr,
		WsMsgCh: websocketMsg,
	}

	go monitorForPLCTimeout(&deps, exit)

	err = service.LoadAllServiceFuncs(store)
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
		return
	}

	var addr = flag.String("addr", "0.0.0.0:"+strconv.Itoa(config.AppPort()), "http service address")
	// mux router
	router := service.InitRouter(deps)

	// to embed react build with go rice
	router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("./web-client/build").HTTPBox()))

	// cors configuration for front-end
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"PUT", "DELETE", "POST", "GET"},
		AllowedHeaders: []string{"*"},
	})

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
	})

	// init web server
	server := negroni.Classic()
	server.Use(c)
	server.UseHandler(router)

	flag.Parse()

	idleConnsClosed := make(chan struct{})

	go func() {

		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
		<-signals

		// We received an interrupt signal, shut down.
		logger.Warnln("..................\n----Application shutting down gracefully ----|\n.............................................|")
		err = deps.Tec.ReachRoomTemp()
		if err != nil {
			logger.Errorln("Couldn't reach the room temp!")
			os.Exit(-1)
		}
		deps.Plc.SwitchOffLidTemp()
		os.Exit(0)
	}()

	server.Run(*addr)
	<-idleConnsClosed

	return
}

func monitorForPLCTimeout(deps *service.Dependencies, exit chan error) {
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
