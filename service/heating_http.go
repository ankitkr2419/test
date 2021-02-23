package service

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func plcHeatingHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var heatingID uuid.UUID
		//here we are hardcoding the shaker no in future this is to be fetched dynamically.
		shakerNo := 3
		vars := mux.Vars(req)
		heatingID, err := parseUUID(vars["id"])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		heat, err := deps.Store.GetHeating(req.Context(), heatingID)
		deck := vars["deck"]

		fmt.Printf("heat object %v", heat)
		switch deck {
		case "A", "B":
			_, err := deps.PlcDeck[deck].Heat(uint16(heat.Temperature), uint16(shakerNo), heat.FollowTemp, heat.Duration)
			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("done"))
	})
}
