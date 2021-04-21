package responses

import(
	"fmt"
)

var(
	PiercingDecodingError = fmt.Errorf("error decoding piercing record")
	PiercingValidationError = fmt.Errorf("error validating piercing record")
	PiercingCreationError = fmt.Errorf("error creating piercing record")

)

// Special errors which are in []byte format
var (
	DataMarshallingError = []byte(`error marshalling data`)
)