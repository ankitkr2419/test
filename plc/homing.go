package plc

import (
	"encoding/json"
	"fmt"
	"mylab/cpagent/responses"
	
	"math"
	"time"

	logger "github.com/sirupsen/logrus"
)

func (d *Compact32Deck) Homing() (response string, err error) {

	d.setHomingPercent(0)

	if d.IsRunInProgress() {
		err = responses.PreviousRunInProgressError
		return "", err
	}

	defer func() {
		if r := recover(); r != nil {
			time.Sleep(2 * time.Second)
			d.ResetRunInProgress()
			logger.Infof("\nRecovering in Homing %v for Deck %v", r, d.name)

			time.Sleep(2 * time.Second)
			if !d.IsMachineHomed() {
				d.ExitCh <- fmt.Errorf("%v_%v_%v", ErrorExtractionMonitor, d.name, err.Error())
				time.Sleep(5 * time.Second)
				response, err = d.Homing()
			}
		}
	}()

	d.SetRunInProgress()
	defer d.ResetRunInProgress()

	d.ResetAborted()

	logger.Infoln("Moving Syringe DOWN till sensor cuts it")
	response, err = d.syringeHoming()
	if err != nil {
		panic(err)
	}

	//web socket response for syringe homing
	d.setHomingPercent(25)

	// NOTE: getHomingPercent will handle both Deck homing percent
	// Similarly getHomingDeckName will handle both Deck Name convention
	wsProgressOperation := WSData{
		Progress: d.getHomingPercent(),
		Deck:     d.getHomingDeckName(),
		Status:   "PROGRESS_HOMING",
		OperationDetails: OperationDetails{
			Message: fmt.Sprintf("successfully homed syringe for deck %v", d.name),
		},
	}

	wsData, err := json.Marshal(wsProgressOperation)
	if err != nil {
		logger.Errorf("error in marshalling web socket data %v", err.Error())
		d.WsErrCh <- fmt.Errorf("%v_%v_%v", ErrorExtractionMonitor, d.name, err.Error())
	}
	d.WsMsgCh <- fmt.Sprintf("progress_homing_%v", string(wsData))

	logger.Infoln("Moving Syringe Module UP till sensor cuts it")
	response, err = d.syringeModuleHoming()
	if err != nil {
		panic(err)
	}

	//web socket response for syringe module homing
	d.setHomingPercent(50.0)

	wsProgressOperation.Progress = d.getHomingPercent()
	wsProgressOperation.OperationDetails.Message = fmt.Sprintf("successfully homed syringe module for deck %v", d.name)

	wsData, err = json.Marshal(wsProgressOperation)
	if err != nil {
		logger.Errorf("error in marshalling web socket data %v", err.Error())
		d.WsErrCh <- fmt.Errorf("%v_%v_%v", ErrorExtractionMonitor, d.name, err.Error())
	}

	d.WsMsgCh <- fmt.Sprintf("progress_homing_%v", string(wsData))

	logger.Infoln("Homing Magnet")
	response, err = d.magnetHoming()
	if err != nil {
		panic(err)
	}

	//web socket response for magnet homing

	d.setHomingPercent(75.0)

	wsProgressOperation.Progress = d.getHomingPercent()
	wsProgressOperation.OperationDetails.Message = fmt.Sprintf("successfully homed magnet for deck %v", d.name)
	wsData, err = json.Marshal(wsProgressOperation)
	if err != nil {
		logger.Errorf("error in marshalling web socket data %v", err.Error())
		d.WsErrCh <- fmt.Errorf("%v_%v_%v", ErrorExtractionMonitor, d.name, err.Error())
	}

	d.WsMsgCh <- fmt.Sprintf("progress_homing_%v", string(wsData))

	// Started Deck Homing
	logger.Infoln("Moving deck forward till sensor cuts it")
	response, err = d.deckHoming()
	if err != nil {
		panic(err)
	}

	d.setHomingPercent(100.0)

	wsProgressOperation.Progress = d.getHomingPercent()
	wsProgressOperation.OperationDetails.Message = fmt.Sprintf("successfully homed deck %v", d.name)

	wsData, err = json.Marshal(wsProgressOperation)
	if err != nil {
		logger.Errorf("error in marshalling web socket data %v", err.Error())
		d.WsErrCh <- fmt.Errorf("%v_%v_%v", ErrorExtractionMonitor, d.name, err.Error())
	}
	d.WsMsgCh <- fmt.Sprintf("progress_homing_%v", string(wsData))

	d.setHomed()

	successWsData := WSData{
		Progress: 100,
		Deck:     d.name,
		Status:   "SUCCESS_HOMING",
		OperationDetails: OperationDetails{
			Message: fmt.Sprintf("successfully homed for deck %v", d.name),
		},
	}
	wsData, err = json.Marshal(successWsData)
	if err != nil {
		logger.Errorf("error in marshalling web socket data %v", err.Error())
		d.WsErrCh <- fmt.Errorf("%v_%v_%v", ErrorExtractionMonitor, d.name, err.Error())
		return "", err
	}
	d.WsMsgCh <- fmt.Sprintf("success_homing_%v", string(wsData))

	logger.Infoln("Homing Completed Successfully")

	return "HOMING SUCCESS", nil
}

// ***NOTE***
// * In Syringe Sensor is DOWN and not UP.
// * This is exactly opposite of Syringe Module and Magnet Up/Down
// * Thus we need ASPIRE (syringe going UP) and DISPENSE (syringe going DOWN)

func (d *Compact32Deck) syringeHoming() (response string, err error) {

	deckAndNumber := DeckNumber{Deck: d.name, Number: K10_Syringe_LHRH}

	logger.Infoln("Syringe is moving down until sensor not cut")

	response, err = d.setupMotor(homingFastSpeed, initialSensorCutSyringePulses, Motors[deckAndNumber]["ramp"], DISPENSE, deckAndNumber.Number)
	if err != nil {
		return
	}

	logger.Infoln("Aspiring and getting cut then aspiring 2000 pulses")
	response, err = d.setupMotor(homingFastSpeed, reverseAfterNonCutPulses, Motors[deckAndNumber]["ramp"], ASPIRE, deckAndNumber.Number)
	if err != nil {
		return
	}

	logger.Infoln("Dispensing until sensor cut")
	response, err = d.setupMotor(homingSlowSpeed, initialSensorCutSyringePulses, Motors[deckAndNumber]["ramp"], DISPENSE, deckAndNumber.Number)
	if err != nil {
		return
	}

	logger.Infoln("Syringe homing is completed")

	return "SYRINGE HOMING COMPLETED", nil
}

func (d *Compact32Deck) syringeModuleHoming() (response string, err error) {

	deckAndNumber := DeckNumber{Deck: d.name, Number: K9_Syringe_Module_LHRH}

	logger.Infoln("Syringe Module moving Up")
	response, err = d.setupMotor(homingFastSpeed, initialSensorCutSyringeModulePulses, Motors[deckAndNumber]["ramp"], UP, deckAndNumber.Number)
	if err != nil {
		return
	}

	logger.Infoln("After First Fast Moving Up and getting Cut")

	logger.Infoln("Syringe Module moving Down 20 mm or More.")
	response, err = d.setupMotor(homingFastSpeed, reverseAfterNonCutPulses, Motors[deckAndNumber]["ramp"], DOWN, deckAndNumber.Number)
	if err != nil {
		return
	}

	logger.Infoln("Syringe Module moving Up")
	response, err = d.setupMotor(homingSlowSpeed, finalSensorCutPulses, Motors[deckAndNumber]["ramp"], UP, deckAndNumber.Number)
	if err != nil {
		return
	}

	logger.Infoln("After Final Slow Moving Up and getting Cut")

	return "SYRINGE HOMING SUCCESS", nil
}

func (d *Compact32Deck) deckHoming() (response string, err error) {

	deckAndNumber := DeckNumber{Deck: d.name, Number: K5_Deck}

	logger.Infoln("Deck is moving forward")
	response, err = d.setupMotor(homingDeckFastSpeed, initialSensorCutDeckPulses, Motors[deckAndNumber]["ramp"], FWD, deckAndNumber.Number)
	if err != nil {
		return
	}

	logger.Infoln("Deck is moving back by and after not cut -> 2000")
	response, err = d.setupMotor(homingDeckFastSpeed, reverseAfterNonCutPulses, Motors[deckAndNumber]["ramp"], REV, deckAndNumber.Number)
	if err != nil {
		return
	}

	logger.Infoln("Deck is moving forward again by 2999")
	response, err = d.setupMotor(homingSlowSpeed, finalSensorCutPulses, Motors[deckAndNumber]["ramp"], FWD, deckAndNumber.Number)
	if err != nil {
		return
	}

	logger.Infoln("Deck homing is completed.")

	return "DECK HOMING SUCCESS", nil
}

func (d *Compact32Deck) magnetHoming() (response string, err error) {
	var magnetDetach float64
	var ok bool
	var pulses uint16
	deckAndNumber := DeckNumber{Deck: d.name, Number: K7_Magnet_Rev_Fwd}

	// Detaching magnet, doesn't matter even if its already detached
	if magnetDetach, ok = consDistance["magnet_detach_for_homing"]; !ok {
		err = fmt.Errorf("magnet_detach_for_homing doesn't exist")
		logger.Errorln(err, d.name)
		return "", err
	}
	logger.Infoln("Magnet is moving backward by 5 mm for detachment")

	pulses = uint16(math.Round(float64(Motors[deckAndNumber]["steps"]) * magnetDetach))

	response, err = d.setupMotor(Motors[deckAndNumber]["fast"], pulses, Motors[deckAndNumber]["ramp"], REV, deckAndNumber.Number)
	if err != nil {
		return
	}

	response, err = d.magnetUpDownHoming()
	if err != nil {
		return
	}
	response, err = d.magnetFwdRevHoming()
	if err != nil {
		return
	}

	return "MAGNET HOMING SUCCESS", nil
}

func (d *Compact32Deck) magnetUpDownHoming() (response string, err error) {

	deckAndNumber := DeckNumber{Deck: d.name, Number: K6_Magnet_Up_Down}

	logger.Infoln("Magnet is moving up")
	response, err = d.setupMotor(homingFastSpeed, initialSensorCutMagnetPulses, Motors[deckAndNumber]["ramp"], UP, deckAndNumber.Number)
	if err != nil {
		return
	}

	// NOTE: Less Pulses used as 2000 cause magnet dash onto 1000ul tips at worst conditions.
	logger.Infoln("Magnet is moving down by and after not cut -> 400")
	response, err = d.setupMotor(homingFastSpeed, reverseAfterNonCutPulsesMagnet, Motors[deckAndNumber]["ramp"], DOWN, deckAndNumber.Number)
	if err != nil {
		return
	}

	logger.Infoln("Magnet is moving up again by 2999 till sensor cuts")
	response, err = d.setupMotor(homingSlowSpeed, finalSensorCutPulses, Motors[deckAndNumber]["ramp"], UP, deckAndNumber.Number)

	logger.Infoln("Magnet Up/Down homing is completed.")

	return "MAGNET UP/DOWN HOMING SUCCESS", nil
}

func (d *Compact32Deck) magnetFwdRevHoming() (response string, err error) {

	deckAndNumber := DeckNumber{Deck: d.name, Number: K7_Magnet_Rev_Fwd}
	var magnetReverseAfterHoming, distanceToTravel float64
	var pulses uint16
	var ok bool

	logger.Infoln("Magnet is moving forward")
	response, err = d.setupMotor(homingFastSpeed, initialSensorCutMagnetPulses, Motors[deckAndNumber]["ramp"], FWD, deckAndNumber.Number)
	if err != nil {
		return
	}

	logger.Infoln("Magnet is moving back by and after not cut -> 2000")
	response, err = d.setupMotor(homingFastSpeed, reverseAfterNonCutPulses, Motors[deckAndNumber]["ramp"], REV, deckAndNumber.Number)
	if err != nil {
		return
	}

	logger.Infoln("Magnet is moving forward again by 2999")
	response, err = d.setupMotor(homingSlowSpeed, finalSensorCutPulses, Motors[deckAndNumber]["ramp"], FWD, deckAndNumber.Number)

	logger.Infoln("Moving Magnet Back by 50mm")

	if magnetReverseAfterHoming, ok = consDistance["magnet_reverse_after_homing"]; !ok {
		err = fmt.Errorf("magnet_reverse_after_homing doesn't exist")
		logger.Errorln(err, d.name)
		return "", err
	}

	// We know the concrete direction here, its reverse
	distanceToTravel = magnetReverseAfterHoming - Positions[deckAndNumber]
	logger.Infoln("Magnet Pos:---> ", Positions[deckAndNumber])
	// Make Travel Distance Positive if was negative
	if distanceToTravel < 0 {
		distanceToTravel *= -1
	}
	pulses = uint16(math.Round(float64(Motors[deckAndNumber]["steps"]) * distanceToTravel))

	response, err = d.setupMotor(homingFastSpeed, pulses, Motors[deckAndNumber]["ramp"], REV, deckAndNumber.Number)
	if err != nil {
		return
	}

	logger.Infoln("Magnet Fwd/Rev homing is completed with reverse pulses added.")

	return "MAGNET FWD/REV HOMING SUCCESS", nil
}
