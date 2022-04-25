package image

import "time"

// Upgrade updates the image, implementing the UpdateInterface interface.
func (image *Image) Upgrade() error {
	if image.status.IsBuilding() {
		return ErrAlreadyBuilding
	}
	image.status = Building
	image.version.Update()
	image.WithTimeout(90 * time.Minute) // 1.5h timeout
	// TODO: implement this with image-builder client.
	return nil
}

// Rollback rolls back the image, implementing the UpdateInterface interface.
func (image *Image) Rollback() error {
	image.Cancel()
	image.status = Success // image rolled back already present
	image.version.Rollback()
	return nil
}

// CheckForUpdate checks for updates, implementing the UpdateInterface interface.
func (image *Image) CheckForUpdate() error {
	// TODO: implement this with image-builder client.
	return nil
}

// IsSuccessful returns true if the image is successfully updated, implementing the UpdateInterface interface.
func (image Image) IsSuccessful() bool {
	return image.status.IsSuccess()
}
