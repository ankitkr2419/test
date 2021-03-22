package compact32

import (
	"fmt"
	"math"
	logger "github.com/sirupsen/logrus"

)


func (d *Compact32Deck) MachineIsHomed() bool{
	if temp, ok := homed.Load(d.name); !ok {
		logger.Errorln("homed isn't loaded!")
	} else if temp.(bool) {
		return true
	}
	return false
}

func (d *Compact32Deck) Homing() (response string, err error) {

	if temp, ok := runInProgress.Load(d.name); !ok {
		err = fmt.Errorf("runInProgress isn't loaded!")
		return
	} else if temp.(bool) {
		err = fmt.Errorf("previous run already in progress... wait or abort it")
		return "", err
	}

	runInProgress.Store(d.name, true)
	defer d.ResetRunInProgress()

	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][5], OFF)
	if err != nil {
		fmt.Println("Inside Switch off Shaker err : ", err, d.name)
		return "", err
	}
	fmt.Println("Switched off the shaker--> for ", d.name)

	aborted.Store(d.name, false)

	fmt.Println("Moving Syringe DOWN till sensor cuts it")
	response, err = d.SyringeHoming()
	if err != nil {
		return
	}

	fmt.Println("Moving Syringe Module UP till sensor cuts it")
	response, err = d.SyringeModuleHoming()
	if err != nil {
		return
	}

	fmt.Println("Homing Magnet")
	response, err = d.MagnetHoming()
	if err != nil {
		return
	}

	fmt.Println("Moving deck forward till sensor cuts it")
	response, err = d.DeckHoming()
	if err != nil {
		return
	}

	homed.Store(d.name, true)

	fmt.Println("Homing Completed Successfully")

	return "HOMING SUCCESS", nil
}

// ***NOTE***
// * In Syringe Sensor is DOWN and not UP.
// * This is exactly opposite of Syringe Module and Magnet Up/Down
// * Thus we need ASPIRE (syringe going UP) and DISPENSE (syringe going DOWN)

func (d *Compact32Deck) SyringeHoming() (response string, err error) {

	deckAndNumber := DeckNumber{Deck: d.name, Number: K10_Syringe_LHRH}

	fmt.Println("Syringe is moving down until sensor not cut")

	response, err = d.SetupMotor(homingFastSpeed, initialSensorCutSyringePulses, motors[deckAndNumber]["ramp"], DISPENSE, deckAndNumber.Number)
	if err != nil {
		return
	}

	fmt.Println("Aspiring and getting cut then aspiring 2000 pulses")
	response, err = d.SetupMotor(homingFastSpeed, reverseAfterNonCutPulses, motors[deckAndNumber]["ramp"], ASPIRE, deckAndNumber.Number)
	if err != nil {
		return
	}

	fmt.Println("Syringe homing is completed")

	return "SYRINGE HOMING COMPLETED", nil
}

func (d *Compact32Deck) SyringeModuleHoming() (response string, err error) {

	deckAndNumber := DeckNumber{Deck: d.name, Number: K9_Syringe_Module_LHRH}

	fmt.Println("Syringe Module moving Up")
	response, err = d.SetupMotor(homingFastSpeed, initialSensorCutSyringeModulePulses, motors[deckAndNumber]["ramp"], UP, deckAndNumber.Number)
	if err != nil {
		return
	}

	fmt.Println("After First Fast Moving Up and getting Cut")

	fmt.Println("Syringe Module moving Down 20 mm or More.")
	response, err = d.SetupMotor(homingFastSpeed, reverseAfterNonCutPulses, motors[deckAndNumber]["ramp"], DOWN, deckAndNumber.Number)
	if err != nil {
		return
	}

	fmt.Println("Syringe Module moving Up")
	response, err = d.SetupMotor(homingSlowSpeed, finalSensorCutPulses, motors[deckAndNumber]["ramp"], UP, deckAndNumber.Number)
	if err != nil {
		return
	}

	fmt.Println("After Final Slow Moving Up and getting Cut")

	return "SYRINGE HOMING SUCCESS", nil
}

func (d *Compact32Deck) DeckHoming() (response string, err error) {

	deckAndNumber := DeckNumber{Deck: d.name, Number: K5_Deck}

	fmt.Println("Deck is moving forward")
	response, err = d.SetupMotor(homingDeckFastSpeed, initialSensorCutDeckPulses, motors[deckAndNumber]["ramp"], FWD, deckAndNumber.Number)
	if err != nil {
		return
	}

	fmt.Println("Deck is moving back by and after not cut -> 2000")
	response, err = d.SetupMotor(homingDeckFastSpeed, reverseAfterNonCutPulses, motors[deckAndNumber]["ramp"], REV, deckAndNumber.Number)
	if err != nil {
		return
	}

	fmt.Println("Deck is moving forward again by 2999")
	response, err = d.SetupMotor(homingSlowSpeed, finalSensorCutPulses, motors[deckAndNumber]["ramp"], FWD, deckAndNumber.Number)
	if err != nil {
		return
	}

	fmt.Println("Deck homing is completed.")

	return "DECK HOMING SUCCESS", nil
}

func (d *Compact32Deck) MagnetHoming() (response string, err error) {
	var magnetDetach float64
	var ok bool
	var pulses uint16
	deckAndNumber := DeckNumber{Deck: d.name, Number: K7_Magnet_Rev_Fwd}

	// Detaching magnet, doesn't matter even if its already detached
	if magnetDetach, ok = consDistance["magnet_detach_for_homing"]; !ok {
		err = fmt.Errorf("magnet_detach_for_homing doesn't exist")
		fmt.Println("Error: ", err, d.name)
		return "", err
	}
	fmt.Println("Magnet is moving backward by 5 mm for detachment")

	pulses = uint16(math.Round(float64(motors[deckAndNumber]["steps"]) * magnetDetach))

	response, err = d.SetupMotor(motors[deckAndNumber]["fast"], pulses, motors[deckAndNumber]["ramp"], REV, deckAndNumber.Number)
	if err != nil {
		return
	}

	response, err = d.MagnetUpDownHoming()
	if err != nil {
		return
	}
	response, err = d.MagnetFwdRevHoming()
	if err != nil {
		return
	}

	return "MAGNET HOMING SUCCESS", nil
}

func (d *Compact32Deck) MagnetUpDownHoming() (response string, err error) {

	deckAndNumber := DeckNumber{Deck: d.name, Number: K6_Magnet_Up_Down}

	fmt.Println("Magnet is moving up")
	response, err = d.SetupMotor(homingFastSpeed, initialSensorCutMagnetPulses, motors[deckAndNumber]["ramp"], UP, deckAndNumber.Number)
	if err != nil {
		return
	}

	// NOTE: Less Pulses used as 2000 cause magnet dash onto 1000ul tips at worst conditions.
	fmt.Println("Magnet is moving down by and after not cut -> 400")
	response, err = d.SetupMotor(homingFastSpeed, reverseAfterNonCutPulsesMagnet, motors[deckAndNumber]["ramp"], DOWN, deckAndNumber.Number)
	if err != nil {
		return
	}

	fmt.Println("Magnet is moving up again by 2999 till sensor cuts")
	response, err = d.SetupMotor(homingSlowSpeed, finalSensorCutPulses, motors[deckAndNumber]["ramp"], UP, deckAndNumber.Number)

	fmt.Println("Magnet Up/Down homing is completed.")

	return "MAGNET UP/DOWN HOMING SUCCESS", nil
}

func (d *Compact32Deck) MagnetFwdRevHoming() (response string, err error) {

	deckAndNumber := DeckNumber{Deck: d.name, Number: K7_Magnet_Rev_Fwd}
	var magnetReverseAfterHoming, distanceToTravel float64
	var pulses uint16
	var ok bool

	fmt.Println("Magnet is moving forward")
	response, err = d.SetupMotor(homingFastSpeed, initialSensorCutMagnetPulses, motors[deckAndNumber]["ramp"], FWD, deckAndNumber.Number)
	if err != nil {
		return
	}

	fmt.Println("Magnet is moving back by and after not cut -> 2000")
	response, err = d.SetupMotor(homingFastSpeed, reverseAfterNonCutPulses, motors[deckAndNumber]["ramp"], REV, deckAndNumber.Number)
	if err != nil {
		return
	}

	fmt.Println("Magnet is moving forward again by 2999")
	response, err = d.SetupMotor(homingSlowSpeed, finalSensorCutPulses, motors[deckAndNumber]["ramp"], FWD, deckAndNumber.Number)

	fmt.Println("Moving Magnet Back by 50mm")

	if magnetReverseAfterHoming, ok = consDistance["magnet_reverse_after_homing"]; !ok {
		err = fmt.Errorf("magnet_reverse_after_homing doesn't exist")
		fmt.Println("Error: ", err, d.name)
		return "", err
	}

	// We know the concrete direction here, its reverse
	distanceToTravel = magnetReverseAfterHoming - positions[deckAndNumber]
	fmt.Println("Magnet Pos:---> ", positions[deckAndNumber])
	// Make Travel Distance Positive if was negative
	if distanceToTravel < 0 {
		distanceToTravel *= -1
	}
	pulses = uint16(math.Round(float64(motors[deckAndNumber]["steps"]) * distanceToTravel))

	response, err = d.SetupMotor(homingFastSpeed, pulses, motors[deckAndNumber]["ramp"], REV, deckAndNumber.Number)
	if err != nil {
		return
	}

	fmt.Println("Magnet Fwd/Rev homing is completed with reverse pulses added.")

	return "MAGNET FWD/REV HOMING SUCCESS", nil
}