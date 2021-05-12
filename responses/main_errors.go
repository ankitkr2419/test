package responses

import (
	"fmt"
)

var (
	UnsupportedPLCError = fmt.Errorf("error unsupported PLC. Valid PLC: 'simulator' or 'compact32'")
	DatabaseInitError   = fmt.Errorf("error database init failed")
	DBAllSetupError     = fmt.Errorf("error loading DB Setups failed")
	PLCAllLoadError     = fmt.Errorf("error loading PLC Functions failed")
	ServiceAllLoadError = fmt.Errorf("error loading Service Functions failed")
)
