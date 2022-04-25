package image

import (
	"encoding/json"
	"errors"
	"strings"
)

// ErrInvalidDist returns an error for an invalid distribution.
var ErrInvalidDist = errors.New("invalid distribution")

// Distribution is the type of the distribution of an image.
type Distribution struct {
	dist string
}

// NewDistribution creates a new distribution.
func NewDistribution(dist string) Distribution {
	if !isValidDist(dist) {
		return Distribution{}
	}
	return Distribution{dist: dist}
}

// String returns the string representation of a distribution.
func (d Distribution) String() string {
	return d.dist
}

// IsZero returns true if the distribution is empty.
func (d Distribution) IsZero() bool {
	return d == Distribution{}
}

// isValidDist returns true if the distribution is valid.
func isValidDist(dist string) bool {
	return dist != "" && strings.TrimSpace(dist) != ""
}

// MarshalJSON creates a custom json marshaller.
func (d Distribution) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

// UnmarshalJSON creates a custom json unmarshaller.
func (d *Distribution) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	if NewDistribution(s).IsZero() {
		return ErrInvalidDist
	}
	d.dist = s
	return nil
}
