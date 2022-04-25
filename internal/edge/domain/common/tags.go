package common

import (
	"encoding/json"
	"strings"

	"github.com/Avielyo10/edge-api/internal/common/models"
)

// Tag is a tag of an image.
type Tag struct {
	tag string
}

// Tags are list of tags.
type Tags struct {
	tags []Tag
}

// NewTag returns a new tag.
func NewTag(tag string) Tag {
	if !isValidTag(tag) {
		return Tag{}
	}
	return Tag{tag: tag}
}

// NewTags returns a new list of tags.
func NewTags(tags ...string) Tags {
	if len(tags) == 0 {
		return Tags{}
	}
	var newTags Tags
	for _, tag := range tags {
		newTags.Add(NewTag(tag))
	}
	return newTags
}

// IsZero returns true if the tag is empty.
func (t Tag) IsZero() bool {
	return t == Tag{}
}

// isValidTag returns true if the tag is valid.
func isValidTag(name string) bool {
	return name != "" && strings.TrimSpace(name) != ""
}

// Add adds a tag to the list.
func (t *Tags) Add(tags ...Tag) {
	for _, tag := range tags {
		if !tag.IsZero() {
			t.tags = append(t.tags, tag)
		}
	}
}

// Remove removes a tag from the list if exist.
func (t *Tags) Remove(tags ...Tag) {
	// create a hashset of tags
	tagSet := make(map[string]bool)
	for _, tag := range t.tags {
		// add all tags to the hashset
		tagSet[tag.String()] = true
	}
	// remove all tags from the hashset
	for _, tag := range tags {
		if tag.IsZero() {
			continue
		}
		delete(tagSet, tag.String())
	}
	// create a new list of tags
	newTags := make([]Tag, 0, len(tagSet))
	for tag := range tagSet {
		newTags = append(newTags, NewTag(tag))
	}
	t.tags = newTags
}

// String returns the string representation of the tag.
func (t Tag) String() string {
	return t.tag
}

// StringArray returns a list of tags as strings.
func (t Tags) StringArray() []string {
	tags := make([]string, 0, len(t.tags))
	for _, tag := range t.tags {
		tags = append(tags, tag.String())
	}
	return tags
}

// Tags returns the list of tags.
func (t Tags) Tags() []Tag {
	return t.tags
}

// MarshalJSON creates a custom json marshaller
func (t Tags) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.StringArray())
}

// UnmarshalJSON creates a custom json unmarshaller
func (t *Tags) UnmarshalJSON(data []byte) error {
	var tags []string
	if err := json.Unmarshal(data, &tags); err != nil {
		return err
	}
	var newTags []Tag
	for _, tag := range tags {
		newTag := NewTag(tag)
		if !newTag.IsZero() {
			newTags = append(newTags, newTag)
		}
	}
	t.tags = newTags
	return nil
}

// MarshalGorm marshals the tags to a gorm model.
func (t Tags) MarshalGorm(account string) []models.Tag {
	var tags []models.Tag
	for _, tag := range t.Tags() {
		tags = append(tags, models.Tag{
			Account: account,
			Name:    tag.String(),
		})
	}
	return tags
}
