package simulator

import (
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"sync"
	"time"
)

type UISimulator struct {
	sync.RWMutex
	name   string // Deck Name
	ExitCh chan string
	ErrCh  chan error
}

func NewUISimulator(exit chan error, deck string) plc.DeckDriver {
	ex := make(chan string)
	s := UISimulator{}
	s.ExitCh = ex
	s.ErrCh = exit
	s.name = deck
	return &s
}

// NameOfDeck returns name of deck
func (us *UISimulator) NameOfDeck() string {
	return us.name
}

// Homing returns success if the machine is  successfully homed
func (us *UISimulator) Homing() (string, error) { return "SUCCESS", nil }

// DeckHoming
func (us *UISimulator) DeckHoming() (string, error)          { return "SUCCESS", nil }
func (us *UISimulator) SyringeHoming() (string, error)       { return "SUCCESS", nil }
func (us *UISimulator) SyringeModuleHoming() (string, error) { return "SUCCESS", nil }
func (us *UISimulator) MagnetHoming() (string, error)        { return "SUCCESS", nil }
func (us *UISimulator) MagnetUpDownHoming() (string, error)  { return "SUCCESS", nil }
func (us *UISimulator) MagnetFwdRevHoming() (string, error)  { return "SUCCESS", nil }
func (us *UISimulator) SwitchOffMotor() (string, error)      { return "SUCCESS", nil }
func (us *UISimulator) ReadExecutedPulses() (string, error)  { return "SUCCESS", nil }
func (us *UISimulator) SetupMotor(uint16, uint16, uint16, uint16, uint16) (string, error) {
	return "SUCCESS", nil
}
func (us *UISimulator) ManualMovement(uint16, uint16, uint16) (string, error) { return "SUCCESS", nil }
func (us *UISimulator) ResetRunInProgress()                                   { return }
func (us *UISimulator) SetRunInProgress()                                     { return }
func (us *UISimulator) SetTimerInProgress()                                   { return }
func (us *UISimulator) ResetTimerInProgress()                                 { return }
func (us *UISimulator) Pause() (string, error)                                { return "SUCCESS", nil }
func (us *UISimulator) Resume() (string, error)                               { return "SUCCESS", nil }
func (us *UISimulator) Abort() (string, error)                                { return "SUCCESS", nil }
func (us *UISimulator) DiscardBoxCleanup() (string, error)                    { return "SUCCESS", nil }
func (us *UISimulator) RestoreDeck() (string, error)                          { return "SUCCESS", nil }
func (us *UISimulator) UVLight(string) (string, error)                        { return "SUCCESS", nil }
func (us *UISimulator) ResumeMotorWithPulses(uint16) (string, error)          { return "SUCCESS", nil }
func (us *UISimulator) Heating(uint16, bool, time.Duration) (string, error)   { return "SUCCESS", nil }
func (us *UISimulator) AspireDispense(aspireDispense db.AspireDispense, cartridgeID int64, tipType string) (response string, err error) {
	return "SUCCESS", nil
}
func (us *UISimulator) TipDocking(td db.TipDock, cartridgeID int64) (response string, err error) {
	return "SUCCESS", nil
}
func (us *UISimulator) TipOperation(to db.TipOperation) (response string, err error) { return "SUCCESS", nil }
func (us *UISimulator) TipPickup(pos int64) (response string, err error)             { return "SUCCESS", nil }
func (us *UISimulator) TipDiscard() (response string, err error)                     { return "SUCCESS", nil }
func (us *UISimulator) AttachDetach(db.AttachDetach) (response string, err error)    { return "SUCCESS", nil }
func (us *UISimulator) AddDelay(db.Delay) (string, error)                            { return "SUCCESS", nil }
func (us *UISimulator) Piercing(pi db.Piercing, cartridgeID int64) (response string, err error) {
	return "SUCCESS", nil
}
