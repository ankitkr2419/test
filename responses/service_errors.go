package responses

import (
	"fmt"
)

var (
	PiercingDecodeError     = fmt.Errorf("error decoding piercing record")
	PiercingValidationError = fmt.Errorf("error validating piercing record")
	PiercingCreateError     = fmt.Errorf("error creating piercing record")
	PiercingFetchError      = fmt.Errorf("error fetching piercing record")
	PiercingUpdateError     = fmt.Errorf("error updating piercing record")

	AspireDispenseDecodeError     = fmt.Errorf("error decoding aspire dispense record")
	AspireDispenseValidationError = fmt.Errorf("error validating aspire dispense record")
	AspireDispenseCreateError     = fmt.Errorf("error creating aspire dispense record")
	AspireDispenseFetchError      = fmt.Errorf("error fetching aspire dispense record")
	AspireDispenseUpdateError     = fmt.Errorf("error updating aspire dispense record")

	DelayDecodeError     = fmt.Errorf("error decoding delay record")
	DelayValidationError = fmt.Errorf("error validating delay record")
	DelayCreateError     = fmt.Errorf("error creating delay record")
	DelayFetchError      = fmt.Errorf("error fetching delay record")
	DelayUpdateError     = fmt.Errorf("error updating delay record")

	AttachDetachDecodeError     = fmt.Errorf("error decoding attach detach record")
	AttachDetachValidationError = fmt.Errorf("error validating attach detach record")
	AttachDetachCreateError     = fmt.Errorf("error creating attach detach record")
	AttachDetachFetchError      = fmt.Errorf("error fetching attach detach record")
	AttachDetachUpdateError     = fmt.Errorf("error updating attach detach record")

	ShakingDecodeError     = fmt.Errorf("error decoding shaking record")
	ShakingValidationError = fmt.Errorf("error validating shaking record")
	ShakingCreateError     = fmt.Errorf("error creating shaking record")
	ShakingFetchError      = fmt.Errorf("error fetching shaking record")
	ShakingUpdateError     = fmt.Errorf("error updating shaking record")

	HeatingDecodeError     = fmt.Errorf("error decoding heating record")
	HeatingValidationError = fmt.Errorf("error validating heating record")
	HeatingCreateError     = fmt.Errorf("error creating heating record")
	HeatingFetchError      = fmt.Errorf("error fetching heating record")
	HeatingUpdateError     = fmt.Errorf("error updating heating record")

	TipDockingDecodeError     = fmt.Errorf("error decoding tip docking record")
	TipDockingValidationError = fmt.Errorf("error validating tip docking record")
	TipDockingCreateError     = fmt.Errorf("error creating tip docking record")
	TipDockingFetchError      = fmt.Errorf("error fetching tip docking record")
	TipDockingUpdateError     = fmt.Errorf("error updating tip docking record")

	TipOperationDecodeError     = fmt.Errorf("error decoding tip operation record")
	TipOperationValidationError = fmt.Errorf("error validating tip operation record")
	TipOperationCreateError     = fmt.Errorf("error creating tip operation record")
	TipOperationFetchError      = fmt.Errorf("error fetching tip operation record")
	TipOperationUpdateError     = fmt.Errorf("error updating tip operation record")

	UUIDParseError = fmt.Errorf("error parsing uuid")

	DelayRangeInvalid = fmt.Errorf("error invalid delay range allowed range is (0, 100]")
	
	SimulatorReservedDelayError = fmt.Errorf("delay is allowed only for simulator")

)

// Special errors which are in []byte format
var (
	DataMarshallingError = []byte(`error marshalling data`)
)
