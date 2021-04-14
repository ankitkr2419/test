package service

import (
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
)

type Dependencies struct {
	Store   db.Storer
	Plc     plc.Driver
	PlcDeck map[string]plc.Common
	ExitCh  <-chan error
	WsErrCh chan error
	WsMsgCh chan string
	// define other service dependencies
}
