package main

import (
	rice "github.com/GeertJohan/go.rice"

	"flag"
	"mylab/cpagent/config"
	"mylab/cpagent/db"

	"mylab/cpagent/plc/simulator"
	"mylab/cpagent/responses"
	"mylab/cpagent/service"

	"net/http"
	"os"
	"strconv"

	"github.com/rs/cors"

	logger "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/urfave/negroni"
)

var cliCommand = []cli.Command{
	{
		Name:  "start",
		Usage: "start [--plc {simulator|compact32}] [--tec {simulator|compact32}] [--test] [--no-extraction] [--no-rtpcr] [--delay range:(0,100] ]",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "plc",
				Value: "compact32",
				Usage: "Choose the PLC. simulator|compact32",
			},
			&cli.StringFlag{
				Name:  "tec",
				Value: "simulator",
				Usage: "Choose the TEC. simulator|compact32",
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
			if c.String("plc") != service.SIM && c.Int("delay") != 50 {
				return responses.SimulatorReservedDelayError
			}
			err := simulator.UpdateDelay(c.Int("delay"))
			if err != nil {
				logger.Error("Re-check delay argument")
				return err
			}
			return getDependenciesAndStartApp(c.String("plc"), c.String("tec"), c.Bool("test"), c.Bool("no-rtpcr"), c.Bool("no-extraction"))
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

func main() {

	err := service.SetLoggersAndFiles()
	if err != nil {
		panic(err)
	}

	config.LoadAllConfs()

	cliApp := cli.NewApp()
	cliApp.Name = config.AppName()
	cliApp.Version = service.Version
	cliApp.Commands = cliCommand

	if err := cliApp.Run(os.Args); err != nil {
		panic(err)
	}
}

func getDependenciesAndStartApp(plcName, tecName string, test, noRTPCR, noExtraction bool) (err error) {
	logger.Println("run in test mode --->", test)
	var deps service.Dependencies

	if deps, err = service.GetAllDependencies(plcName, tecName, test, noRTPCR, noExtraction); err != nil {
		logger.Errorln("Getting Dependencies failed!")
		return
	}

	return startApp(deps)
}

func startApp(deps service.Dependencies) (err error) {

	err = service.LoadAllSetups(deps.Store)
	if err != nil {
		logger.Errorln("loading All Setups failed!")
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

	go service.WaitForGracefulShutdown(deps, idleConnsClosed)

    // handle global panics
    defer func(){
        if r := recover(); r != nil {
            logger.Errorln("Program panicked: ", r)
            service.ShutDownGracefully(deps)
        }
    }()

	server.Run(*addr)
	<-idleConnsClosed

	return
}
