package image

import "errors"

// OutputType is the type of the output of the image.
type OutputType struct {
	out string
}

// enum
var (
	ISO = OutputType{"rhel-edge-installer"}
	TAR = OutputType{"rhel-edge-commit"}
)

// ErrInvalidOutputType is returned when the output type is invalid.
var ErrInvalidOutputType = errors.New("invalid output type")

// available output types
var outputTypes = []OutputType{ISO, TAR}

// NewOutputTypeFromString creates a new output type from a string.
func NewOutputTypeFromString(out string) (OutputType, error) {
	for _, outputType := range outputTypes {
		if outputType.out == out {
			return outputType, nil
		}
	}
	return OutputType{}, ErrInvalidOutputType
}

// NewOutputType returns a new output type, TAR by default.
func NewOutputType(out ...string) []OutputType {
	if len(out) == 0 {
		return []OutputType{TAR}
	}
	// make a set of output types
	outputTypes := make(map[OutputType]bool)
	// always add TAR
	outputTypes[TAR] = true
	for _, outputType := range out {
		outputType, err := NewOutputTypeFromString(outputType)
		if err != nil {
			continue
		}
		outputTypes[outputType] = true
	}
	// convert to slice
	outputTypesSlice := make([]OutputType, 0, len(outputTypes))
	for outputType := range outputTypes {
		outputTypesSlice = append(outputTypesSlice, outputType)
	}
	return outputTypesSlice
}

// IsZero returns true if the output type is empty.
func (o OutputType) IsZero() bool {
	return o == OutputType{}
}

// String returns the string representation of an output type.
func (o OutputType) String() string {
	return o.out
}

// IsISO returns true if the output type is ISO.
func (o OutputType) IsISO() bool {
	return o == ISO
}

// IsTAR returns true if the output type is TAR.
func (o OutputType) IsTAR() bool {
	return o == TAR
}

// IsAvailableOutputType returns true if the output type is available.
func IsAvailableOutputType(out string) bool {
	for _, outputType := range outputTypes {
		if outputType.out == out {
			return true
		}
	}
	return false
}

// MarshalJSON marshals the output type to JSON.
func (o OutputType) MarshalJSON() ([]byte, error) {
	return []byte(`"` + o.out + `"`), nil
}

// UnmarshalJSON unmarshals the output type from JSON.
func (o *OutputType) UnmarshalJSON(data []byte) error {
	if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		return ErrInvalidOutputType
	}
	o.out = string(data[1 : len(data)-1])
	return nil
}
