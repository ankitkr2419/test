package service

import (
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"mylab/cpagent/tec"
	"net/http"
)

func runTECHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		file := db.GetExcelFile(tec.LogsPath, "test_output")
		p := plc.Stage{
			Holding: []plc.Step{
				plc.Step{TargetTemp: 25, RampUpTemp: 2, HoldTime: 1, DataCapture: false},
				plc.Step{TargetTemp: 50, RampUpTemp: 2, HoldTime: 1, DataCapture: false},
				plc.Step{TargetTemp: 95, RampUpTemp: 2, HoldTime: 1, DataCapture: false},
			},
			Cycle: []plc.Step{
				plc.Step{TargetTemp: 95, RampUpTemp: 2, HoldTime: 1, DataCapture: false},
				plc.Step{TargetTemp: 60, RampUpTemp: 2, HoldTime: 1, DataCapture: false},
			},
			CycleCount: 3,
		}

		go startExp(deps, p, file)

		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: "Run Started success"})
	})
}
