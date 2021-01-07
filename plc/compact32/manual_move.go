package compact32

import (
	"encoding/binary"
)

func (d *Compact32Deck) ManualMovement(motorNum, direction, pulses uint16) (response string, err error) {
	sensorHasCut = false
	aborted = false
	var sensorAddressBytes = []byte{0x08, 0x02}
	sensorAddressUint16 := binary.BigEndian.Uint16(sensorAddressBytes)

	return d.SetupMotor(uint16(2000), pulses, uint16(100), direction, motorNum, uint16(0xff00), uint16(sensorAddressUint16))
}
