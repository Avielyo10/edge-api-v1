package image

import (
	"errors"
	"strconv"
)

var (
	ErrInvalidVersion = errors.New("invalid version")
)

// Version is a version of an image.
type Version struct {
	number uint
}

// IsZero returns true if the version is empty.
func (v Version) IsZero() bool {
	return v == Version{}
}

// Update increments the version of an image.
func (v *Version) Update() {
	v.number++
}

// Rollback decrements the version of an image.
func (v *Version) Rollback() {
	if v.number > 1 {
		v.number--
	}
}

// NewVersion creates a new version of an image.
func NewVersion(number uint) (Version, error) {
	if number == 0 {
		return Version{}, ErrInvalidVersion
	}
	return Version{number}, nil
}

// Uint returns the version as uint64.
func (v Version) Uint() uint {
	return v.number
}

// MarshalJSON marshals the version to JSON.
func (v Version) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatUint(uint64(v.number), 10)), nil
}

// UnmarshalJSON unmarshals the version from JSON.
func (v *Version) UnmarshalJSON(data []byte) error {
	number, err := strconv.ParseUint(string(data), 10, 64)
	if err != nil {
		return err
	}
	*v, err = NewVersion(uint(number))
	return err
}
