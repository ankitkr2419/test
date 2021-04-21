package responses

import(
	"fmt"
)

var(
	PiercingDecodeError = fmt.Errorf("error decoding piercing record")
	PiercingValidationError = fmt.Errorf("error validating piercing record")
	PiercingCreateError = fmt.Errorf("error creating piercing record")
	PiercingFetchError = fmt.Errorf("error fetching piercing record")
	PiercingUpdateError = fmt.Errorf("error updating piercing record")
	UUIDParseError = fmt.Errorf("error parsing uuid")
)

// Special errors which are in []byte format
var (
	DataMarshallingError = []byte(`error marshalling data`)
)