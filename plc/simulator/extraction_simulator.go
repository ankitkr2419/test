package simulator

import (
	"encoding/json"
	"fmt"
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"sync"
	"time"
)

type ExtractionSimulator struct {
	sync.RWMutex
	name    string // Deck Name
	WsMsgCh chan string
	ExitCh  chan string
	ErrCh   chan error
}

func NewExtractionSimulator(wsMsgch chan string, exit chan error, deck string) plc.DeckDriver {
	ex := make(chan string)
	s := ExtractionSimulator{}
	s.WsMsgCh = wsMsgch

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

func (us *ExtractionSimulator) Homing() (string, error) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 1)

		wsProgressOperation := plc.WSData{
			Progress: float64(10 * i),
			Deck:     us.name,
			Status:   "PROGRESS_HOMING",
			OperationDetails: plc.OperationDetails{
				Message: fmt.Sprintf("successfully homed %v for deck %v", i, us.name),
			},
		}

		wsData, err := json.Marshal(wsProgressOperation)
		if err != nil {
			us.ErrCh <- err
		}
		us.WsMsgCh <- fmt.Sprintf("progress_homing_%v", string(wsData))
	}
	us.WsMsgCh <- fmt.Sprintf("success_homing_successfully homed for deck %v", us.name)
	return "SUCCESS", nil
}

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
func (us *ExtractionSimulator) UVLight(string) (string, error) {

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 1)
		// the format for message in every operation must be in the format:
		// progress/success_OPERATION NAME_OPERATION MESSAGE
		us.WsMsgCh <- fmt.Sprintf("progress_uvLight_uv light cleanup in progress for deck %s ", us.name)
	}
	us.WsMsgCh <- fmt.Sprintf("success_uvLight_uv Light Completed Successfully %v", us.name)
	return "SUCCESS", nil

	return "SUCCESS", nil
}
func (us *ExtractionSimulator) Heating(db.Heating) (string, error) {
	time.Sleep(time.Second * 1)

	return "SUCCESS", nil
}
func (us *ExtractionSimulator) AspireDispense(aspireDispense db.AspireDispense, cartridgeID int64, tipType string) (response string, err error) {
	time.Sleep(time.Second * 1)

	return "SUCCESS", nil
}
func (us *ExtractionSimulator) TipDocking(td db.TipDock, cartridgeID int64) (response string, err error) {
	return "SUCCESS", nil
}
func (us *ExtractionSimulator) TipOperation(to db.TipOperation) (response string, err error) {
	time.Sleep(time.Second * 1)

	return "SUCCESS", nil
}
func (us *ExtractionSimulator) TipPickup(pos int64) (response string, err error) {
	time.Sleep(time.Second * 1)

	return "SUCCESS", nil
}
func (us *ExtractionSimulator) TipDiscard() (response string, err error) {
	time.Sleep(time.Second * 1)

	return "SUCCESS", nil
}
func (us *ExtractionSimulator) AttachDetach(db.AttachDetach) (response string, err error) {
	time.Sleep(time.Second * 1)

	return "SUCCESS", nil
}
func (us *ExtractionSimulator) AddDelay(db.Delay) (string, error) {
	time.Sleep(time.Second * 1)

	return "SUCCESS", nil
}
func (us *ExtractionSimulator) Piercing(pi db.Piercing, cartridgeID int64) (response string, err error) {
	time.Sleep(time.Second * 1)

	return "SUCCESS", nil
}

func (us *ExtractionSimulator) DiscardTipAndHome(bool) (response string, err error) {
	time.Sleep(time.Second * 1)

	return "SUCCESS", nil
}
func (us *ExtractionSimulator) Shaking(db.Shaker) (string, error) {
	return "SUCCESS", nil

}
