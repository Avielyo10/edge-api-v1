package image

import (
	"encoding/json"
	"errors"

	"golang.org/x/crypto/ssh"

	"github.com/Avielyo10/edge-api/internal/common/models"
)

// User defines the model for a ISO user
type User struct {
	username string
	sshKey   string
}

// ErrInvalidUser returns an error for an invalid user
var ErrInvalidUser = errors.New("invalid user")

// NewUser returns a new user
func NewUser(username, sshKey string) (User, error) {
	if !isValidUser(username, sshKey) {
		return User{}, ErrInvalidUser
	}
	return User{username: username, sshKey: sshKey}, nil
}

// ValidSSHKey returns true if the ssh key is valid
func (u User) ValidSSHKey() bool {
	_, _, _, _, err := ssh.ParseAuthorizedKey([]byte(u.sshKey))
	return err == nil
}

// ValidUsername returns true if the username is valid
func (u User) ValidUsername() bool {
	return u.username != ""
}

// IsZero returns true if the user is empty
func (u User) IsZero() bool {
	return u == User{}
}

// Valid returns true if the user is valid
func isValidUser(username, sshKey string) bool {
	u := User{username: username, sshKey: sshKey}
	return u.ValidUsername() && u.ValidSSHKey()
}

// Username returns the username
func (u User) Username() string {
	return u.username
}

// SSHKey returns the ssh key
func (u User) SSHKey() string {
	return u.sshKey
}

// MarshalJSON creates a custom json marshaller
func (u User) MarshalJSON() ([]byte, error) {
	return []byte(`{"username":"` + u.username + `","ssh_key":"` + u.sshKey + `"}`), nil
}

// UnmarshalJSON creates a custom json unmarshaller
func (u *User) UnmarshalJSON(data []byte) error {
	var v struct {
		Username string `json:"username"`
		SSHKey   string `json:"ssh_key"`
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	if !isValidUser(v.Username, v.SSHKey) {
		return ErrInvalidUser
	}
	*u = User{username: v.Username, sshKey: v.SSHKey}
	return nil
}

// MarshalGorm marshals the user to a gorm model.
func (u User) MarshalGorm() *models.User {
	return &models.User{
		Name:   u.username,
		SSHKey: u.sshKey,
	}
}
