package service

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func plcShakingHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var shakingID uuid.UUID

		vars := mux.Vars(req)
		shakingID, err := parseUUID(vars["id"])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		shaker, err := deps.Store.ShowShaking(req.Context(), shakingID)
		deck := vars["deck"]
		fmt.Printf("shaker object %v", shaker)
		switch deck {
		case "A", "B":
			_, err := deps.PlcDeck[deck].Shake(shaker)
			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		fmt.Fprintf(rw, fmt.Sprintf("%v", shaker))

	})

}
