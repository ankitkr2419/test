package service

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func discardAndHomeHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		var discard bool
		var err error
		vars := mux.Vars(req)
		deck := vars["deck"]

		discardTip := vars["discard"]

		if discard, err = strconv.ParseBool(discardTip); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(`Invalid boolean value for tip discard option`))
			err := fmt.Errorf("Discard option Should be boolean only")
			deps.WsErrCh <- err
		}

		switch deck {
		case "A", "B":
			rw.WriteHeader(http.StatusOK)
			rw.Write([]byte(`discard-tip-and-home in progress`))
			go deps.PlcDeck[deck].DiscardTipAndHome(discard)
		default:
			err := fmt.Errorf("Check your deck name")
			deps.WsErrCh <- err
		}

	})
}
