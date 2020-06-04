package simulator

import "mylab/cpagent/plc"

var plcIO plcRegistors

type plcRegistors struct {
	d holdingRegistors // PLC holding resistors
	m coilRegistors    // PLC coil resistors
}

type holdingRegistors struct {
	heartbeat      uint16    // heartbeat register (W)
	currentTemp    uint16    // Current temperature (R)
	currentCycle   uint16    // Current cycle (R)
	idealLidTemp   uint16    // Ideal Lid temperature (W)
	currentLidTemp uint16    // Current Lid temperature (R)
	emissions      [6]uint16 // Well Emission data 96x6 registers (R)
	stage          plc.Stage // contains cycle data
	errCode        uint16    // error code (R)
}

type coilRegistors struct {
	homingSuccess  uint16 // homing success (R)
	homingErr      uint16 // homing error (R)
	startStopCycle uint16 // Start / Stop Cycle (W)
	restartCycle   uint16 // Restart Cycle (if rebooted during a run! (R)
	signalErr      uint16 // Signal Error (R)
	emmissionFlag  uint16 // Well Emmission register data  ON: PLC write & OFF: Read (RW)
	cycleCompleted uint16 // Cycle completed (R)
}
