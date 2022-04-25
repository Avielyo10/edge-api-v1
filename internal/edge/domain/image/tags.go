package image

import "github.com/Avielyo10/edge-api/internal/edge/domain/common"

// AddTag adds a tag to the image.
func (i *Image) AddTag(tag ...common.Tag) {
	i.tags.Add(tag...)
}

// RemoveTag removes a tag from the image.
func (i *Image) RemoveTag(tag ...common.Tag) {
	i.tags.Remove(tag...)
}
