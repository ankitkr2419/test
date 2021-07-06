package service

import (
	"mylab/cpagent/plc"
	"mylab/cpagent/tec"
	"net/http"
)

func runTECHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		file := plc.GetExcelFile(tec.LogsPath, "test_output")
		p := plc.Stage{
			Holding: []plc.Step{
				plc.Step{25, 2, 0},
				plc.Step{50, 2, 0},
				plc.Step{95, 2, 0},
			},
			Cycle: []plc.Step{
				plc.Step{95, 2, 0},
				plc.Step{60, 2, 0},
			},
			CycleCount: 3,
		}

		go startExp(deps, p, file)

		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: "Run Started success"})
	})
}
