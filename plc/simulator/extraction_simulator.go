package simulator

import (
	"mylab/cpagent/plc"
)

func NewExtractionSimulator(wsMsgch chan string, wsErrch chan error, exit chan error, deck string) plc.Common {
	s := plc.Compact32Deck{}
	s.WsMsgCh = wsMsgch

	driver := SimulatorDriver{}

	s.DeckDriver = &driver
	s.ExitCh = exit
	s.WsMsgCh = wsMsgch
	s.WsErrCh = wsErrch

	s.Name = deck

	return &s
}