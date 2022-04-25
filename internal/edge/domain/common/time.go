package common

import (
	"encoding/json"
	"time"
)

// imageTime struct contains the time of an image.
type Time struct {
	createdAt time.Time
	updatedAt time.Time
	deletedAt time.Time
}

// NewTime returns a new time.
func NewTime(createdAt, updatedAt, deletedAt time.Time) Time {
	return Time{createdAt: createdAt, updatedAt: updatedAt, deletedAt: deletedAt}
}

// CreatedAt returns the created at time of an image.
func (t Time) CreatedAt() time.Time {
	return t.createdAt
}

// UpdatedAt returns the updated at time of an image.
func (t Time) UpdatedAt() time.Time {
	return t.updatedAt
}

// DeletedAt returns the deleted at time of an image.
func (t Time) DeletedAt() time.Time {
	return t.deletedAt
}

// IsZero returns true if the time is empty.
func (t Time) IsZero() bool {
	return t == Time{}
}

// IsTimestampZero returns true if the timestamp is zero.

// MarshalJSON marshals the time to JSON.
func (t Time) MarshalJSON() ([]byte, error) {
	tmp := struct {
		CreatedAt string `json:"created_at,omitempty"`
		UpdatedAt string `json:"updated_at,omitempty"`
		DeletedAt string `json:"deleted_at,omitempty"`
	}{
		CreatedAt: t.createdAt.Format(time.RFC3339Nano),
		UpdatedAt: t.updatedAt.Format(time.RFC3339Nano),
		DeletedAt: t.deletedAt.Format(time.RFC3339Nano),
	}
	return json.Marshal(tmp)
}

// UnmarshalJSON unmarshals the time from JSON.
func (t *Time) UnmarshalJSON(data []byte) error {
	tmp := struct {
		CreatedAt string `json:"created_at,omitempty"`
		UpdatedAt string `json:"updated_at,omitempty"`
		DeletedAt string `json:"deleted_at,omitempty"`
	}{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	if tmp.CreatedAt != "" {
		if createdAt, err := time.Parse(time.RFC3339Nano, tmp.CreatedAt); err != nil {
			return err
		} else {
			t.createdAt = createdAt
		}
	}
	if tmp.UpdatedAt != "" {
		if updatedAt, err := time.Parse(time.RFC3339Nano, tmp.UpdatedAt); err != nil {
			return err
		} else {
			t.updatedAt = updatedAt
		}
	}
	if tmp.DeletedAt != "" {
		if deletedAt, err := time.Parse(time.RFC3339Nano, tmp.DeletedAt); err != nil {
			return err
		} else {
			t.deletedAt = deletedAt
		}
	}
	return nil
}
