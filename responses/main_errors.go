package responses

import (
	"fmt"
)

var (
	UnsupportedPLCError = fmt.Errorf("error unsupported PLC. Valid PLC: 'simulator' or 'compact32'")
	UnsupportedTECError = fmt.Errorf("error unsupported TEC. Valid PLC: 'simulator' or 'compact32'")
	UnknownCase         = fmt.Errorf("error unknown case reached")
	DatabaseInitError   = fmt.Errorf("error database init failed")
	DBAllSetupError     = fmt.Errorf("error loading DB Setups failed")
	PLCAllLoadError     = fmt.Errorf("error loading PLC Functions failed")
	ServiceAllLoadError = fmt.Errorf("error loading Service Functions failed")
	WriteToFileError    = fmt.Errorf("failed to write output to log file")
)
