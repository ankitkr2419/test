package compact32

import (
	"time"
)

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
		100:  uint16(0x0064), // heartbeat register (W)
		101:  uint16(0x0065), // offset for holding stage config (W)
		113:  uint16(0x0071), // offset for cycling stage config (W)
		131:  uint16(0x0083), // Number of cycles to run (W)
		132:  uint16(0x0084), // Current temperature (R)
		133:  uint16(0x0085), // Current cycle (R)
		134:  uint16(0x0085), // Ideal Lid temperature (W)
		135:  uint16(0x0086), // Current Lid temperature (R)
		410:  uint16(0x019A), // General regisrter offset for values! (unused for now)
		2000: uint16(0x07D0), // Well Emission data 96x6 registers (R)
		2577: uint16(0x0A11), // error code (R)
	},
	// Coil registers: ON:0xFF00, OFF: 0x0000
	"M": map[int]uint16{
		100: uint16(0x0064), // homing success (R)
		101: uint16(0x0065), // homing error (R)
		102: uint16(0x0066), // Start / Stop Cycle (W)
		103: uint16(0x0067), // UnUsed
		104: uint16(0x0068), // Restart Cycle (if rebooted during a run! (R)
		105: uint16(0x0069), // Signal Error (R)
		106: uint16(0x006A), // Well Emmission register data  ON: PLC write & OFF: Read (RW)
		107: uint16(0x006B), // Cycle completed (R)
	},
}

var LOOKUP map[string]string = map[string]string{
	"heartbeat": "D101",
}
