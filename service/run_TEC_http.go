package service

import (
	"mylab/cpagent/plc"
	"net/http"
)

func runTECHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		p := plc.Stage{
			Holding: []plc.Step{
				plc.Step{25, 2, 1, false},
				plc.Step{50, 2, 1, false},
				plc.Step{95, 2, 1, false},
			},
			Cycle: []plc.Step{
				plc.Step{95, 2, 1, false},
				plc.Step{60, 2, 1, false},
			},
			CycleCount: 3,
		}

		go startExp(deps, p)

		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: "Run Started success"})
	})
}
