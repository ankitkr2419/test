package service

import (
	"net/http"
	"mylab/cpagent/plc"
	
)

func runTECHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		var err error

		p := plc.Stage{
		Holding: []plc.Step{
			plc.Step{65.3, 10, 5},
			plc.Step{85.3, 10, 5},
			plc.Step{95, 10, 5},
		},
		Cycle: []plc.Step{
			// plc.Step{60, 10, 10},
			plc.Step{95, 10, 5},
			plc.Step{85, 10, 5},
			plc.Step{75, 10, 5},
			plc.Step{65, 10, 5},
		},
		CycleCount: 3,
	}
		err = startExp(deps, p)
		if err != nil{
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: err.Error()})
			return
		}

		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: "Run Started success"} )
	})
}