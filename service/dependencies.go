package service

import (
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
)

type Dependencies struct {
	Store  db.Storer
	Plc    plc.Driver
	ExitCh <-chan error
	// define other service dependencies
}
