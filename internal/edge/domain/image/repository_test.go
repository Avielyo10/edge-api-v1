package image

import (
	"context"
	"reflect"
	"testing"
	"time"

	"gorm.io/gorm"

	"github.com/Avielyo10/edge-api/internal/common/models"
	"github.com/Avielyo10/edge-api/internal/edge/domain/common"
	"github.com/redhatinsights/edge-api/config"
	"github.com/redhatinsights/platform-go-middlewares/identity"
)

func TestImage_MarshalGorm(t *testing.T) {
	config.Init()
	config.Get().Auth = true

	validSSHKey := "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDFjRxF1E73z1K9AjltDkuJyGUW3YluTEAW6PvHEZH6vnzNHI+cut716lGGRFHlYk1Fk51Q/92ZlynJ/HqByaK/MJppkQSL4x3KEm6s5ciwXbVEb3ct4waTgqxPD9gy7NN0uzbrhQMillb50yZgox6d9A/JmyRA1Dlai/esrlKfZ4wtSUl+CMsPoVxC6pIsh1YqUWE7S/dvXsQ8V+O7H0sdXAkZMg09kLUOQe3fliTMg6wppW+tb30g4MWAbHSrXksL1TpYjmP0M+stNetO2EIZ07bc8KpQhZybdM8LUhhPGuZXuKzIlwbkDI7C1yLv574wOYCjG/zk7Zu9qO7p6u8x valid@sshkey"
	validName, _ := common.NewName("valid-name")
	validVersion, _ := NewVersion(1)
	validRepos := NewRepos(
		NewRepo("test", "http://test.com"),
		NewRepo("test2", "http://test2.com"),
	)
	now := time.Now()
	ctx := context.WithValue(context.Background(), identity.Key, identity.XRHID{
		Identity: identity.Identity{
			AccountNumber: "0000000",
		},
	})
	account := common.DefaultAccount.String()
	type fields struct {
		ctx          context.Context
		cancel       context.CancelFunc
		uuid         string
		name         common.Name
		description  string
		timing       common.Time
		status       Status
		version      Version
		distribution Distribution
		user         User
		packages     Packages
		repos        Repos
		installer    Installer
		outputType   []OutputType
		tags         common.Tags
	}
	tests := []struct {
		name   string
		fields fields
		want   *models.Image
	}{
		{
			name: "should succeed",
			fields: fields{
				ctx:          ctx,
				cancel:       context.CancelFunc(func() {}),
				uuid:         "valid-uuid",
				name:         validName,
				description:  "valid-description",
				timing:       common.NewTime(now, now, time.Time{}),
				status:       Success,
				version:      validVersion,
				distribution: Distribution{"rhel8"},
				user: User{
					username: "valid-username",
					sshKey:   validSSHKey,
				},
				packages: Packages{requiredPackages: []Package{
					{"ansible"},
					{"rhc"},
				}, packages: []Package{
					{"vim"},
				}},
				repos:      validRepos,
				installer:  Installer{},
				outputType: []OutputType{TAR},
				tags:       common.NewTags("tag1", "tag2"),
			},
			want: &models.Image{
				Model: models.Model{
					CreatedAt: now,
					UpdatedAt: now,
					DeletedAt: gorm.DeletedAt{},
				},
				Account:      account,
				UUID:         "valid-uuid",
				Name:         "valid-name",
				Description:  "valid-description",
				Distribution: "rhel8",
				Status:       "success",
				Version:      1,
				User: models.User{
					Name:    "valid-username",
					SSHKey:  validSSHKey,
					Account: account,
				},
				Packages: []models.Package{
					{Name: "ansible", Account: account},
					{Name: "rhc", Account: account},
					{Name: "vim", Account: account},
				},
				Repos: []models.Repo{
					{Name: "test", URL: "http://test.com", Account: account},
					{Name: "test2", URL: "http://test2.com", Account: account},
				},
				Installer:   models.Installer{Account: account},
				OutputTypes: []string{"rhel-edge-commit"},
				Tags: []models.Tag{
					{Name: "tag1", Account: account},
					{Name: "tag2", Account: account},
				},
			},
		},
		{
			name: "invalid context",
			fields: fields{
				ctx:          context.Background(),
				cancel:       context.CancelFunc(func() {}),
				uuid:         "valid-uuid",
				name:         validName,
				description:  "valid-description",
				timing:       common.NewTime(now, now, time.Time{}),
				status:       Success,
				version:      validVersion,
				distribution: Distribution{"rhel8"},
				user: User{
					username: "valid-username",
					sshKey:   validSSHKey,
				},
				packages: Packages{requiredPackages: []Package{
					{"ansible"},
					{"rhc"},
				}, packages: []Package{
					{"vim"},
				}},
				repos:      validRepos,
				installer:  Installer{},
				outputType: []OutputType{TAR},
				tags:       common.NewTags("tag1", "tag2"),
			},
			want: nil,
		},
		{
			name:   "zero image",
			fields: fields{},
			want:   nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			image := Image{
				ctx:          tt.fields.ctx,
				cancel:       tt.fields.cancel,
				uuid:         tt.fields.uuid,
				name:         tt.fields.name,
				description:  tt.fields.description,
				timing:       tt.fields.timing,
				status:       tt.fields.status,
				version:      tt.fields.version,
				distribution: tt.fields.distribution,
				user:         tt.fields.user,
				packages:     tt.fields.packages,
				repos:        tt.fields.repos,
				installer:    tt.fields.installer,
				outputType:   tt.fields.outputType,
				tags:         tt.fields.tags,
			}
			if got := image.MarshalGorm(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Image.MarshalGorm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImage_MarshalRedis(t *testing.T) {
	validSSHKey := "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDFjRxF1E73z1K9AjltDkuJyGUW3YluTEAW6PvHEZH6vnzNHI+cut716lGGRFHlYk1Fk51Q/92ZlynJ/HqByaK/MJppkQSL4x3KEm6s5ciwXbVEb3ct4waTgqxPD9gy7NN0uzbrhQMillb50yZgox6d9A/JmyRA1Dlai/esrlKfZ4wtSUl+CMsPoVxC6pIsh1YqUWE7S/dvXsQ8V+O7H0sdXAkZMg09kLUOQe3fliTMg6wppW+tb30g4MWAbHSrXksL1TpYjmP0M+stNetO2EIZ07bc8KpQhZybdM8LUhhPGuZXuKzIlwbkDI7C1yLv574wOYCjG/zk7Zu9qO7p6u8x valid@sshkey"
	validName, _ := common.NewName("valid-name")
	validVersion, _ := NewVersion(1)
	validRepos := NewRepos(
		NewRepo("test", "http://test.com"),
		NewRepo("test2", "http://test2.com"),
	)
	now := time.Now()
	type fields struct {
		ctx          context.Context
		cancel       context.CancelFunc
		uuid         string
		name         common.Name
		description  string
		timing       common.Time
		status       Status
		version      Version
		distribution Distribution
		user         User
		packages     Packages
		repos        Repos
		installer    Installer
		outputType   []OutputType
		tags         common.Tags
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{
			name: "should succeed",
			fields: fields{
				ctx:          context.Background(),
				cancel:       context.CancelFunc(func() {}),
				uuid:         "valid-uuid",
				name:         validName,
				description:  "valid-description",
				timing:       common.NewTime(now, now, now),
				status:       Success,
				version:      validVersion,
				distribution: Distribution{"rhel8"},
				user: User{
					username: "valid-username",
					sshKey:   validSSHKey,
				},
				packages:   Packages{},
				repos:      validRepos,
				installer:  Installer{},
				outputType: []OutputType{},
				tags:       common.NewTags("tag1", "tag2"),
			},
			want: []byte(`{"uuid":"valid-uuid","name":"valid-name","description":"valid-description","status":"success","version":1,"distribution":"rhel8","user":{"username":"valid-username","ssh_key":"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDFjRxF1E73z1K9AjltDkuJyGUW3YluTEAW6PvHEZH6vnzNHI+cut716lGGRFHlYk1Fk51Q/92ZlynJ/HqByaK/MJppkQSL4x3KEm6s5ciwXbVEb3ct4waTgqxPD9gy7NN0uzbrhQMillb50yZgox6d9A/JmyRA1Dlai/esrlKfZ4wtSUl+CMsPoVxC6pIsh1YqUWE7S/dvXsQ8V+O7H0sdXAkZMg09kLUOQe3fliTMg6wppW+tb30g4MWAbHSrXksL1TpYjmP0M+stNetO2EIZ07bc8KpQhZybdM8LUhhPGuZXuKzIlwbkDI7C1yLv574wOYCjG/zk7Zu9qO7p6u8x valid@sshkey"},"packages":null,"repos":[{"name":"test","url":"http://test.com"},{"name":"test2","url":"http://test2.com"}],"installer":{},"tags":["tag1","tag2"],"created_at":"` + now.Format(time.RFC3339Nano) + `","updated_at":"` + now.Format(time.RFC3339Nano) + `","deleted_at":"` + now.Format(time.RFC3339Nano) + `"}`),
		},
		{
			name:   "should be nil",
			fields: fields{},
			want:   nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			image := Image{
				ctx:          tt.fields.ctx,
				cancel:       tt.fields.cancel,
				uuid:         tt.fields.uuid,
				name:         tt.fields.name,
				description:  tt.fields.description,
				timing:       tt.fields.timing,
				status:       tt.fields.status,
				version:      tt.fields.version,
				distribution: tt.fields.distribution,
				user:         tt.fields.user,
				packages:     tt.fields.packages,
				repos:        tt.fields.repos,
				installer:    tt.fields.installer,
				outputType:   tt.fields.outputType,
				tags:         tt.fields.tags,
			}
			if got := image.MarshalRedis(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Image.MarshalRedis() = %v, want %v", got, tt.want)
			}
		})
	}
}
