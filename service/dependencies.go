package service

import (
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"mylab/cpagent/tec"
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
