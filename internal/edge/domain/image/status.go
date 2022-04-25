package image

import "errors"

// Define the available statuses.
var (
	Success  = Status{"success"}
	Building = Status{"building"}
	Error    = Status{"error"}
)

// All available statuses.
var availableStatuses = []Status{
	Success,
	Building,
	Error,
}

// Status errors
var (
	ErrAlreadyBuilding = errors.New("image already building")
	ErrInvalidStatus   = errors.New("invalid status")
)

// Status represents the status of an image.
// Using struct instead of `type Status string` to allow us to ensure,
// that we have full control of what values are possible.
type Status struct {
	state string
}

// NewStatusFromString creates a new status from a string.
func NewStatusFromString(state string) (Status, error) {
	for _, status := range availableStatuses {
		if status.state == state {
			return status, nil
		}
	}
	return Status{}, ErrInvalidStatus
}

// NewStatus returns a new status, building by default.
func NewStatus() Status {
	return Building
}

// IsZero returns true if the status is empty.
func (s Status) IsZero() bool {
	return s == Status{}
}

// String returns the string representation of a status.
func (s Status) String() string {
	return s.state
}

// IsSuccess returns true if the status is success.
func (s Status) IsSuccess() bool {
	return s == Success
}

// IsBuilding returns true if the status is building.
func (s Status) IsBuilding() bool {
	return s == Building
}

// IsError returns true if the status is error.
func (s Status) IsError() bool {
	return s == Error
}

// MarshalJSON marshals the status to JSON.
func (s Status) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.state + `"`), nil
}

// UnmarshalJSON unmarshals the status from JSON.
func (s *Status) UnmarshalJSON(data []byte) error {
	if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		return ErrInvalidStatus
	}
	state, err := NewStatusFromString(string(data[1 : len(data)-1]))
	if err != nil {
		return err
	}
	*s = state
	return nil
}
