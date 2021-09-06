package simulator

import (
	"mylab/cpagent/plc"
)

/* REGISTERS_EXTRACTION Mappings:
 * Here address represents the register address
 * Here value represents the uint16 value we have received
 *	plc.DeckA => {
		"D" => {
	 		MODBUS_EXTRACTION[plc.DeckA]["D"][address] => uint16(value)
  		},
   		"M" => {
	 		MODBUS_EXTRACTION[plc.DeckA]["M"][address] => uint16(value)
  		}
	}
	plc.DeckB => {
		"D" => {
	 		MODBUS_EXTRACTION[plc.DeckB]["D"][address] => uint16(value)
  		},
   		"M" => {
	 		MODBUS_EXTRACTION[plc.DeckB]["M"][address] => uint16(value)
  		}
	}


   So, REGISTERS_EXTRACTION[plc.DeckA]["D"][ MODBUS_EXTRACTION[plc.DeckA]["D"][200] ] will give us the register value at D 200 for deck A
   And REGISTERS_EXTRACTION[plc.DeckB]["M"][ MODBUS_EXTRACTION[plc.DeckB]["M"][0] ] will give us register value at M 0 for deck B
*/

// TODO: Use sync.Map
var REGISTERS_EXTRACTION map[string]map[string]map[uint16]uint16 = map[string]map[string]map[uint16]uint16{
	// Registers of Deck A
	plc.DeckA: map[string]map[uint16]uint16{
		// Data Registers
		"D": map[uint16]uint16{
			plc.MODBUS_EXTRACTION[plc.DeckA]["D"][200]: uint16(0),             // Motor Speed (W)
			plc.MODBUS_EXTRACTION[plc.DeckA]["D"][202]: uint16(0),             // Motor Pulses (W)
			plc.MODBUS_EXTRACTION[plc.DeckA]["D"][204]: uint16(0),             // Motor Ramp Rate (W)
			plc.MODBUS_EXTRACTION[plc.DeckA]["D"][206]: uint16(0),             // Motor Direction. 1=Towards Sensor, 0=Against Sensor (W)
			plc.MODBUS_EXTRACTION[plc.DeckA]["D"][208]: uint16(0),             // Shaker Temperature Set Value (W)
			plc.MODBUS_EXTRACTION[plc.DeckA]["D"][210]: uint16(roomTemp * 10), // Shaker Temperature present value LH (R)
			plc.MODBUS_EXTRACTION[plc.DeckA]["D"][212]: uint16(0),             // Pulses Executed (R)
			plc.MODBUS_EXTRACTION[plc.DeckA]["D"][214]: uint16(0),             // Heartbeat. PC=2, PLC=1, time=200ms (W/R)
			plc.MODBUS_EXTRACTION[plc.DeckA]["D"][216]: uint16(0),             // Shaker RPM present value (R)
			plc.MODBUS_EXTRACTION[plc.DeckA]["D"][218]: uint16(0),             // Shaker RPM set value (Note 1) (W)
			plc.MODBUS_EXTRACTION[plc.DeckA]["D"][220]: uint16(0),             // Shaker selection (Note 2) (W)
			plc.MODBUS_EXTRACTION[plc.DeckA]["D"][222]: uint16(0),             // Shaker heater selection (Note 3) (W)
			plc.MODBUS_EXTRACTION[plc.DeckA]["D"][224]: uint16(roomTemp * 10), // Shaker temperature present value RH (R)
			plc.MODBUS_EXTRACTION[plc.DeckA]["D"][226]: uint16(0),             // Motor Number (W)
			plc.MODBUS_EXTRACTION[plc.DeckA]["D"][504]: uint16(0),             // PID Tuning Completed LH (R) (4 = DONE)
			plc.MODBUS_EXTRACTION[plc.DeckA]["D"][534]: uint16(0),             // PID Tuning Completed RH (R) (4 = DONE)
		},
		// Coil registers: ON:0xFF00, OFF: 0x0000
		"M": map[uint16]uint16{
			plc.MODBUS_EXTRACTION[plc.DeckA]["M"][0]: uint16(0), // Motor ON/OFF . (W)
			plc.MODBUS_EXTRACTION[plc.DeckA]["M"][1]: uint16(0), // Pulses Completion bit. ON = uint16(1), OFF = uint16(0) (R)
			plc.MODBUS_EXTRACTION[plc.DeckA]["M"][2]: uint16(0), // Sensor Has Cut. ON = uint16(3), OFF = uint16(2)(R)
			plc.MODBUS_EXTRACTION[plc.DeckA]["M"][3]: uint16(0), // Heater ON/OFF (W)
			plc.MODBUS_EXTRACTION[plc.DeckA]["M"][4]: uint16(0), // PID calibration LH ON/OFF (W)
			plc.MODBUS_EXTRACTION[plc.DeckA]["M"][5]: uint16(0), // Shaker ON/OFF (W)
			plc.MODBUS_EXTRACTION[plc.DeckA]["M"][6]: uint16(0), // UV Light ON/OFF (W)
			plc.MODBUS_EXTRACTION[plc.DeckA]["M"][7]: uint16(0), // Light ON/OFF (W)
			plc.MODBUS_EXTRACTION[plc.DeckA]["M"][8]: uint16(0), // PID calibration RH ON/OFF (W)
		},
	},
	// Deck B
	plc.DeckB: map[string]map[uint16]uint16{
		// Data Registers
		"D": map[uint16]uint16{
			plc.MODBUS_EXTRACTION[plc.DeckB]["D"][200]: uint16(0),             // Motor Speed (W)
			plc.MODBUS_EXTRACTION[plc.DeckB]["D"][202]: uint16(0),             // Motor Pulses (W)
			plc.MODBUS_EXTRACTION[plc.DeckB]["D"][204]: uint16(0),             // Motor Ramp Rate (W)
			plc.MODBUS_EXTRACTION[plc.DeckB]["D"][206]: uint16(0),             // Motor Direction. 1=Towards Sensor, 0=Against Sensor (W)
			plc.MODBUS_EXTRACTION[plc.DeckB]["D"][208]: uint16(0),             // Shaker Temperature Set Value (W)
			plc.MODBUS_EXTRACTION[plc.DeckB]["D"][210]: uint16(roomTemp * 10), // Shaker Temperature present value LH (R)
			plc.MODBUS_EXTRACTION[plc.DeckB]["D"][212]: uint16(0),             // Pulses Executed (R)
			plc.MODBUS_EXTRACTION[plc.DeckB]["D"][214]: uint16(0),             // Heartbeat. PC=2, PLC=1, time=200ms (W/R)
			plc.MODBUS_EXTRACTION[plc.DeckB]["D"][216]: uint16(0),             // Shaker RPM present value (R)
			plc.MODBUS_EXTRACTION[plc.DeckB]["D"][218]: uint16(0),             // Shaker RPM set value (Note 1) (W)
			plc.MODBUS_EXTRACTION[plc.DeckB]["D"][220]: uint16(0),             // Shaker selection (Note 2) (W)
			plc.MODBUS_EXTRACTION[plc.DeckB]["D"][222]: uint16(0),             // Shaker heater selection (Note 3) (W)
			plc.MODBUS_EXTRACTION[plc.DeckB]["D"][224]: uint16(roomTemp * 10), // Shaker temperature present value RH (R)
			plc.MODBUS_EXTRACTION[plc.DeckB]["D"][226]: uint16(0),             // Motor Number (W)
			plc.MODBUS_EXTRACTION[plc.DeckA]["D"][504]: uint16(0),             // PID Tuning Completed LH (R) (4 = DONE)
			plc.MODBUS_EXTRACTION[plc.DeckA]["D"][534]: uint16(0),             // PID Tuning Completed RH (R) (4 = DONE)
		},
		// Coil registers: ON:0xFF00, OFF: 0x0000
		"M": map[uint16]uint16{
			plc.MODBUS_EXTRACTION[plc.DeckB]["M"][0]: uint16(0), // Motor ON/OFF . (W)
			plc.MODBUS_EXTRACTION[plc.DeckB]["M"][1]: uint16(0), // Pulses Completion bit. ON = uint16(1), OFF = uint16(0) (R)
			plc.MODBUS_EXTRACTION[plc.DeckB]["M"][2]: uint16(0), // Sensor Has Cut. ON = uint16(3), OFF = uint16(2)(R)
			plc.MODBUS_EXTRACTION[plc.DeckB]["M"][3]: uint16(0), // Heater ON/OFF (W)
			plc.MODBUS_EXTRACTION[plc.DeckB]["M"][4]: uint16(0), // PID calibration LH ON/OFF (W)
			plc.MODBUS_EXTRACTION[plc.DeckB]["M"][5]: uint16(0), // Shaker ON/OFF (W)
			plc.MODBUS_EXTRACTION[plc.DeckB]["M"][6]: uint16(0), // UV Light ON/OFF (W)
			plc.MODBUS_EXTRACTION[plc.DeckB]["M"][7]: uint16(0), // Light ON/OFF (W)
			plc.MODBUS_EXTRACTION[plc.DeckB]["M"][8]: uint16(0), // PID calibration RH ON/OFF (W)
		},
	},
}
