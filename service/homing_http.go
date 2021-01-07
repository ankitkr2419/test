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

func syringeHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		response, err := deps.PlcDeckA.SyringeHoming()
		if err != nil {
			fmt.Fprintf(rw, err.Error())
		} else {
			fmt.Fprintf(rw, response)
		}
	})
}

func syringeModuleHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		response, err := deps.PlcDeckA.SyringeModuleHoming()
		if err != nil {
			fmt.Fprintf(rw, err.Error())
		} else {
			fmt.Fprintf(rw, response)
		}
	})
}

func magnetHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		response, err := deps.PlcDeckA.MagnetHoming()
		if err != nil {
			fmt.Fprintf(rw, err.Error())
		} else {
			fmt.Fprintf(rw, response)
		}
	})
}

func homingHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		response, err := deps.PlcDeckA.Homing()
		if err != nil {
			fmt.Fprintf(rw, err.Error())
		} else {
			fmt.Fprintf(rw, response)
		}
	})
}
