package simulator

type plcRegistors struct {
	d holdingRegistors // PLC holding resistors
	m coilRegistors    // PLC coil resistors
}

type holdingRegistors struct {
	heartbeat      uint16 // heartbeat register (W)
	currentTemp    uint16 // Current temperature (R)
	currentCycle   uint16 // Current cycle (R)
	idealLidTemp   uint16 // Ideal Lid temperature (W)
	currentLidTemp uint16 // Current Lid temperature (R)
	errCode        uint16 // error code (R)

}

type coilRegistors struct {
	homingSuccess  uint16 // homing success (R)
	homingErr      uint16 // homing error (R)
	startStopCycle uint16 // Start / Stop Cycle (W)
	restartCycle   uint16 // Restart Cycle (if rebooted during a run! (R)
	signalErr      uint16 // Signal Error (R)
	emissionFlag   uint16 // Well Emmission register data  ON: PLC write & OFF: Read (RW)
	cycleCompleted uint16 // Cycle completed (R)

	//updated
	cycleStart  uint16 // cycle start
	homing      uint16 // homing button
	resetValues uint16 // reset values
	setRotation uint16 // set rotation button
}
