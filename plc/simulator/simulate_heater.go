package simulator

import(
   "time"
)

// TODO: Handle Heater M 3
/*
   When Heater On
   - Monitor Heater
   - Set Present Value
   - Calibrate Temperature
   When Heater Off
   - Close Heater Monitoring
*/

type Celsius float64

const(
   // roomTemp already declared in pcr
   // rampUpTemp is temperature increase with heater ON
   rampUpRate Celsius = 0.2   // increase temp by 0.2 degree celsius per second when heater is ON
   // rampDownTemp is temperature decrease with heater OFF
   rampDownRate Celsius = 0.05 // decrease temp by 0.05 degree celsius per second when heater is OFF
)

func (d *SimulatorDriver) simulateOnHeater() (err error) {
   
   for {
		if !d.isHeaterInProgress() {
			break
		}
		time.Sleep(200 * time.Millisecond)
	}
	d.setHeaterInProgress()
	defer d.resetHeaterInProgress()

   

	return
}
