package compact32

import (
	"fmt"
	"math"
	"mylab/cpagent/db"
	"time"
)

/****ALGORITHM******
1. Call Tip Pick Up at position 3
2. Pierce well after well
3. Call Tip Discard

********/

func (d *Compact32Deck) Piercing(pi db.Piercing, cartridgeID int64) (response string, err error) {

	// Move Deck to Extraction Cartridge start
	distToTravel = cd["extraction_cartridge_start"] - at_X
	response, err = m.setUpMotor(uint16(motors[0].Fast), uint16(math.Round(float64(motors[0].Steps)*distToTravel)), uint16(motors[0].Ramp), uint16(Fwd), uint16(motors[0].Number), On, Off)
	if err != nil {
		return
	}

	at_X = cd["extraction_cartridge_start"]

	fmt.Println("Completed Move Deck to Extraction Cartridge start")
	time.Sleep(1 * time.Second)

	//	wellsDistanceInMM := []float64{17.5, 24.5, 33.3, 41.2, 51.2, 62.99, 74.9, 84.3}

	//*************************
	// Pierce Well after Well *
	//*************************
	//	for wellNum, distance := range wellsDistanceInMM {
	// Well Distances -> from "extraction_cartridge_start"
	for _, c := range cartriage {

		distToTravel = c.Distance - at_X
		// Move Deck to reach the well
		response, err = m.setUpMotor(uint16(motors[0].Fast), uint16(math.Round(float64(motors[0].Steps)*distToTravel)), uint16(motors[0].Ramp), uint16(Fwd), uint16(motors[0].Number), On, Off)
		if err != nil {
			return
		}

		at_X = cd["extraction_cartridge_start"] + c.Distance

		fmt.Println("Completed Move Deck to reach the wellNum ", c.WellNum)
		time.Sleep(1 * time.Second)

		// Pierce the Well
		// Move Syringe Module Down Fwd 31mm
		response, err = m.setUpMotor(uint16(motors[1].Fast), uint16(math.Round(float64(motors[1].Steps)*cd["piercing_height"])), uint16(motors[1].Ramp), uint16(Fwd), uint16(motors[1].Number), On, Off)
		if err != nil {
			return
		}

		fmt.Println("Completed Move Syringe Module Down Fwd 31mm wellNum ", c.WellNum)
		time.Sleep(1 * time.Second)

		// Move Syringe Module Up Rev Fwd 31mm
		response, err = m.setUpMotor(uint16(motors[1].Fast), uint16(math.Round(float64(motors[1].Steps)*cd["piercing_height"])), uint16(motors[1].Ramp), uint16(Rev), uint16(motors[1].Number), On, Off)
		if err != nil {
			return
		}

		fmt.Println("Completed Move Syringe Module Up Rev 31mm wellNum ", c.WellNum)
		time.Sleep(1 * time.Second)

	}

	//**************
	// Tip Discard *
	//**************

	// Move Deck Fwd 150.54mm

	distToTravel = cd["discard_big_hole"] - at_X
	response, err = m.setUpMotor(uint16(motors[0].Fast), uint16(math.Round(float64(motors[0].Steps)*distToTravel)), uint16(motors[0].Ramp), uint16(Fwd), uint16(motors[0].Number), On, Off)
	if err != nil {
		return
	}

	at_X = cd["discard_big_hole"]

	fmt.Println("Completed Move Deck Fwd 150.54mm")
	time.Sleep(1 * time.Second)

	// Move Syringe Module Down Fwd 83.9mm
	response, err = m.setUpMotor(uint16(motors[1].Fast), uint16(math.Round(float64(motors[1].Steps)*cd["piercing_tip_discard_height"])), uint16(motors[1].Ramp), uint16(Rev), uint16(motors[1].Number), On, Off)
	if err != nil {
		return
	}

	fmt.Println("Completed Move Syringe Module Down Fwd 83.9mm")
	time.Sleep(1 * time.Second)

	// Move Deck Rev 6.8mm to cut Hold of that tip
	distToTravel = at_X - cd["discard_small_hole"]
	response, err = m.setUpMotor(uint16(motors[0].Fast), uint16(math.Round(float64(motors[0].Steps)*distToTravel)), uint16(motors[0].Ramp), uint16(Fwd), uint16(motors[0].Number), On, Off)
	if err != nil {
		return
	}

	at_X = cd["discard_small_hole"]

	fmt.Println("Completed Move Deck Rev 6.8mm to cut Hold of that tip")
	time.Sleep(1 * time.Second)

	// Move Syringe Module Up Slow Rev 7.5mm
	response, err = m.setUpMotor(uint16(motors[1].Slow), uint16(math.Round(float64(motors[1].Steps)*cd["discard_tip_slow_up"])), uint16(motors[1].Ramp), uint16(Rev), uint16(motors[1].Number), On, Off)
	if err != nil {
		return
	}

	fmt.Println("Completed Move Syringe Module Up Slow Rev 7.5mm")
	time.Sleep(1 * time.Second)

	// Move Syringe Module Up Rev 137.5mm
	response, err = m.setUpMotor(uint16(motors[1].Fast), uint16(math.Round(float64(motors[1].Steps)*cd["syringe_end_max"])), uint16(motors[1].Ramp), uint16(Rev), uint16(motors[1].Number), On, Off)
	if err != nil {
		return
	}

	fmt.Println("Completed Move Syringe Module Up Rev 137.5mm")
	time.Sleep(1 * time.Second)

	fmt.Println("Home the machine now.")

	return "RUN SUCCESS", nil
}

/*

func (m MotorDriver) piercingHardcoded() (response string, err error) {



	// Well Distances - relative
	wellsDistanceInMM := []float64{98.24, 7.9, 7.9, 7.9, 9.9, 11.9, 11.9, 9.4}

	//*************************
	// Pierce Well after Well *
	//*************************
	for wellNum, distance := range wellsDistanceInMM {

		// Move Deck to reach the well
		response, err = m.setUpMotor(uint16(Fast), uint16(math.Round(DeckPulses*distance)), uint16(Ramp), uint16(Fwd), uint16(K5Deck), On, Off)
		if err != nil {
			return
		}

		fmt.Println("Completed Move Deck to reach the wellNum ", wellNum)
		time.Sleep(1 * time.Second)

		// Pierce the Well
		// Move Syringe Module Down Fwd 31mm
		response, err = m.setUpMotor(uint16(Fast), uint16(math.Round(NotDeckPulses*31)), uint16(Ramp), uint16(Fwd), uint16(K9SyringeModuleLHRH), On, Off)
		if err != nil {
			return
		}

		fmt.Println("Completed Move Syringe Module Down Fwd 31mm wellNum ", wellNum)
		time.Sleep(1 * time.Second)

		// Move Syringe Module Up Rev Fwd 31mm
		response, err = m.setUpMotor(uint16(Fast), uint16(math.Round(NotDeckPulses*31)), uint16(Ramp), uint16(Rev), uint16(K9SyringeModuleLHRH), On, Off)
		if err != nil {
			return
		}

		fmt.Println("Completed Move Syringe Module Up Rev 31mm wellNum ", wellNum)
		time.Sleep(1 * time.Second)

	}

	//**************
	// Tip Discard *
	//**************

	// Move Deck Fwd 150.54mm
	response, err = m.setUpMotor(uint16(Fast), uint16(math.Round(DeckPulses*150.54)), uint16(Ramp), uint16(Fwd), uint16(K5Deck), On, Off)
	if err != nil {
		return
	}

	fmt.Println("Completed Move Deck Fwd 150.54mm")
	time.Sleep(1 * time.Second)

	// Move Syringe Module Down Fwd 83.9mm
	response, err = m.setUpMotor(uint16(Fast), uint16(math.Round(NotDeckPulses*83.9)), uint16(Ramp), uint16(Fwd), uint16(K9SyringeModuleLHRH), On, Off)
	if err != nil {
		return
	}

	fmt.Println("Completed Move Syringe Module Down Fwd 83.9mm")
	time.Sleep(1 * time.Second)

	// Move Deck Rev 6.8mm to cut Hold of that tip
	response, err = m.setUpMotor(uint16(Fast), uint16(math.Round(DeckPulses*6.8)), uint16(Ramp), uint16(Fwd), uint16(K5Deck), On, Off)
	if err != nil {
		return
	}

	fmt.Println("Completed Move Deck Rev 6.8mm to cut Hold of that tip")
	time.Sleep(1 * time.Second)

	// Move Syringe Module Up Slow Rev 7.5mm
	response, err = m.setUpMotor(uint16(Slow), uint16(math.Round(NotDeckPulses*7.5)), uint16(Ramp), uint16(Rev), uint16(K9SyringeModuleLHRH), On, Off)
	if err != nil {
		return
	}

	fmt.Println("Completed Move Syringe Module Up Slow Rev 7.5mm")
	time.Sleep(1 * time.Second)

	// Move Syringe Module Up Rev 137.5mm
	response, err = m.setUpMotor(uint16(Fast), uint16(math.Round(NotDeckPulses*137.5)), uint16(Ramp), uint16(Rev), uint16(K9SyringeModuleLHRH), On, Off)
	if err != nil {
		return
	}

	fmt.Println("Completed Move Syringe Module Up Rev 137.5mm")
	time.Sleep(1 * time.Second)

	fmt.Println("Home the machine now.")

	return "RUN SUCCESS", nil
}

func (m MotorDriver) piercing() (response string, err error) {

	// **************
	// Tip PickUp	*
	//***************

	// check if run is already running, i.e check if motor is on and completion is off
	response, err = m.isRunInProgress()
	if err != nil {
		return
	}

	distToTravel := cd["piercing_tip"] - at_X
	// Move Deck Fwd 39.02mm
	response, err = m.setUpMotor(uint16(motors[0].Fast), uint16(math.Round(float64(motors[0].Steps)*distToTravel)), uint16(motors[0].Ramp), uint16(Fwd), uint16(motors[0].Number), On, Off)
	if err != nil {
		return
	}

	at_X = cd["piercing_tip"]

	fmt.Println("Completed Move Deck Fwd 39.02mm")
	time.Sleep(1 * time.Second)

	// Move Syringe Module Down Fwd 103mm
	response, err = m.setUpMotor(uint16(motors[1].Fast), uint16(math.Round(float64(motors[1].Steps)*cd["syringe_module_fast_down"])), uint16(motors[1].Ramp), uint16(Fwd), uint16(motors[1].Number), On, Off)
	if err != nil {
		return
	}

	fmt.Println("Completed Move Syringe Module Down Fwd 103mm")
	time.Sleep(1 * time.Second)

	// Move Syringe Module Down Slow Fwd 5mm
	response, err = m.setUpMotor(uint16(motors[1].Slow), uint16(math.Round(float64(motors[1].Steps)*cd["syringe_module_slow_down"])), uint16(motors[1].Ramp), uint16(Fwd), uint16(motors[1].Number), On, Off)
	if err != nil {
		return
	}

	fmt.Println("Completed Move Syringe Module Down Slow Fwd 5mm")
	time.Sleep(1 * time.Second)

	// Move Syringe Module Up Rev 59mm
	response, err = m.setUpMotor(uint16(motors[1].Fast), uint16(math.Round(float64(motors[1].Steps)*cd["pickup_piercing_tip_up"])), uint16(motors[1].Ramp), uint16(Rev), uint16(motors[1].Number), On, Off)
	if err != nil {
		return
	}

	fmt.Println("Completed Move Syringe Module Up Rev 59mm")
	time.Sleep(1 * time.Second)

	// Move Deck to Extraction Cartridge start
	distToTravel = cd["extraction_cartridge_start"] - at_X
	response, err = m.setUpMotor(uint16(motors[0].Fast), uint16(math.Round(float64(motors[0].Steps)*distToTravel)), uint16(motors[0].Ramp), uint16(Fwd), uint16(motors[0].Number), On, Off)
	if err != nil {
		return
	}

	at_X = cd["extraction_cartridge_start"]

	fmt.Println("Completed Move Deck to Extraction Cartridge start")
	time.Sleep(1 * time.Second)

	//	wellsDistanceInMM := []float64{17.5, 24.5, 33.3, 41.2, 51.2, 62.99, 74.9, 84.3}

	//*************************
	// Pierce Well after Well *
	//*************************
	//	for wellNum, distance := range wellsDistanceInMM {
	// Well Distances -> from "extraction_cartridge_start"
	for _, c := range cartriage {

		distToTravel = c.Distance - at_X
		// Move Deck to reach the well
		response, err = m.setUpMotor(uint16(motors[0].Fast), uint16(math.Round(float64(motors[0].Steps)*distToTravel)), uint16(motors[0].Ramp), uint16(Fwd), uint16(motors[0].Number), On, Off)
		if err != nil {
			return
		}

		at_X = cd["extraction_cartridge_start"] + c.Distance

		fmt.Println("Completed Move Deck to reach the wellNum ", c.WellNum)
		time.Sleep(1 * time.Second)

		// Pierce the Well
		// Move Syringe Module Down Fwd 31mm
		response, err = m.setUpMotor(uint16(motors[1].Fast), uint16(math.Round(float64(motors[1].Steps)*cd["piercing_height"])), uint16(motors[1].Ramp), uint16(Fwd), uint16(motors[1].Number), On, Off)
		if err != nil {
			return
		}

		fmt.Println("Completed Move Syringe Module Down Fwd 31mm wellNum ", c.WellNum)
		time.Sleep(1 * time.Second)

		// Move Syringe Module Up Rev Fwd 31mm
		response, err = m.setUpMotor(uint16(motors[1].Fast), uint16(math.Round(float64(motors[1].Steps)*cd["piercing_height"])), uint16(motors[1].Ramp), uint16(Rev), uint16(motors[1].Number), On, Off)
		if err != nil {
			return
		}

		fmt.Println("Completed Move Syringe Module Up Rev 31mm wellNum ", c.WellNum)
		time.Sleep(1 * time.Second)

	}

	//**************
	// Tip Discard *
	//**************

	// Move Deck Fwd 150.54mm

	distToTravel = cd["discard_big_hole"] - at_X
	response, err = m.setUpMotor(uint16(motors[0].Fast), uint16(math.Round(float64(motors[0].Steps)*distToTravel)), uint16(motors[0].Ramp), uint16(Fwd), uint16(motors[0].Number), On, Off)
	if err != nil {
		return
	}

	at_X = cd["discard_big_hole"]

	fmt.Println("Completed Move Deck Fwd 150.54mm")
	time.Sleep(1 * time.Second)

	// Move Syringe Module Down Fwd 83.9mm
	response, err = m.setUpMotor(uint16(motors[1].Fast), uint16(math.Round(float64(motors[1].Steps)*cd["piercing_tip_discard_height"])), uint16(motors[1].Ramp), uint16(Rev), uint16(motors[1].Number), On, Off)
	if err != nil {
		return
	}

	fmt.Println("Completed Move Syringe Module Down Fwd 83.9mm")
	time.Sleep(1 * time.Second)

	// Move Deck Rev 6.8mm to cut Hold of that tip
	distToTravel = at_X - cd["discard_small_hole"]
	response, err = m.setUpMotor(uint16(motors[0].Fast), uint16(math.Round(float64(motors[0].Steps)*distToTravel)), uint16(motors[0].Ramp), uint16(Fwd), uint16(motors[0].Number), On, Off)
	if err != nil {
		return
	}

	at_X = cd["discard_small_hole"]

	fmt.Println("Completed Move Deck Rev 6.8mm to cut Hold of that tip")
	time.Sleep(1 * time.Second)

	// Move Syringe Module Up Slow Rev 7.5mm
	response, err = m.setUpMotor(uint16(motors[1].Slow), uint16(math.Round(float64(motors[1].Steps)*cd["discard_tip_slow_up"])), uint16(motors[1].Ramp), uint16(Rev), uint16(motors[1].Number), On, Off)
	if err != nil {
		return
	}

	fmt.Println("Completed Move Syringe Module Up Slow Rev 7.5mm")
	time.Sleep(1 * time.Second)

	// Move Syringe Module Up Rev 137.5mm
	response, err = m.setUpMotor(uint16(motors[1].Fast), uint16(math.Round(float64(motors[1].Steps)*cd["syringe_end_max"])), uint16(motors[1].Ramp), uint16(Rev), uint16(motors[1].Number), On, Off)
	if err != nil {
		return
	}

	fmt.Println("Completed Move Syringe Module Up Rev 137.5mm")
	time.Sleep(1 * time.Second)

	fmt.Println("Home the machine now.")

	return "RUN SUCCESS", nil
}
*/
