package main

// @APITitle Main
// @APIDescription Main API for Microservices in Go!

import (
	"errors"
	"fmt"
	"mylab/mylabdiscoveries/config"
	"mylab/mylabdiscoveries/db"
	"mylab/mylabdiscoveries/service"
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
			Usage: "start server",
			Action: func(c *cli.Context) error {
				return startApp(c.Args().Get(0))
			},
		},
		{
			Name:  "create_migration",
			Usage: "create migration file",
			Action: func(c *cli.Context) error {
				switch c.Args().Get(1) {
				case "dev":
					return db.CreateMigrationFileSQLite(c.Args().Get(0))
				case "prod":
					return db.CreateMigrationFile(c.Args().Get(0))
				default:
					logger.WithField("mode", c.Args().Get(1)).Info("Select valid mode")
					return errors.New("Invalid Mode")
				}
			},
		},
		{
			Name:  "migrate",
			Usage: "run db migrations",
			Action: func(c *cli.Context) error {
				switch c.Args().Get(0) {
				case "dev":
					return db.RunMigrationsSQLite()
				case "prod":
					return db.RunMigrations()
				default:
					logger.WithField("mode", c.Args().Get(1)).Info("Select valid mode")
					return errors.New("Invalid Mode")
				}
			},
		},
		{
			Name:  "rollback",
			Usage: "rollback migrations",
			Action: func(c *cli.Context) error {
				switch c.Args().Get(1) {
				case "dev":
					return db.RollbackMigrationsSQLite(c.Args().Get(0))
				case "prod":
					return db.RollbackMigrations(c.Args().Get(0))
				default:
					logger.WithField("mode", c.Args().Get(1)).Info("Select valid mode")
					return errors.New("Invalid Mode")
				}

			},
		},
	}

	if err := cliApp.Run(os.Args); err != nil {
		panic(err)
	}
}

func startApp(Mode string) (err error) {
	var store db.Storer
	switch Mode {
	case "dev":
		store, err = db.InitSQL()
		if err != nil {
			logger.WithField("err", err.Error()).Error("Database init failed")
			return
		}
	case "prod":
		store, err = db.Init()
		if err != nil {
			logger.WithField("err", err.Error()).Error("Database init failed")
			return
		}
	default:
		logger.WithField("mode", Mode).Info("Select valid mode")
		return
	}

	deps := service.Dependencies{
		Store: store,
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
