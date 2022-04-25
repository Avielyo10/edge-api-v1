package image

import (
	"context"
	"reflect"
	"testing"

	"github.com/Avielyo10/edge-api/internal/edge/domain/common"
)

func TestImage_Upgrade(t *testing.T) {
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
		wantErr bool
		want    Image
	}{
		{
			name: "Test Upgrade",
			fields: fields{
				ctx:          context.Background(),
				cancel:       func() {},
				uuid:         "",
				name:         common.Name{},
				description:  "",
				timing:       common.Time{},
				status:       Status{"success"},
				version:      Version{1},
				distribution: Distribution{},
				user:         User{},
				packages:     Packages{},
				repos:        Repos{},
				installer:    Installer{},
				outputType:   []OutputType{},
				tags:         common.Tags{},
			},
			wantErr: false,
			want: Image{
				ctx:          context.Background(),
				cancel:       func() {},
				uuid:         "",
				name:         common.Name{},
				description:  "",
				timing:       common.Time{},
				status:       Status{"building"},
				version:      Version{2},
				distribution: Distribution{},
				user:         User{},
				packages:     Packages{},
				repos:        Repos{},
				installer:    Installer{},
				outputType:   []OutputType{},
				tags:         common.Tags{},
			},
		},
		{
			name: "Test Upgrade with error",
			fields: fields{
				ctx:          context.Background(),
				cancel:       func() {},
				uuid:         "",
				name:         common.Name{},
				description:  "",
				timing:       common.Time{},
				status:       Status{"building"},
				version:      Version{},
				distribution: Distribution{},
				user:         User{},
				packages:     Packages{},
				repos:        Repos{},
				installer:    Installer{},
				outputType:   []OutputType{},
				tags:         common.Tags{},
			},
			wantErr: true,
			want: Image{
				ctx:          context.Background(),
				cancel:       func() {},
				uuid:         "",
				name:         common.Name{},
				description:  "",
				timing:       common.Time{},
				status:       Status{"building"},
				version:      Version{2},
				distribution: Distribution{},
				user:         User{},
				packages:     Packages{},
				repos:        Repos{},
				installer:    Installer{},
				outputType:   []OutputType{},
				tags:         common.Tags{},
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
			if err := image.Upgrade(); (err != nil) != tt.wantErr {
				t.Errorf("Image.Upgrade() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if !reflect.DeepEqual(image.version, tt.want.version) {
					t.Errorf("Image.Upgrade() version = %v, want %v", image.version, tt.want.version)
				}
				if !reflect.DeepEqual(image.status, tt.want.status) {
					t.Errorf("Image.Upgrade() status = %v, want %v", image.status, tt.want.status)
				}
			}
		})
	}
}

func TestImage_Rollback(t *testing.T) {
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
		wantErr bool
		want    Image
	}{
		{
			name: "Test Rollback",
			fields: fields{
				ctx:          context.Background(),
				cancel:       func() {},
				uuid:         "",
				name:         common.Name{},
				description:  "",
				timing:       common.Time{},
				status:       Status{"building"},
				version:      Version{2},
				distribution: Distribution{},
				user:         User{},
				packages:     Packages{},
				repos:        Repos{},
				installer:    Installer{},
				outputType:   []OutputType{},
				tags:         common.Tags{},
			},
			wantErr: false,
			want: Image{
				ctx:          context.Background(),
				cancel:       func() {},
				uuid:         "",
				name:         common.Name{},
				description:  "",
				timing:       common.Time{},
				status:       Status{"success"},
				version:      Version{1},
				distribution: Distribution{},
				user:         User{},
				packages:     Packages{},
				repos:        Repos{},
				installer:    Installer{},
				outputType:   []OutputType{},
				tags:         common.Tags{},
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
			if err := image.Rollback(); (err != nil) != tt.wantErr {
				t.Errorf("Image.Rollback() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if !reflect.DeepEqual(image.version, tt.want.version) {
					t.Errorf("Image.Rollback() version = %v, want %v", image.version, tt.want.version)
				}
				if !reflect.DeepEqual(image.status, tt.want.status) {
					t.Errorf("Image.Rollback() status = %v, want %v", image.status, tt.want.status)
				}
			}
		})
	}
}

func TestImage_CheckForUpdate(t *testing.T) {
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
		wantErr bool
	}{
		{
			name: "Test CheckForUpdate",
			fields: fields{
				ctx:          context.Background(),
				cancel:       func() {},
				uuid:         "",
				name:         common.Name{},
				description:  "",
				timing:       common.Time{},
				status:       Status{"success"},
				version:      Version{1},
				distribution: Distribution{},
				user:         User{},
				packages:     Packages{},
				repos:        Repos{},
				installer:    Installer{},
				outputType:   []OutputType{},
				tags:         common.Tags{},
			},
			wantErr: false,
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
			if err := image.CheckForUpdate(); (err != nil) != tt.wantErr {
				t.Errorf("Image.CheckForUpdate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestImage_IsSuccessful(t *testing.T) {
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
			name: "Test IsSuccessful",
			fields: fields{
				ctx:          context.Background(),
				cancel:       func() {},
				uuid:         "",
				name:         common.Name{},
				description:  "",
				timing:       common.Time{},
				status:       Status{"success"},
				version:      Version{1},
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
		{
			name: "Test IsSuccessful",
			fields: fields{
				ctx:          context.Background(),
				cancel:       func() {},
				uuid:         "",
				name:         common.Name{},
				description:  "",
				timing:       common.Time{},
				status:       Status{"failure"},
				version:      Version{1},
				distribution: Distribution{},
				user:         User{},
				packages:     Packages{},
				repos:        Repos{},
				installer:    Installer{},
				outputType:   []OutputType{},
				tags:         common.Tags{},
			},
			want: false,
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
			if got := image.IsSuccessful(); got != tt.want {
				t.Errorf("Image.IsSuccessful() = %v, want %v", got, tt.want)
			}
		})
	}
}
