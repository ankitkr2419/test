package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Manual struct {
	MotorNum  int `json:"motor_number"`
	Pulses    int `json:"pulses"`
	Direction int `json:"direction"`
}

func manualHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		var m Manual
		fmt.Println(req.Body)
		err := json.NewDecoder(req.Body).Decode(&m)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			fmt.Println("Error decoding manual data")
			return
		}

		switch {
		case m.MotorNum <= 4 || m.MotorNum > 10:
			rw.WriteHeader(http.StatusBadRequest)
			err = fmt.Errorf("Select motor num in only in between 5-10")
			fmt.Println(err)
			return
		case m.Direction != 0 && m.Direction != 1:
			rw.WriteHeader(http.StatusBadRequest)
			err = fmt.Errorf("Select motor direction in only as 0 or 1")
			fmt.Println(err)
			return
		case m.Pulses > 10000:
			rw.WriteHeader(http.StatusBadRequest)
			err = fmt.Errorf("Consider pulses only less than or equal to 10000")
			fmt.Println(err)
			return
		}

		response, err := deps.PlcDeckA.ManualMovement(uint16(m.MotorNum), uint16(m.Direction), uint16(m.Pulses))
		if err != nil {
			fmt.Fprintf(rw, err.Error())
		} else {
			fmt.Fprintf(rw, response)
		}
	})
}

func pauseHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		response, err := deps.PlcDeckA.Pause()
		if err != nil {
			fmt.Fprintf(rw, err.Error())
		} else {
			fmt.Fprintf(rw, response)
		}
	})
}

func resumeHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		response, err := deps.PlcDeckA.Resume()
		if err != nil {
			fmt.Fprintf(rw, err.Error())
		} else {
			fmt.Fprintf(rw, response)
		}
	})
}

func abortHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		response, err := deps.PlcDeckA.Abort()
		if err != nil {
			fmt.Fprintf(rw, err.Error())
		} else {
			fmt.Fprintf(rw, response)
		}
	})
}
