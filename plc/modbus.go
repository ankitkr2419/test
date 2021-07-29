package plc

import (
	"time"
)

const (
	ON  = uint16(0xFF00)
	OFF = uint16(0x0000)
)

const (
	UP            = uint16(1)
	DOWN          = uint16(0)
	FWD           = uint16(1)
	REV           = uint16(0)
	ASPIRE        = uint16(0)
	DISPENSE      = uint16(1)
	TowardsSensor = uint16(1)
	AgainstSensor = uint16(0)
	SensorUncut   = 2
	SensorCut     = 3
)

// *** NOTE ***
// For Syringe ASPIRE means syringe goes UP and DISPENSE means syringe goes DOWN
// This is because of hardware compatibility
// 1 means towards sensor, 0 means against sensor
// ************

/* Modbus Mappings:
 *
 * "D" => {
	 101 => uint16(0x0065)
	 2000 => uint16(0x07D0)
   },
   "M" => {
	 132 => uint16(0x1084)
   }

   So, MODBUS["D"][101] will give us the exact modbus address of 0x0065.
*/

const (
	heartbeat time.Duration = 200 * time.Millisecond
)

var MODBUS map[string]map[int]uint16 = map[string]map[int]uint16{
	// Data Registers
	"D": map[int]uint16{
		//updated addresses
		23: uint16(0x1017), //register for setting rotation pulses
		44: uint16(0x102C), //start address of register For FAM
		60: uint16(0x103C), //start address of register For VIC

		100: uint16(0x1064), // heartbeat register (W)
		// 101:  uint16(0x1065), // offset for holding stage config (W)
		// 113:  uint16(0x1071), // offset for cycling stage config (W)
		// 131:  uint16(0x1083), // Number of cycles to run (W)
		// 132:  uint16(0x1084), // Current temperature (R)
		// 133:  uint16(0x1085), // Current cycle (R)
		134: uint16(0x1086), // Ideal Lid temperature (W)
		135: uint16(0x1087), // Current Lid temperature (R)
		410: uint16(0x119A), // General register offset for values! (unused for now)
		// 2000: uint16(0x17D0), // Well Emission data 96x6 registers (R)
		2577: uint16(0x1A11), // error code (R)
	},
	// Coil registers: ON:0xFF00, OFF: 0x0000
	"M": map[int]uint16{

		//Updated addresses
		1: uint16(0x0801), //combined Homing rt-pcr
		2: uint16(0x0802), //combined Homing rt-pcr
		// 14: uint16(0x080E), //combined with M15 ON/OFF rotate button
		// 15: uint16(0x080F), //combined with M14 ON/OFF rotate button
		20: uint16(0x0814), //combined with M21 ON/OFF cycle button
		21: uint16(0x0815), //combined with M20 ON/OFF cycle button
		25: uint16(0x0819), //reset values

		100: uint16(0x0864), // homing success (R)
		101: uint16(0x0865), // homing error (R)
		// 102: uint16(0x0866), // Start / Stop Cycle (W)
		103: uint16(0x0867), // UnUsed
		104: uint16(0x0868), // Restart Cycle (if rebooted during a run! (R)
		105: uint16(0x0869), // Signal Error (R)
		106: uint16(0x086A), // Well Emmission register data  ON: PLC write & OFF: Read (RW)
		107: uint16(0x086B), // Cycle completed (R)
		109: uint16(0x086D), // Lid Heating On (W)
	},
}

var LOOKUP map[string]string = map[string]string{
	"heartbeat": "D100",
}

/* MODBUS_EXTRACTION Mappings:
 *
 *	DeckA => {
		"D" => {
	 		101 => uint16(0x0065)
	 		2000 => uint16(0x07D0)
  		},
   		"M" => {
	 		8: uint16(0x0808)
  		 }
	}
	DeckB => {
		"D" => {
	 		200: uint16(0x10C8)
			226: uint16(0x10E2)
  		},
   		"M" => {
	 		0: uint16(0x0800)
  		 }
	}


   So, MODBUS_EXTRACTION[DeckA]["D"][101] will give us the exact modbus address of 0x0065 for DECK A.
   And MODBUS_EXTRACTION[DeckB]["M"][0] will give us the exact modbus address of 0x0800 for DECK B.
*/

var MODBUS_EXTRACTION map[string]map[string]map[int]uint16 = map[string]map[string]map[int]uint16{
	// Deck A
	DeckA: map[string]map[int]uint16{
		// Data Registers
		"D": map[int]uint16{
			200: uint16(0x10C8), // Motor Speed (W)
			202: uint16(0x10CA), // Motor Pulses (W)
			204: uint16(0x10CC), // Motor Ramp Rate (W)
			206: uint16(0x10CE), // Motor Direction. 1=Towards Sensor, 0=Against Sensor (W)
			208: uint16(0x10D0), // Shaker Temperature Set Value (W)
			210: uint16(0x10D2), // Shaker Temperature present value LH (R)
			212: uint16(0x10D4), // Pulses Executed (R)
			214: uint16(0x10D6), // Heartbeat. PC=2, PLC=1, time=200ms (W/R)
			216: uint16(0x10D8), // Shaker Pulses present value (R)
			218: uint16(0x10DA), // Shaker Pulses set value (Note 1) (W)
			220: uint16(0x10DC), // Shaker selection (Note 2) (W)
			222: uint16(0x10DE), // Shaker heater selection (Note 3) (W)
			224: uint16(0x10E0), // Shaker temperature present value RH (R)
			226: uint16(0x10E2), // Motor Number (W)
		},
		// Coil registers: ON:0xFF00, OFF: 0x0000
		"M": map[int]uint16{
			0: uint16(0x0800), // Motor ON/OFF . (W)
			1: uint16(0x0801), // Pulses Completion bit. ON = uint16(1), OFF = uint16(0) (R)
			2: uint16(0x0802), // Sensor Has Cut. ON = uint16(3), OFF = uint16(2)(R)
			3: uint16(0x0803), // Heater ON/OFF (W)
			4: uint16(0x0804), // PID calibration LH ON/OFF (W)
			5: uint16(0x0805), // Shaker ON/OFF (W)
			6: uint16(0x0806), // UV Light ON/OFF (W)
			7: uint16(0x0807), // Light ON/OFF (W)
			8: uint16(0x0808), // PID calibration RH ON/OFF (W)
		},
	},
	// Deck B
	DeckB: map[string]map[int]uint16{
		// Data Registers
		"D": map[int]uint16{
			200: uint16(0x10C8), // Motor Speed (W)
			202: uint16(0x10CA), // Motor Pulses (W)
			204: uint16(0x10CC), // Motor Ramp Rate (W)
			206: uint16(0x10CE), // Motor Direction. 1=Towards Sensor, 0=Against Sensor (W)
			208: uint16(0x10D0), // Shaker Temperature Set Value (W)
			210: uint16(0x10D2), // Shaker Temperature present value LH (R)
			212: uint16(0x10D4), // Pulses Executed (R)
			214: uint16(0x10D6), // Heartbeat. PC=2, PLC=1, time=200ms (W/R)
			216: uint16(0x10D8), // Shaker Pulses present value (R)
			218: uint16(0x10DA), // Shaker Pulses set value (Note 1) (W)
			220: uint16(0x10DC), // Shaker selection (Note 2) (W)
			222: uint16(0x10DE), // Shaker heater selection (Note 3) (W)
			224: uint16(0x10E0), // Shaker temperature present value RH (R)
			226: uint16(0x10E2), // Motor Number (W)
		},
		// Coil registers: ON:0xFF00, OFF: 0x0000
		"M": map[int]uint16{
			0: uint16(0x0800), // Motor ON/OFF . (W)
			1: uint16(0x0801), // Pulses Completion bit. ON = uint16(1), OFF = uint16(0) (R)
			2: uint16(0x0802), // Sensor Has Cut. ON = uint16(3), OFF = uint16(2)(R)
			3: uint16(0x0803), // Heater ON/OFF (W)
			4: uint16(0x0804), // PID calibration LH ON/OFF (W)
			5: uint16(0x0805), // Shaker ON/OFF (W)
			6: uint16(0x0806), // UV Light ON/OFF (W)
			7: uint16(0x0807), // Light ON/OFF (W)
			8: uint16(0x0808), // PID calibration RH ON/OFF (W)
		},
	},
	/*
	   ***Note 1: Shaker RPM***
	   1 = 500 RPM  = 6500 Pulses
	   2 = 800 RPM  = 10500 Pulses
	   3 = 1100 RPM = 14500 Pulses

	   ***Note 2: Shaker selection***
	   1 = LH shaker ON
	   2 = RH Shaker ON
	   3 = LH + RH shaker ON

	   ***Note 3: Shaker heater selection***
	   1 = LH shaker heater ON
	   2 = RH shaker heater ON
	   3 = LH + RH shaker heater ON
	*/
}
