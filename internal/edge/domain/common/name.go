package common

import (
	"encoding/json"
	"errors"
	"strings"
)

// Name is a name of an entity, cannot be empty or fill with spaces.
type Name struct {
	name string
}

// ErrInvalidName is returned when the name is invalid.
var ErrInvalidName = errors.New("invalid name")

// NewName creates a new name.
func NewName(name string) (Name, error) {
	if !isValidName(name) {
		return Name{}, ErrInvalidName
	}
	return Name{name: name}, nil
}

// IsZero returns true if the name is empty.
func (n Name) IsZero() bool {
	return n == Name{}
}

// isValidName returns true if the name is valid.
func isValidName(name string) bool {
	return name != "" && strings.TrimSpace(name) != ""
}

// String returns the name as string.
func (n Name) String() string {
	return n.name
}

// MarshalJSON implements the json.Marshaler interface.
func (n Name) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (n *Name) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	n.name = s
	return nil
}
