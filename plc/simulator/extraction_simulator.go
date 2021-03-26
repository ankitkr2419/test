package simulator

import (
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"sync"
	"time"
)

type ExtractionSimulator struct {
	sync.RWMutex
	name   string // Deck Name
	ExitCh chan string
	ErrCh  chan error
}

func NewExtractionSimulator(exit chan error, deck string) plc.DeckDriver {
	ex := make(chan string)
	s := ExtractionSimulator{}
	s.ExitCh = ex
	s.ErrCh = exit
	s.name = deck
	return &s
}

// NameOfDeck returns name of deck
func (us *ExtractionSimulator) NameOfDeck() string {
	return us.name
}

// Homing returns success if the machine is  successfully homed
func (us *ExtractionSimulator) Homing() (string, error) { return "SUCCESS", nil }
func (us *ExtractionSimulator) ManualMovement(uint16, uint16, uint16) (string, error) {
	return "SUCCESS", nil
}
func (us *ExtractionSimulator) IsMachineHomed() bool               { return true }
func (us *ExtractionSimulator) IsRunInProgress() bool              { return false }
func (us *ExtractionSimulator) ResetRunInProgress()                { return }
func (us *ExtractionSimulator) SetRunInProgress()                  { return }
func (us *ExtractionSimulator) Pause() (string, error)             { return "SUCCESS", nil }
func (us *ExtractionSimulator) Resume() (string, error)            { return "SUCCESS", nil }
func (us *ExtractionSimulator) Abort() (string, error)             { return "SUCCESS", nil }
func (us *ExtractionSimulator) DiscardBoxCleanup() (string, error) { return "SUCCESS", nil }
func (us *ExtractionSimulator) RestoreDeck() (string, error)       { return "SUCCESS", nil }
func (us *ExtractionSimulator) UVLight(string) (string, error)     { return "SUCCESS", nil }
func (us *ExtractionSimulator) Heating(uint16, bool, time.Duration) (string, error) {
	return "SUCCESS", nil
}
func (us *ExtractionSimulator) AspireDispense(aspireDispense db.AspireDispense, cartridgeID int64, tipType string) (response string, err error) {
	return "SUCCESS", nil
}
func (us *ExtractionSimulator) TipDocking(td db.TipDock, cartridgeID int64) (response string, err error) {
	return "SUCCESS", nil
}
func (us *ExtractionSimulator) TipOperation(to db.TipOperation) (response string, err error) {
	return "SUCCESS", nil
}
func (us *ExtractionSimulator) TipPickup(pos int64) (response string, err error) {
	return "SUCCESS", nil
}
func (us *ExtractionSimulator) TipDiscard() (response string, err error) { return "SUCCESS", nil }
func (us *ExtractionSimulator) AttachDetach(db.AttachDetach) (response string, err error) {
	return "SUCCESS", nil
}
func (us *ExtractionSimulator) AddDelay(db.Delay) (string, error) { return "SUCCESS", nil }
func (us *ExtractionSimulator) Piercing(pi db.Piercing, cartridgeID int64) (response string, err error) {
	return "SUCCESS", nil
}