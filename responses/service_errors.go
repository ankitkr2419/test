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

	UUIDParseError = fmt.Errorf("error parsing uuid")
)

// Special errors which are in []byte format
var (
	DataMarshallingError = []byte(`error marshalling data`)
)
