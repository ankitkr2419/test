package service

import (
	"fmt"
	"net/http"
)

func deckHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		response, err := deps.PlcDeckA.DeckHoming()
		if err != nil {
			fmt.Fprintf(rw, err.Error())
		} else {
			fmt.Fprintf(rw, response)
		}
	})
}

/*
func (deps Dependencies) deckHoming() (response string, err error) {

	// M2
	var sensorAddressBytes = []byte{0x08, 0x02}
	sensorAddressUint16 := binary.BigEndian.Uint16(sensorAddressBytes)

	sensorHasCut = false
	fmt.Println("Deck is moving forward")
	response, err = m.SetupMotor(uint16(2000), uint16(59199), uint16(100), uint16(1), uint16(5), uint16(0xff00), sensorAddressUint16)
	if err != nil {
		return
	}

	sensorHasCut = false
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Deck is moving back by and after not cut -> 2000")
	response, err = m.SetupMotor(uint16(2000), uint16(19999), uint16(100), uint16(0), uint16(5), uint16(0xff00), sensorAddressUint16)
	if err != nil {
		return
	}

	time.Sleep(100 * time.Millisecond)

	fmt.Println("Deck is moving forward again by 2999")
	response, err = m.SetupMotor(uint16(500), uint16(2999), uint16(100), uint16(1), uint16(5), uint16(0xff00), sensorAddressUint16)

	fmt.Println("Deck homing is completed.")

	return "DECK HOMING SUCCESS", nil
}
*/
