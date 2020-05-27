package main

// @APITitle Main
// @APIDescription Main API for Microservices in Go!

import (
	"fmt"
	"mylab/cpagent/config"
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"mylab/cpagent/plc/compact32"
	"mylab/cpagent/plc/simulator"
	"mylab/cpagent/service"
	"os"
	"strconv"

	logger "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/urfave/negroni"
)

func main() {
	logger.SetFormatter(&logger.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "02-01-2006 15:04:05",
	})

	config.Load("application")

	cliApp := cli.NewApp()
	cliApp.Name = config.AppName()
	cliApp.Version = "1.0.0"
	cliApp.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "start server [--plc {simulator|compact32}]",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "plc",
					Value: "compact32",
					Usage: "Choose the PLC. simulator|compact32",
				},
			},
			Action: func(c *cli.Context) error {
				return startApp(c.String("plc"))
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
	}

	if err := cliApp.Run(os.Args); err != nil {
		panic(err)
	}
}

func startApp(plcName string) (err error) {
	var store db.Storer
	var driver plc.Driver

	if plcName != "simulator" && plcName != "compact32" {
		logger.Error("Unsupported PLC. Valid PLC: 'simulator' or 'compact32'")
		return
	}

	exit := make(chan error)
	// PLC work in a completely separate go-routine!
	if plcName == "compact32" {
		driver = compact32.NewCompact32Driver(exit)
	} else {
		driver = simulator.NewSimulator(exit)
	}

	// The exit plan incase there is a feedback from the driver to abort/exit
	go func() {
		err = <-exit
		logger.WithField("err", err.Error()).Error("PLC Driver has requested exi")
		// TODO: Handle exit gracefully
		// We need to call the API on the Web to display the error and restart, abort or call service!
	}()

	store, err = db.Init()
	if err != nil {
		logger.WithField("err", err.Error()).Error("Database init failed")
		return
	}

	deps := service.Dependencies{
		Store: store,
		Plc:   driver,
	}

	// mux router
	router := service.InitRouter(deps)

	// init web server
	server := negroni.Classic()
	server.UseHandler(router)

	port := config.AppPort() // This can be changed to the service port number via environment variable.
	addr := fmt.Sprintf(":%s", strconv.Itoa(port))

	server.Run(addr)
	return
}
