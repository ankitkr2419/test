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
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/goburrow/modbus"

	"github.com/rs/cors"

	logger "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/urfave/negroni"
)

// variables for Binary Build info
var Version, User, Machine, CommitID, Branch, BuiltOn string

func main() {
	logger.SetFormatter(&logger.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "02-01-2006 15:04:05",
	})

	// logging output to file and console
	var filename = "utils/output_log.txt"
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
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
			Usage: "start server [--plc {simulator|compact32}] [--delay range:(0,100] ]",
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
				&cli.IntFlag{
					Name:  "delay",
					Value: 50,
					Usage: "Input a delay in range (0, 100]",
				},
			},
			Action: func(c *cli.Context) error {
				if c.String("plc") != "simulator" && c.Int("delay") != 50 {
					return responses.SimulatorReservedDelayError
				}
				err := simulator.UpdateDelay(c.Int("delay"))
				if err != nil {
					logger.Error("Re-check delay argument")
					return err
				}
				return startApp(c.String("plc"), c.Bool("test"))
			},
		},
		{
			Name:  "create_migration",
			Usage: "create migration file",
			Action: func(c *cli.Context) error {
				return db.CreateMigrationFile(c.Args().Get(0))
			},
		},
		{
			Name:  "migrate",
			Usage: "run db migrations",
			Action: func(c *cli.Context) error {
				return db.RunMigrations()
			},
		},
		{
			Name:  "rollback",
			Usage: "rollback migrations",
			Action: func(c *cli.Context) error {

				return db.RollbackMigrations(c.Args().Get(0))
			},
		},
		{
			Name:  "import",
			Usage: "import [--recipename RECIPE_NAME] [--csv CSV_ABSOLUTE_PATH] ",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "recipename",
					Value: "recipe",
					Usage: "put recipe name",
				},
				&cli.StringFlag{
					Name:  "csv",
					Value: "recipe.csv",
					Usage: "put recipe's csv complete file path",
				},
			},
			Action: func(c *cli.Context) error {
				return db.ImportCSV(c.String("recipename"), c.String("csv"))
			},
		},
		{
			Name:  "version",
			Usage: "version",
			Action: func(c *cli.Context) {
				printBinaryInfo()
			},
		},
	}

	if err := cliApp.Run(os.Args); err != nil {
		panic(err)
	}
}

func startApp(plcName string, test bool) (err error) {
	var store db.Storer
	var driver plc.Driver
	var handler *modbus.RTUClientHandler
	var driverDeckA plc.Extraction
	var driverDeckB plc.Extraction

	if plcName != "simulator" && plcName != "compact32" {
		logger.Errorln(responses.UnsupportedPLCError)
		return
	}

	exit := make(chan error)

	websocketMsg := make(chan string)

	websocketErr := make(chan error)

	// PLC work in a completely separate go-routine!
	if plcName == "compact32" {
		driver = compact32.NewCompact32Driver(websocketMsg, websocketErr, exit, test)
		driverDeckA, handler = compact32.NewCompact32DeckDriverA(websocketMsg, websocketErr, exit, test)
		driverDeckB = compact32.NewCompact32DeckDriverB(websocketMsg, exit, test, handler)
	} else {
		driver = simulator.NewSimulator(exit)
		driverDeckA = simulator.NewExtractionSimulator(websocketMsg, websocketErr, exit, "A")
		driverDeckB = simulator.NewExtractionSimulator(websocketMsg, websocketErr, exit, "B")

	}

	store, err = db.Init()
	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.DatabaseInitError)
		return
	}

	plcDeckMap := map[string]plc.Extraction{
		"A": driverDeckA,
		"B": driverDeckB,
	}

	deps := service.Dependencies{
		Store:   store,
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
	server.Run(*addr)
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
				"A": driverDeckA,
				"B": driverDeckB,
			}
			deps.PlcDeck = plcDeckMap
		default:
			time.Sleep(5 * time.Second)
		}
	}
}

func printBinaryInfo() {
	fmt.Printf("\nVersion\t\t: %v \nUser\t\t: %v \nMachine\t\t: %v \nBranch\t\t: %v \nCommitID\t: %v \nBuilt\t\t: %v\n", Version, User, Machine, Branch, CommitID, BuiltOn)
}
