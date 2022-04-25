package image

import (
	"encoding/json"

	"github.com/Avielyo10/edge-api/internal/common/models"
)

// Installer defines the model for a ISO installer
type Installer struct {
	isoURL       string
	composeJobID string
	checksum     string
}

// NewInstaller returns a new installer
func NewInstaller(isoURL, composeJobID, checksum string) Installer {
	return Installer{
		isoURL:       isoURL,
		composeJobID: composeJobID,
		checksum:     checksum,
	}
}

// ISOURL returns the iso url of the installer
func (i Installer) ISOURL() string {
	return i.isoURL
}

// ComposeJobID returns the compose job id of the installer
func (i Installer) ComposeJobID() string {
	return i.composeJobID
}

// Checksum returns the checksum of the installer
func (i Installer) Checksum() string {
	return i.checksum
}

// MarshalGorm marshals the installer
func (i Installer) MarshalGorm() *models.Installer {
	return &models.Installer{
		ISOURL:       i.ISOURL(),
		ComposeJobID: i.ComposeJobID(),
		Checksum:     i.Checksum(),
	}
}

// UnmarshalGorm unmarshals the installer
func (i Installer) UnmarshalGorm(in *models.Installer) Installer {
	if in != nil {
		return Installer{
			isoURL:       in.ISOURL,
			composeJobID: in.ComposeJobID,
			checksum:     in.Checksum,
		}
	}
	return Installer{}
}

// MarshalJSON creates a custom json marshaller
func (i Installer) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ISOURL       string `json:"iso_url,omitempty"`
		ComposeJobID string `json:"compose_job_id,omitempty"`
		Checksum     string `json:"checksum,omitempty"`
	}{
		ISOURL:       i.ISOURL(),
		ComposeJobID: i.ComposeJobID(),
		Checksum:     i.Checksum(),
	})
}

// UnmarshalJSON creates a custom json unmarshaller
func (i *Installer) UnmarshalJSON(data []byte) error {
	var tmp struct {
		ISOURL       string `json:"iso_url,omitempty"`
		ComposeJobID string `json:"compose_job_id,omitempty"`
		Checksum     string `json:"checksum,omitempty"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	i.isoURL = tmp.ISOURL
	i.composeJobID = tmp.ComposeJobID
	i.checksum = tmp.Checksum
	return nil
}
