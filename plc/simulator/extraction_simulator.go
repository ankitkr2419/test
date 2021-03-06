package simulator

import (
	"mylab/cpagent/plc"
)

func NewExtractionSimulator(wsMsgch chan string, wsErrch chan error, exit chan error, deck string) plc.Extraction {
	s := plc.Compact32Deck{}
	s.WsMsgCh = wsMsgch

	driver := SimulatorDriver{DeckName: deck}

	s.DeckDriver = &driver
	s.ExitCh = exit
	s.WsMsgCh = wsMsgch
	s.WsErrCh = wsErrch

	plc.SetDeckName(&s, deck)

	go loadUtils()
	return &s
}
