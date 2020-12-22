package main

import (
	rice "github.com/GeertJohan/go.rice"

	"flag"
	"mylab/cpagent/config"
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"mylab/cpagent/plc/compact32"
	"mylab/cpagent/plc/simulator"
	"mylab/cpagent/service"
	"net/http"
	"os"
	"strconv"

	"github.com/rs/cors"
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

	// config file to configure dye & targets
	config.Load("config")

	// simulator config file to configure controls & wells in simulator
	config.Load("simulator")

	// config file to configure motors
	config.Load("motor_config")

	// config file to configure consumable distance
	config.Load("consumable_config")

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
				&cli.BoolFlag{
					Name:  "test",
					Usage: "Run in test mode!",
				},
			},
			Action: func(c *cli.Context) error {
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
	}

	if err := cliApp.Run(os.Args); err != nil {
		panic(err)
	}
}

func startApp(plcName string, test bool) (err error) {
	var store db.Storer
	var driver plc.Driver

	if plcName != "simulator" && plcName != "compact32" {
		logger.Error("Unsupported PLC. Valid PLC: 'simulator' or 'compact32'")
		return
	}

	exit := make(chan error)

	websocketMsg := make(chan string)

	websocketErr := make(chan error)

	// PLC work in a completely separate go-routine!
	if plcName == "compact32" {
		driver = compact32.NewCompact32Driver(exit, test)
	} else {
		driver = simulator.NewSimulator(exit)
	}

	store, err = db.Init()
	if err != nil {
		logger.WithField("err", err.Error()).Error("Database init failed")
		return
	}

	deps := service.Dependencies{
		Store:   store,
		Plc:     driver,
		ExitCh:  exit,
		WsErrCh: websocketErr,
		WsMsgCh: websocketMsg,
	}

	// setup Db with dyes & targets
	err = db.Setup(store)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Setup Dyes & Targets failed")
		return
	}

	// setup Db with motors
	err = db.SetupMotor(store)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Setup Motors failed")
		return
	}

	// setup Db with consumable distance
	err = db.SetupConsumable(store)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Setup Cosumable Distance failed")
		return
	}

	// add default User
	u := db.User{
		Username: "admin",
		Password: service.MD5Hash("admin"),
		Role:     "administrator",
	}
	db.AddDefaultUser(store, u)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Setup Default User failed")
		return
	}

	var addr = flag.String("addr", "localhost:"+strconv.Itoa(config.AppPort()), "http service address")
	// mux router
	router := service.InitRouter(deps)

	// to embed react build with go rice
	router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("./web-client/build").HTTPBox()))

	// cors configuration for front-end
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"PUT", "DELETE", "POST", "GET"},
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
