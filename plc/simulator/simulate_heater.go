package simulator

import(
   "fmt"
   "time"
	logger "github.com/sirupsen/logrus"
	"mylab/cpagent/plc"

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
   // roomTemp already declared in pcr.go
   // rampUpTemp is temperature increase with heater ON
   rampUpRate Celsius = 0.2   // increase temp by 0.2 degree celsius per second when heater is ON
   // rampDownTemp is temperature decrease with heater OFF
   rampDownRate Celsius = 0.05 // decrease temp by 0.05 degree celsius per second when heater is OFF
)

// Currently we only handle if both shaker's start together and not separate

// WARN: Be careful with temperature, its multiplied by 10 for machine
func (d *SimulatorDriver) simulateOnHeater() (err error) {
   
   // ASK: Do we need to Handle for separate shaker part ?
   heaterNum :=  d.readRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][222] )
   if heaterNum != uint16(3) {
      err = fmt.Errorf("simulator support is only provided for LH, RH shaker heaters together only")
      return
   }
   for {
		if !d.isHeaterInProgress() {
			break
		}
		time.Sleep(200 * time.Millisecond)
	}
	d.setHeaterInProgress()
	defer d.resetHeaterInProgress()

   logger.Infoln("Heater has started!!")

   // Update Temperature every sec
   d.updateTemperature()

	return
}

func (d *SimulatorDriver) updateTemperature() {

   for{
      time.Sleep(time.Second)
      if plc.OFF == d.readRegister("M", plc.MODBUS_EXTRACTION[d.DeckName]["M"][3]) {
         go d.coolDown()
         return
      }
   
      targetTemp := d.readRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][208] )
      currentTempLH :=  d.readRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][210] )
      currentTempRH :=  d.readRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][224] )
   
      // 1 degree extra
      logger.Infoln("Heating Up -> targetTemp :", targetTemp , ", currentTempLH :", currentTempLH, ", currentTempRH", currentTempRH)

      if currentTempLH  < targetTemp  + 10 {
         d.setRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][210] , currentTempLH  + uint16(rampUpRate * 10))
      } 
      if currentTempRH  < targetTemp  + 10{
         d.setRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][224] , currentTempRH  + uint16(rampUpRate * 10))
      }
   }
}

// cool Down to room Temperature if heater isn't in progress, update every 2 sec
func (d *SimulatorDriver) coolDown() {
   for{
      time.Sleep(2 * time.Second)
      if d.isHeaterInProgress() {
         return
      }

      currentTempLH :=  d.readRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][210] )
      currentTempRH :=  d.readRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][224] )
   
      logger.Infoln("Cooling Down -> currentTempLH :", currentTempLH, ", currentTempRH", currentTempRH)
      // 1 degree extra
      if currentTempLH > uint16(roomTemp * 10 + 10) {
         d.setRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][210] , currentTempLH - uint16(rampDownRate * 20))
      } 
      if currentTempRH > uint16(roomTemp * 10 + 10) {
         d.setRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][224] , currentTempRH + uint16(rampDownRate * 20))
      }
   }
}