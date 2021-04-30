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
			go deps.PlcDeck[deck].DiscardTipAndHome(discard)
			rw.Header().Add("Content-Type", "application/json")
			rw.WriteHeader(http.StatusOK)
			rw.Write([]byte(fmt.Sprintf(`"msg":"discard-tip-and-home in progress","deck": "%v"`, deck)))
		default:
			err := fmt.Errorf("Check your deck name")
			deps.WsErrCh <- err
		}

		if err != nil {
			fmt.Fprintf(rw, err.Error())
			fmt.Println(err.Error())
			rw.WriteHeader(http.StatusBadRequest)
		}

	})
}