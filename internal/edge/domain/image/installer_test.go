package image

import (
	"reflect"
	"testing"

	"github.com/Avielyo10/edge-api/internal/common/models"
	"github.com/bxcodec/faker/v3"
)

func TestNewInstaller(t *testing.T) {
	type args struct {
		isoURL       string
		composeJobID string
		checksum     string
	}
	type testsStruct struct {
		name string
		args args
		want Installer
	}
	tests := make([]testsStruct, 3)
	if err := faker.FakeData(&tests); err != nil {
		t.Fatal(err)
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := NewInstaller(tt.args.isoURL, tt.args.composeJobID, tt.args.checksum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInstaller() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInstaller_ISOURL(t *testing.T) {
	type fields struct {
		isoURL       string
		composeJobID string
		checksum     string
	}
	type testsStruct struct {
		name   string
		fields fields
		want   string
	}
	tests := make([]testsStruct, 3)
	if err := faker.FakeData(&tests); err != nil {
		t.Fatal(err)
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			i := Installer{
				isoURL:       tt.fields.isoURL,
				composeJobID: tt.fields.composeJobID,
				checksum:     tt.fields.checksum,
			}
			if got := i.ISOURL(); got != tt.want {
				t.Errorf("Installer.ISOURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInstaller_ComposeJobID(t *testing.T) {
	type fields struct {
		isoURL       string
		composeJobID string
		checksum     string
	}
	type testsStruct struct {
		name   string
		fields fields
		want   string
	}
	tests := make([]testsStruct, 3)
	if err := faker.FakeData(&tests); err != nil {
		t.Fatal(err)
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			i := Installer{
				isoURL:       tt.fields.isoURL,
				composeJobID: tt.fields.composeJobID,
				checksum:     tt.fields.checksum,
			}
			if got := i.ComposeJobID(); got != tt.want {
				t.Errorf("Installer.ComposeJobID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInstaller_Checksum(t *testing.T) {
	type fields struct {
		isoURL       string
		composeJobID string
		checksum     string
	}
	type testsStruct struct {
		name   string
		fields fields
		want   string
	}
	tests := make([]testsStruct, 3)
	if err := faker.FakeData(&tests); err != nil {
		t.Fatal(err)
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			i := Installer{
				isoURL:       tt.fields.isoURL,
				composeJobID: tt.fields.composeJobID,
				checksum:     tt.fields.checksum,
			}
			if got := i.Checksum(); got != tt.want {
				t.Errorf("Installer.Checksum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInstaller_MarshalGorm(t *testing.T) {
	type fields struct {
		isoURL       string
		composeJobID string
		checksum     string
	}
	tests := []struct {
		name   string
		fields fields
		want   *models.Installer
	}{
		{
			name: "success",
			fields: fields{
				isoURL:       "https://example.com/iso.iso",
				composeJobID: "12345",
				checksum:     "12345",
			},
			want: &models.Installer{
				ISOURL:       "https://example.com/iso.iso",
				ComposeJobID: "12345",
				Checksum:     "12345",
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			i := Installer{
				isoURL:       tt.fields.isoURL,
				composeJobID: tt.fields.composeJobID,
				checksum:     tt.fields.checksum,
			}
			if got := i.MarshalGorm(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Installer.MarshalGorm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInstaller_UnmarshalGorm(t *testing.T) {
	type fields struct {
		isoURL       string
		composeJobID string
		checksum     string
	}
	type args struct {
		in *models.Installer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Installer
	}{
		{
			name: "success",
			fields: fields{
				isoURL:       "https://example.com/iso.iso",
				composeJobID: "12345",
				checksum:     "12345",
			},
			args: args{
				in: &models.Installer{
					ISOURL:       "https://example.com/iso.iso",
					ComposeJobID: "12345",
					Checksum:     "12345",
				},
			},
			want: Installer{
				isoURL:       "https://example.com/iso.iso",
				composeJobID: "12345",
				checksum:     "12345",
			},
		},
		{
			name: "empty",
			fields: fields{
				isoURL:       "",
				composeJobID: "",
				checksum:     "",
			},
			args: args{
				in: nil,
			},
			want: Installer{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			i := Installer{
				isoURL:       tt.fields.isoURL,
				composeJobID: tt.fields.composeJobID,
				checksum:     tt.fields.checksum,
			}
			if got := i.UnmarshalGorm(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Installer.UnmarshalGorm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInstaller_MarshalJSON(t *testing.T) {
	type fields struct {
		isoURL       string
		composeJobID string
		checksum     string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				isoURL:       "https://example.com/iso.iso",
				composeJobID: "12345",
				checksum:     "12345",
			},
			want: []byte(`{"iso_url":"https://example.com/iso.iso","compose_job_id":"12345","checksum":"12345"}`),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			i := Installer{
				isoURL:       tt.fields.isoURL,
				composeJobID: tt.fields.composeJobID,
				checksum:     tt.fields.checksum,
			}
			got, err := i.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Installer.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Installer.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInstaller_UnmarshalJSON(t *testing.T) {
	type fields struct {
		isoURL       string
		composeJobID string
		checksum     string
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
			name: "success",
			fields: fields{
				isoURL:       "https://example.com/iso.iso",
				composeJobID: "12345",
				checksum:     "12345",
			},
			args: args{
				data: []byte(`{"iso_url":"https://example.com/iso.iso","compose_job_id":"12345","checksum":"12345"}`),
			},
		},
		{
			name: "empty",
			fields: fields{
				isoURL:       "",
				composeJobID: "",
				checksum:     "",
			},
			args: args{
				data: []byte(`{}`),
			},
		},
		{
			name: "invalid",
			fields: fields{
				isoURL:       "",
				composeJobID: "",
				checksum:     "",
			},
			args: args{
				data: []byte(`["fake","data"]`),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			i := &Installer{
				isoURL:       tt.fields.isoURL,
				composeJobID: tt.fields.composeJobID,
				checksum:     tt.fields.checksum,
			}
			if err := i.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Installer.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
