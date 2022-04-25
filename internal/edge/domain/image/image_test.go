package image

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/Avielyo10/edge-api/internal/edge/domain/common"
	"github.com/redhatinsights/edge-api/config"
	"github.com/redhatinsights/platform-go-middlewares/identity"
)

func TestNewImage(t *testing.T) {
	validSSHKey := "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDFjRxF1E73z1K9AjltDkuJyGUW3YluTEAW6PvHEZH6vnzNHI+cut716lGGRFHlYk1Fk51Q/92ZlynJ/HqByaK/MJppkQSL4x3KEm6s5ciwXbVEb3ct4waTgqxPD9gy7NN0uzbrhQMillb50yZgox6d9A/JmyRA1Dlai/esrlKfZ4wtSUl+CMsPoVxC6pIsh1YqUWE7S/dvXsQ8V+O7H0sdXAkZMg09kLUOQe3fliTMg6wppW+tb30g4MWAbHSrXksL1TpYjmP0M+stNetO2EIZ07bc8KpQhZybdM8LUhhPGuZXuKzIlwbkDI7C1yLv574wOYCjG/zk7Zu9qO7p6u8x valid@sshkey"
	validName, _ := common.NewName("valid-name")
	validVersion, _ := NewVersion(1)
	validRepos := NewRepos(
		NewRepo("test", "http://test.com"),
		NewRepo("test2", "http://test2.com"),
	)
	type args struct {
		uuid         string
		name         string
		description  string
		distribution string
		status       string
		username     string
		sshKey       string
		outputType   []string
		tags         []string
		packages     []string
		version      uint
		repos        []interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    Image
		wantErr bool
	}{
		{
			name: "should succeed",
			args: args{
				uuid:       "00000000-0000-0000-0000-000000000000",
				name:       validName.String(),
				status:     "success",
				username:   "valid user",
				sshKey:     validSSHKey,
				outputType: []string{TAR.String()},
				tags:       []string{"test"},
				packages:   []string{"test"},
				version:    1,
				repos: []interface{}{
					map[string]interface{}{
						"name": "test",
						"url":  "http://test.com",
					},
					map[string]interface{}{
						"name": "test2",
						"url":  "http://test2.com",
					},
				},
			},
			want: Image{
				uuid:   "00000000-0000-0000-0000-000000000000",
				name:   validName,
				status: Success,
				user: User{
					username: "valid user",
					sshKey:   validSSHKey,
				},
				outputType: []OutputType{
					TAR,
				},
				tags: common.NewTags("test"),
				packages: NewPackages(
					"test",
				),
				version: validVersion,
				repos:   validRepos,
			},
			wantErr: false,
		},
		{
			name: "should fail, bad name",
			args: args{
				uuid:       "00000000-0000-0000-0000-000000000000",
				name:       "",
				status:     "success",
				username:   "valid user",
				sshKey:     validSSHKey,
				outputType: []string{TAR.String()},
				tags:       []string{"test"},
				packages:   []string{"test"},
				version:    1,
				repos: []interface{}{
					map[string]interface{}{
						"name": "test",
						"url":  "http://test.com",
					},
					map[string]interface{}{
						"name": "test2",
						"url":  "http://test2.com",
					},
				},
			},
			want: Image{
				uuid:   "00000000-0000-0000-0000-000000000000",
				name:   validName,
				status: Success,
				user: User{
					username: "valid user",
					sshKey:   validSSHKey,
				},
				outputType: []OutputType{
					TAR,
				},
				tags: common.NewTags("test"),
				packages: NewPackages(
					"test",
				),
				version: validVersion,
				repos:   validRepos,
			},
			wantErr: true,
		},
		{
			name: "should fail, bad user",
			args: args{
				uuid:       "00000000-0000-0000-0000-000000000000",
				name:       validName.String(),
				status:     "success",
				username:   "",
				sshKey:     validSSHKey,
				outputType: []string{TAR.String()},
				tags:       []string{"test"},
				packages:   []string{"test"},
				version:    1,
				repos: []interface{}{
					map[string]interface{}{
						"name": "test",
						"url":  "http://test.com",
					},
					map[string]interface{}{
						"name": "test2",
						"url":  "http://test2.com",
					},
				},
			},
			want: Image{
				uuid:   "00000000-0000-0000-0000-000000000000",
				name:   validName,
				status: Success,
				user: User{
					username: "valid user",
					sshKey:   validSSHKey,
				},
				outputType: []OutputType{
					TAR,
				},
				tags: common.NewTags("test"),
				packages: NewPackages(
					"test",
				),
				version: validVersion,
				repos:   validRepos,
			},
			wantErr: true,
		},
		{
			name: "should fail, bad version",
			args: args{
				uuid:       "00000000-0000-0000-0000-000000000000",
				name:       validName.String(),
				status:     "success",
				username:   "valid user",
				sshKey:     validSSHKey,
				outputType: []string{TAR.String()},
				tags:       []string{"test"},
				packages:   []string{"test"},
				version:    0,
				repos: []interface{}{
					map[string]interface{}{
						"name": "test",
						"url":  "http://test.com",
					},
					map[string]interface{}{
						"name": "test2",
						"url":  "http://test2.com",
					},
				},
			},
			want: Image{
				uuid:   "00000000-0000-0000-0000-000000000000",
				name:   validName,
				status: Success,
				user: User{
					username: "valid user",
					sshKey:   validSSHKey,
				},
				outputType: []OutputType{
					TAR,
				},
				tags: common.NewTags("test"),
				packages: NewPackages(
					"test",
				),
				version: validVersion,
				repos:   validRepos,
			},
			wantErr: true,
		},
		{
			name: "should fail, bad status",
			args: args{
				uuid:       "00000000-0000-0000-0000-000000000000",
				name:       validName.String(),
				status:     "invalid",
				username:   "valid user",
				sshKey:     validSSHKey,
				outputType: []string{TAR.String()},
				tags:       []string{"test"},
				packages:   []string{"test"},
				version:    1,
				repos: []interface{}{
					map[string]interface{}{
						"name": "test",
						"url":  "http://test.com",
					},
					map[string]interface{}{
						"name": "test2",
						"url":  "http://test2.com",
					},
				},
			},
			want: Image{
				uuid:   "00000000-0000-0000-0000-000000000000",
				name:   validName,
				status: Success,
				user: User{
					username: "valid user",
					sshKey:   validSSHKey,
				},
				outputType: []OutputType{
					TAR,
				},
				tags: common.NewTags("test"),
				packages: NewPackages(
					"test",
				),
				version: validVersion,
				repos:   validRepos,
			},
			wantErr: true,
		},
		{
			name: "should fail, bad repos",
			args: args{
				uuid:       "00000000-0000-0000-0000-000000000000",
				name:       validName.String(),
				status:     "success",
				username:   "valid user",
				sshKey:     validSSHKey,
				outputType: []string{TAR.String()},
				tags:       []string{"test"},
				packages:   []string{"test"},
				version:    1,
				repos:      []interface{}{`"`},
			},
			want: Image{
				uuid:   "00000000-0000-0000-0000-000000000000",
				name:   validName,
				status: Success,
				user: User{
					username: "valid user",
					sshKey:   validSSHKey,
				},
				outputType: []OutputType{
					TAR,
				},
				tags: common.NewTags("test"),
				packages: NewPackages(
					"test",
				),
				version: validVersion,
				repos:   validRepos,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := NewImage(tt.args.uuid, tt.args.name, tt.args.description, tt.args.distribution, tt.args.status, tt.args.username, tt.args.sshKey, tt.args.outputType, tt.args.tags, tt.args.packages, tt.args.version, tt.args.repos)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got.ctx = nil    // ignore context
			got.cancel = nil // ignore context
			if !reflect.DeepEqual(got, tt.want) && !tt.wantErr {
				t.Errorf("NewImage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewImageWithContext(t *testing.T) {
	ctx := context.Background()
	validSSHKey := "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDFjRxF1E73z1K9AjltDkuJyGUW3YluTEAW6PvHEZH6vnzNHI+cut716lGGRFHlYk1Fk51Q/92ZlynJ/HqByaK/MJppkQSL4x3KEm6s5ciwXbVEb3ct4waTgqxPD9gy7NN0uzbrhQMillb50yZgox6d9A/JmyRA1Dlai/esrlKfZ4wtSUl+CMsPoVxC6pIsh1YqUWE7S/dvXsQ8V+O7H0sdXAkZMg09kLUOQe3fliTMg6wppW+tb30g4MWAbHSrXksL1TpYjmP0M+stNetO2EIZ07bc8KpQhZybdM8LUhhPGuZXuKzIlwbkDI7C1yLv574wOYCjG/zk7Zu9qO7p6u8x valid@sshkey"
	validName, _ := common.NewName("valid-name")
	validVersion, _ := NewVersion(1)
	validRepos := NewRepos(
		NewRepo("test", "http://test.com"),
		NewRepo("test2", "http://test2.com"),
	)
	type args struct {
		ctx          context.Context
		uuid         string
		name         string
		description  string
		distribution string
		status       string
		username     string
		sshKey       string
		outputType   []string
		tags         []string
		packages     []string
		version      uint
		repos        []interface{}
	}
	tests := []struct {
		name      string
		args      args
		wantImage Image
		wantErr   bool
	}{
		{
			name: "should succeed",
			args: args{
				ctx:          ctx,
				uuid:         "00000000-0000-0000-0000-000000000000",
				name:         validName.String(),
				status:       "success",
				username:     "valid user",
				sshKey:       validSSHKey,
				distribution: "rhel8",
				outputType:   []string{TAR.String()},
				tags:         []string{"test"},
				packages:     []string{"test"},
				version:      1,
				repos: []interface{}{
					map[string]interface{}{
						"name": "test",
						"url":  "http://test.com",
					},
					map[string]interface{}{
						"name": "test2",
						"url":  "http://test2.com",
					},
				},
			},
			wantImage: Image{
				ctx:    ctx,
				uuid:   "00000000-0000-0000-0000-000000000000",
				name:   validName,
				status: Success,
				user: User{
					username: "valid user",
					sshKey:   validSSHKey,
				},
				outputType: []OutputType{
					TAR,
				},
				distribution: NewDistribution("rhel8"),
				tags:         common.NewTags("test"),
				packages: NewPackages(
					"test",
				),
				version: validVersion,
				repos:   validRepos,
			},
			wantErr: false,
		},
		{
			name: "nil context",
			args: args{
				ctx:        nil,
				uuid:       "00000000-0000-0000-0000-000000000000",
				name:       validName.String(),
				status:     "success",
				username:   "valid user",
				sshKey:     validSSHKey,
				outputType: []string{TAR.String()},
				tags:       []string{"test"},
				packages:   []string{"test"},
				version:    1,
				repos: []interface{}{
					map[string]interface{}{
						"name": "test",
						"url":  "http://test.com",
					},
					map[string]interface{}{
						"name": "test2",
						"url":  "http://test2.com",
					},
				},
			},
			wantImage: Image{},
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotImage, err := NewImageWithContext(tt.args.ctx, tt.args.uuid, tt.args.name,
				tt.args.description, tt.args.distribution, tt.args.status, tt.args.username, tt.args.sshKey, tt.args.outputType, tt.args.tags, tt.args.packages, tt.args.version, tt.args.repos)
			tt.wantImage.ctx = nil    // ignore context
			tt.wantImage.cancel = nil // ignore context
			gotImage.ctx = nil        // ignore context
			gotImage.cancel = nil     // ignore context
			if (err != nil) != tt.wantErr {
				t.Errorf("NewImageWithContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotImage, tt.wantImage) {
				t.Errorf("NewImageWithContext() = %v, want %v", gotImage, tt.wantImage)
			}
		})
	}
}

func TestImage_IsZero(t *testing.T) {
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
		want   bool
	}{
		{
			name: "should succeed",
			fields: fields{
				ctx:          nil,
				cancel:       nil,
				uuid:         "",
				name:         common.Name{},
				description:  "",
				timing:       common.Time{},
				status:       Status{},
				version:      Version{},
				distribution: Distribution{},
				user:         User{},
				packages:     Packages{},
				repos:        Repos{},
				installer:    Installer{},
				outputType:   []OutputType{},
				tags:         common.Tags{},
			},
			want: true,
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
			if got := image.IsZero(); got != tt.want {
				t.Errorf("Image.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImage_Account(t *testing.T) {
	validSSHKey := "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDFjRxF1E73z1K9AjltDkuJyGUW3YluTEAW6PvHEZH6vnzNHI+cut716lGGRFHlYk1Fk51Q/92ZlynJ/HqByaK/MJppkQSL4x3KEm6s5ciwXbVEb3ct4waTgqxPD9gy7NN0uzbrhQMillb50yZgox6d9A/JmyRA1Dlai/esrlKfZ4wtSUl+CMsPoVxC6pIsh1YqUWE7S/dvXsQ8V+O7H0sdXAkZMg09kLUOQe3fliTMg6wppW+tb30g4MWAbHSrXksL1TpYjmP0M+stNetO2EIZ07bc8KpQhZybdM8LUhhPGuZXuKzIlwbkDI7C1yLv574wOYCjG/zk7Zu9qO7p6u8x valid@sshkey"
	validName, _ := common.NewName("valid-name")
	validVersion, _ := NewVersion(1)
	validRepos := NewRepos(
		NewRepo("test", "http://test.com"),
		NewRepo("test2", "http://test2.com"),
	)
	config.Init()

	validReq, _ := http.NewRequest("GET", "/", nil)
	validReq = validReq.WithContext(context.WithValue(validReq.Context(), identity.Key, identity.XRHID{
		Identity: identity.Identity{
			AccountNumber: "0000000",
		},
	}))
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
		name    string
		fields  fields
		want    common.Account
		wantErr bool
	}{
		{
			name: "should succeed",
			fields: fields{
				ctx:    validReq.Context(),
				uuid:   "00000000-0000-0000-0000-000000000000",
				name:   validName,
				status: Success,
				user: User{
					username: "valid user",
					sshKey:   validSSHKey,
				},
				outputType: []OutputType{
					TAR,
				},
				tags:     common.Tags{},
				packages: Packages{},
				version:  validVersion,
				repos:    validRepos,
			},
			want:    common.DefaultAccount,
			wantErr: false,
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
			got, err := image.Account()
			if (err != nil) != tt.wantErr {
				t.Errorf("Image.Account() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Image.Account() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImage_UUID(t *testing.T) {
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
		want   string
	}{
		{
			name: "should succeed",
			fields: fields{
				uuid: "00000000-0000-0000-0000-000000000000",
			},
			want: "00000000-0000-0000-0000-000000000000",
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
			if got := image.UUID(); got != tt.want {
				t.Errorf("Image.UUID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImage_Version(t *testing.T) {
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
		want   Version
	}{
		{
			name: "should succeed",
			fields: fields{
				version: Version{1},
			},
			want: Version{1},
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
			if got := image.Version(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Image.Version() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImage_User(t *testing.T) {
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
		want   User
	}{
		{
			name: "should succeed",
			fields: fields{
				user: User{
					username: "test",
				},
			},
			want: User{
				username: "test",
			},
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
			if got := image.User(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Image.User() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImage_Status(t *testing.T) {
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
		want   Status
	}{
		{
			name: "should succeed",
			fields: fields{
				status: Success,
			},
			want: Success,
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
			if got := image.Status(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Image.Status() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImage_Context(t *testing.T) {
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
		want   context.Context
	}{
		{
			name: "should succeed",
			fields: fields{
				ctx: context.Background(),
			},
			want: context.Background(),
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
			if got := image.Context(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Image.Context() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImage_Packages(t *testing.T) {
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
		want   Packages
	}{
		{
			name: "should succeed",
			fields: fields{
				packages: Packages{packages: []Package{
					{"vim"},
				}},
			},
			want: Packages{packages: []Package{
				{"vim"},
			}},
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
			if got := image.Packages(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Image.Packages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImage_Repos(t *testing.T) {
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
		want   Repos
	}{
		{
			name: "should succeed",
			fields: fields{
				repos: Repos{repos: []*Repo{}},
			},
			want: Repos{repos: []*Repo{}},
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
			if got := image.Repos(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Image.Repos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImage_Tags(t *testing.T) {
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
		want   common.Tags
	}{
		{
			name: "should succeed",
			fields: fields{
				tags: common.Tags{},
			},
			want: common.Tags{},
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
			if got := image.Tags(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Image.Tags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImage_Name(t *testing.T) {
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
		want   common.Name
	}{
		{
			name: "should succeed",
			fields: fields{
				name: common.Name{},
			},
			want: common.Name{},
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
			if got := image.Name(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Image.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImage_Description(t *testing.T) {
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
		want   string
	}{
		{
			name: "should succeed",
			fields: fields{
				description: "blah",
			},
			want: "blah",
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
			if got := image.Description(); got != tt.want {
				t.Errorf("Image.Description() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImage_Distribution(t *testing.T) {
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
		want   Distribution
	}{
		{
			name: "should succeed",
			fields: fields{
				distribution: Distribution{},
			},
			want: Distribution{},
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
			if got := image.Distribution(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Image.Distribution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImage_Installer(t *testing.T) {
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
		want   Installer
	}{
		{
			name: "should succeed",
			fields: fields{
				installer: Installer{},
			},
			want: Installer{},
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
			if got := image.Installer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Image.Installer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImage_OutputTypes(t *testing.T) {
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
		want   []OutputType
	}{
		{
			name: "should succeed",
			fields: fields{
				outputType: []OutputType{TAR},
			},
			want: []OutputType{TAR},
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
			if got := image.OutputTypes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Image.OutputTypes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImage_CreatedAt(t *testing.T) {
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
		want   time.Time
	}{
		{
			name: "should succeed",
			fields: fields{
				timing: common.Time{},
			},
			want: time.Time{},
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
			if got := image.CreatedAt(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Image.CreatedAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImage_UpdatedAt(t *testing.T) {
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
		want   time.Time
	}{
		{
			name: "should succeed",
			fields: fields{
				timing: common.Time{},
			},
			want: time.Time{},
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
			if got := image.UpdatedAt(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Image.UpdatedAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImage_DeletedAt(t *testing.T) {
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
		want   time.Time
	}{
		{
			name: "should succeed",
			fields: fields{
				timing: common.Time{},
			},
			want: time.Time{},
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
			if got := image.DeletedAt(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Image.DeletedAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImage_Cancel(t *testing.T) {
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
	}{
		{
			name: "should succeed",
			fields: fields{
				ctx:    context.Background(),
				cancel: func() {},
			},
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
			image.Cancel()
		})
	}
}

func TestImage_SetTime(t *testing.T) {
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
	type args struct {
		timing common.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   common.Time
	}{
		{
			name: "should succeed",
			fields: fields{
				ctx:    context.Background(),
				timing: common.Time{},
			},
			args: args{
				timing: common.NewTime(now, now, now),
			},
			want: common.NewTime(now, now, now),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			image := &Image{
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
			image.SetTime(tt.args.timing)
			if !reflect.DeepEqual(image.timing, tt.want) {
				t.Errorf("Image.SetTime() = %v, want %v", image.timing, tt.want)
			}
		})
	}
}

func TestImage_SetNameAndDesc(t *testing.T) {
	validName, _ := common.NewName("valid-name")
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
	type args struct {
		name        common.Name
		description string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantName common.Name
		wantDesc string
	}{
		{
			name: "should succeed",
			fields: fields{
				name:        validName,
				description: "",
			},
			args: args{
				name:        validName,
				description: "valid-description",
			},
			wantName: validName,
			wantDesc: "valid-description",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			image := &Image{
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
			image.SetNameAndDesc(tt.args.name, tt.args.description)
		})
	}
}

func TestImage_WithCancel(t *testing.T) {
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
	}{
		{
			name: "should succeed",
			fields: fields{
				ctx:    context.Background(),
				cancel: context.CancelFunc(func() {}),
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			image := &Image{
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
			image.WithCancel()
		})
	}
}

func TestImage_WithTimeout(t *testing.T) {
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
	type args struct {
		timeout time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "should succeed",
			fields: fields{
				ctx:    context.Background(),
				cancel: context.CancelFunc(func() {}),
			},
			args: args{
				timeout: time.Second,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			image := &Image{
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
			image.WithTimeout(tt.args.timeout)
		})
	}
}

func TestImage_Done(t *testing.T) {
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
		want   <-chan struct{}
	}{
		{
			name: "should succeed",
			fields: fields{
				ctx:    context.Background(),
				cancel: context.CancelFunc(func() {}),
			},
			want: nil,
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
			if got := image.Done(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Image.Done() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImage_MarshalJSON(t *testing.T) {
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
		name    string
		fields  fields
		want    []byte
		wantErr bool
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
				tags:       common.Tags{},
			},
			want:    []byte(`{"uuid":"valid-uuid","name":"valid-name","description":"valid-description","status":"success","version":1,"distribution":"rhel8","user":{"username":"valid-username","ssh_key":"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDFjRxF1E73z1K9AjltDkuJyGUW3YluTEAW6PvHEZH6vnzNHI+cut716lGGRFHlYk1Fk51Q/92ZlynJ/HqByaK/MJppkQSL4x3KEm6s5ciwXbVEb3ct4waTgqxPD9gy7NN0uzbrhQMillb50yZgox6d9A/JmyRA1Dlai/esrlKfZ4wtSUl+CMsPoVxC6pIsh1YqUWE7S/dvXsQ8V+O7H0sdXAkZMg09kLUOQe3fliTMg6wppW+tb30g4MWAbHSrXksL1TpYjmP0M+stNetO2EIZ07bc8KpQhZybdM8LUhhPGuZXuKzIlwbkDI7C1yLv574wOYCjG/zk7Zu9qO7p6u8x valid@sshkey"},"packages":null,"repos":[{"name":"test","url":"http://test.com"},{"name":"test2","url":"http://test2.com"}],"installer":{},"tags":[],"created_at":"` + now.Format(time.RFC3339Nano) + `","updated_at":"` + now.Format(time.RFC3339Nano) + `","deleted_at":"` + now.Format(time.RFC3339Nano) + `"}`),
			wantErr: false,
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
			got, err := image.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Image.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Image.MarshalJSON() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestImage_UnmarshalJSON(t *testing.T) {
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
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "should succeed",
			fields: fields{},
			args: args{
				data: []byte(`{"uuid":"valid-uuid","name":"valid-name","description":"valid-description","status":"success","version":1,"distribution":"rhel8","user":{"username":"valid-username","ssh_key":"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDFjRxF1E73z1K9AjltDkuJyGUW3YluTEAW6PvHEZH6vnzNHI+cut716lGGRFHlYk1Fk51Q/92ZlynJ/HqByaK/MJppkQSL4x3KEm6s5ciwXbVEb3ct4waTgqxPD9gy7NN0uzbrhQMillb50yZgox6d9A/JmyRA1Dlai/esrlKfZ4wtSUl+CMsPoVxC6pIsh1YqUWE7S/dvXsQ8V+O7H0sdXAkZMg09kLUOQe3fliTMg6wppW+tb30g4MWAbHSrXksL1TpYjmP0M+stNetO2EIZ07bc8KpQhZybdM8LUhhPGuZXuKzIlwbkDI7C1yLv574wOYCjG/zk7Zu9qO7p6u8x valid@sshkey"},"packages":null,"repos":[{"name":"test","url":"http://test.com"},{"name":"test2","url":"http://test2.com"}],"installer":{},"tags":[],"created_at":"` + now.Format(time.RFC3339Nano) + `","updated_at":"` + now.Format(time.RFC3339Nano) + `","deleted_at":"` + now.Format(time.RFC3339Nano) + `"}`),
			},
			wantErr: false,
		},
		{
			name:   "should fail",
			fields: fields{},
			args: args{
				data: []byte(``),
			},
			wantErr: true,
		},
		{
			name:   "should fail, invalid created_at",
			fields: fields{},
			args: args{
				data: []byte(`{"uuid":"valid-uuid","name":"valid-name","description":"valid-description","status":"success","version":1,"distribution":"rhel8","user":{"username":"valid-username","ssh_key":"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDFjRxF1E73z1K9AjltDkuJyGUW3YluTEAW6PvHEZH6vnzNHI+cut716lGGRFHlYk1Fk51Q/92ZlynJ/HqByaK/MJppkQSL4x3KEm6s5ciwXbVEb3ct4waTgqxPD9gy7NN0uzbrhQMillb50yZgox6d9A/JmyRA1Dlai/esrlKfZ4wtSUl+CMsPoVxC6pIsh1YqUWE7S/dvXsQ8V+O7H0sdXAkZMg09kLUOQe3fliTMg6wppW+tb30g4MWAbHSrXksL1TpYjmP0M+stNetO2EIZ07bc8KpQhZybdM8LUhhPGuZXuKzIlwbkDI7C1yLv574wOYCjG/zk7Zu9qO7p6u8x valid@sshkey"},"packages":null,"repos":[{"name":"test","url":"http://test.com"},{"name":"test2","url":"http://test2.com"}],"installer":{},"tags":[],"created_at":"","updated_at":"` + now.Format(time.RFC3339Nano) + `","deleted_at":"` + now.Format(time.RFC3339Nano) + `"}`),
			},
			wantErr: true,
		},
		{
			name:   "should fail, invalid updated_at",
			fields: fields{},
			args: args{
				data: []byte(`{"uuid":"valid-uuid","name":"valid-name","description":"valid-description","status":"success","version":1,"distribution":"rhel8","user":{"username":"valid-username","ssh_key":"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDFjRxF1E73z1K9AjltDkuJyGUW3YluTEAW6PvHEZH6vnzNHI+cut716lGGRFHlYk1Fk51Q/92ZlynJ/HqByaK/MJppkQSL4x3KEm6s5ciwXbVEb3ct4waTgqxPD9gy7NN0uzbrhQMillb50yZgox6d9A/JmyRA1Dlai/esrlKfZ4wtSUl+CMsPoVxC6pIsh1YqUWE7S/dvXsQ8V+O7H0sdXAkZMg09kLUOQe3fliTMg6wppW+tb30g4MWAbHSrXksL1TpYjmP0M+stNetO2EIZ07bc8KpQhZybdM8LUhhPGuZXuKzIlwbkDI7C1yLv574wOYCjG/zk7Zu9qO7p6u8x valid@sshkey"},"packages":null,"repos":[{"name":"test","url":"http://test.com"},{"name":"test2","url":"http://test2.com"}],"installer":{},"tags":[],"created_at":"` + now.Format(time.RFC3339Nano) + `","updated_at":"","deleted_at":"` + now.Format(time.RFC3339Nano) + `"}`),
			},
			wantErr: true,
		},
		{
			name:   "should fail, invalid deleted_at",
			fields: fields{},
			args: args{
				data: []byte(`{"uuid":"valid-uuid","name":"valid-name","description":"valid-description","status":"success","version":1,"distribution":"rhel8","user":{"username":"valid-username","ssh_key":"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDFjRxF1E73z1K9AjltDkuJyGUW3YluTEAW6PvHEZH6vnzNHI+cut716lGGRFHlYk1Fk51Q/92ZlynJ/HqByaK/MJppkQSL4x3KEm6s5ciwXbVEb3ct4waTgqxPD9gy7NN0uzbrhQMillb50yZgox6d9A/JmyRA1Dlai/esrlKfZ4wtSUl+CMsPoVxC6pIsh1YqUWE7S/dvXsQ8V+O7H0sdXAkZMg09kLUOQe3fliTMg6wppW+tb30g4MWAbHSrXksL1TpYjmP0M+stNetO2EIZ07bc8KpQhZybdM8LUhhPGuZXuKzIlwbkDI7C1yLv574wOYCjG/zk7Zu9qO7p6u8x valid@sshkey"},"packages":null,"repos":[{"name":"test","url":"http://test.com"},{"name":"test2","url":"http://test2.com"}],"installer":{},"tags":[],"created_at":"` + now.Format(time.RFC3339Nano) + `","updated_at":"` + now.Format(time.RFC3339Nano) + `","deleted_at":""}`),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			image := &Image{
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
			if err := image.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Image.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			fmt.Println(image)
		})
	}
}

func TestUnmarshalImageFromDatabase(t *testing.T) {
	validSSHKey := "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDFjRxF1E73z1K9AjltDkuJyGUW3YluTEAW6PvHEZH6vnzNHI+cut716lGGRFHlYk1Fk51Q/92ZlynJ/HqByaK/MJppkQSL4x3KEm6s5ciwXbVEb3ct4waTgqxPD9gy7NN0uzbrhQMillb50yZgox6d9A/JmyRA1Dlai/esrlKfZ4wtSUl+CMsPoVxC6pIsh1YqUWE7S/dvXsQ8V+O7H0sdXAkZMg09kLUOQe3fliTMg6wppW+tb30g4MWAbHSrXksL1TpYjmP0M+stNetO2EIZ07bc8KpQhZybdM8LUhhPGuZXuKzIlwbkDI7C1yLv574wOYCjG/zk7Zu9qO7p6u8x valid@sshkey"
	validName, _ := common.NewName("valid-name")
	validVersion, _ := NewVersion(1)
	validRepos := NewRepos(
		NewRepo("test", "http://test.com"),
		NewRepo("test2", "http://test2.com"),
	)
	now := time.Now()
	type args struct {
		ctx          context.Context
		uuid         string
		name         string
		description  string
		distribution string
		status       string
		username     string
		sshKey       string
		outputType   []string
		tags         []string
		packages     []string
		version      uint
		repos        []interface{}
		createdAt    time.Time
		updatedAt    time.Time
		deletedAt    time.Time
	}
	tests := []struct {
		name      string
		args      args
		wantImage Image
		wantErr   bool
	}{
		{
			name: "should succeed, valid data",
			args: args{
				ctx:          context.Background(),
				uuid:         "valid-uuid",
				name:         "valid-name",
				description:  "valid-description",
				distribution: "rhel8",
				status:       "success",
				username:     "valid-username",
				sshKey:       validSSHKey,
				outputType:   []string{"tar"},
				tags:         []string{"tag1", "tag2"},
				packages:     []string{"package1", "package2"},
				version:      1,
				repos: []interface{}{
					map[string]interface{}{
						"name": "test",
						"url":  "http://test.com",
					},
					map[string]interface{}{
						"name": "test2",
						"url":  "http://test2.com",
					},
				},
				createdAt: now,
				updatedAt: now,
				deletedAt: now,
			},
			wantImage: Image{
				ctx:         context.Background(),
				uuid:        "valid-uuid",
				name:        validName,
				description: "valid-description",
				status:      Success,
				user: User{
					username: "valid-username",
					sshKey:   validSSHKey,
				},
				outputType:   []OutputType{TAR},
				tags:         common.NewTags([]string{"tag1", "tag2"}...),
				packages:     NewPackages([]string{"package1", "package2"}...),
				version:      validVersion,
				distribution: NewDistribution("rhel8"),
				timing:       common.NewTime(now, now, now),
				repos:        validRepos,
				installer:    Installer{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotImage, err := UnmarshalImageFromDatabase(tt.args.ctx, tt.args.uuid, tt.args.name,
				tt.args.description, tt.args.distribution, tt.args.status, tt.args.username, tt.args.sshKey, tt.args.outputType, tt.args.tags, tt.args.packages, tt.args.version, tt.args.repos, tt.args.createdAt, tt.args.updatedAt, tt.args.deletedAt)
			gotImage.ctx = nil        // ignore context
			gotImage.cancel = nil     // ignore context
			tt.wantImage.ctx = nil    // ignore context
			tt.wantImage.cancel = nil // ignore context
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalImageFromDatabase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotImage, tt.wantImage) {
				t.Errorf("UnmarshalImageFromDatabase() = %v, want %v", gotImage, tt.wantImage)
			}
		})
	}
}
