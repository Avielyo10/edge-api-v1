package image

import (
	"context"
	"reflect"
	"sort"
	"testing"

	"github.com/Avielyo10/edge-api/internal/common/models"
	"github.com/Avielyo10/edge-api/internal/edge/domain/common"
)

func TestNewPackages(t *testing.T) {
	type args struct {
		packages []string
	}
	tests := []struct {
		name string
		args args
		want Packages
	}{
		{
			name: "test",
			args: args{
				packages: []string{"test"},
			},
			want: Packages{
				requiredPackages: []Package{
					{"ansible"},
					{"rhc"},
					{"rhc-worker-playbook"},
					{"subscription-manager"},
					{"subscription-manager-plugin-ostree"},
					{"insights-client"},
				},
				packages: []Package{{name: "test"}},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := NewPackages(tt.args.packages...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPackages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPackage(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want Package
	}{
		{
			name: "test",
			args: args{
				name: "test",
			},
			want: Package{
				name: "test",
			},
		},
		{
			name: "empty",
			args: args{
				name: "",
			},
			want: Package{},
		},
		{
			name: "multiple spaces",
			args: args{
				name: "  ",
			},
			want: Package{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := NewPackage(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPackage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidPackage(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid",
			args: args{
				name: "test",
			},
			want: true,
		},
		{
			name: "empty",
			args: args{
				name: "",
			},
			want: false,
		},
		{
			name: "multiple spaces",
			args: args{
				name: "  ",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := isValidPackage(tt.args.name); got != tt.want {
				t.Errorf("isValidPackage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPackage_IsZero(t *testing.T) {
	type fields struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "empty",
			fields: fields{
				name: "",
			},
			want: true,
		},
		{
			name: "test",
			fields: fields{
				name: "test",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			p := Package{
				name: tt.fields.name,
			}
			if got := p.IsZero(); got != tt.want {
				t.Errorf("Package.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPackages_Packages(t *testing.T) {
	type fields struct {
		requiredPackages []Package
		packages         []Package
	}
	tests := []struct {
		name   string
		fields fields
		want   []Package
	}{
		{
			name: "test",
			fields: fields{
				requiredPackages: []Package{
					{"ansible"},
					{"rhc"},
					{"rhc-worker-playbook"},
					{"subscription-manager"},
					{"subscription-manager-plugin-ostree"},
					{"insights-client"},
				},
				packages: []Package{
					{"test"},
				},
			},
			want: []Package{
				{"ansible"},
				{"rhc"},
				{"rhc-worker-playbook"},
				{"subscription-manager"},
				{"subscription-manager-plugin-ostree"},
				{"insights-client"},
				{"test"},
			},
		},
		{
			name: "empty",
			fields: fields{
				requiredPackages: []Package{
					{"ansible"},
					{"rhc"},
					{"rhc-worker-playbook"},
					{"subscription-manager"},
					{"subscription-manager-plugin-ostree"},
					{"insights-client"},
				},
				packages: []Package{},
			},
			want: []Package{
				{"ansible"},
				{"rhc"},
				{"rhc-worker-playbook"},
				{"subscription-manager"},
				{"subscription-manager-plugin-ostree"},
				{"insights-client"},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			p := Packages{
				requiredPackages: tt.fields.requiredPackages,
				packages:         tt.fields.packages,
			}
			if got := p.Packages(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Packages.Packages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPackages_Add(t *testing.T) {
	type fields struct {
		requiredPackages []Package
		packages         []Package
	}
	type args struct {
		pkg []Package
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []Package
	}{
		{
			name: "test",
			fields: fields{
				requiredPackages: []Package{
					{"ansible"},
					{"rhc"},
					{"rhc-worker-playbook"},
					{"subscription-manager"},
					{"subscription-manager-plugin-ostree"},
					{"insights-client"},
				},
				packages: []Package{
					{"test"},
				},
			},
			args: args{
				pkg: []Package{
					{"test2"},
				},
			},
			want: []Package{
				{"ansible"},
				{"rhc"},
				{"rhc-worker-playbook"},
				{"subscription-manager"},
				{"subscription-manager-plugin-ostree"},
				{"insights-client"},
				{"test"},
				{"test2"},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			p := &Packages{
				requiredPackages: tt.fields.requiredPackages,
				packages:         tt.fields.packages,
			}
			p.Add(tt.args.pkg...)
			if !reflect.DeepEqual(p.Packages(), tt.want) {
				t.Errorf("Packages.Add() = %v, want %v", p.packages, tt.want)
			}
		})
	}
}

func TestPackages_Remove(t *testing.T) {
	type fields struct {
		requiredPackages []Package
		packages         []Package
	}
	type args struct {
		pkgs []Package
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []Package
	}{
		{
			name: "test",
			fields: fields{
				requiredPackages: []Package{
					{"ansible"},
					{"rhc"},
					{"rhc-worker-playbook"},
					{"subscription-manager"},
					{"subscription-manager-plugin-ostree"},
					{"insights-client"},
				},
				packages: []Package{
					{"test"},
					{"test2"},
				},
			},
			args: args{
				pkgs: []Package{
					{"test"},
				},
			},
			want: []Package{
				{"ansible"},
				{"rhc"},
				{"rhc-worker-playbook"},
				{"subscription-manager"},
				{"subscription-manager-plugin-ostree"},
				{"insights-client"},
				{"test2"},
			},
		},
		{
			name: "try remove non-existing",
			fields: fields{
				requiredPackages: []Package{
					{"ansible"},
					{"rhc"},
					{"rhc-worker-playbook"},
					{"subscription-manager"},
					{"subscription-manager-plugin-ostree"},
					{"insights-client"},
				},
				packages: []Package{
					{"test"},
					{"test2"},
				},
			},
			args: args{
				pkgs: []Package{
					{"test3"},
				},
			},
			want: []Package{
				{"ansible"},
				{"rhc"},
				{"rhc-worker-playbook"},
				{"subscription-manager"},
				{"subscription-manager-plugin-ostree"},
				{"insights-client"},
				{"test"},
				{"test2"},
			},
		},
		{
			name: "try remove required",
			fields: fields{
				requiredPackages: []Package{
					{"ansible"},
					{"rhc"},
					{"rhc-worker-playbook"},
					{"subscription-manager"},
					{"subscription-manager-plugin-ostree"},
					{"insights-client"},
				},
				packages: []Package{
					{"test"},
					{"test2"},
				},
			},
			args: args{
				pkgs: []Package{
					{"ansible"},
					{"test2"},
				},
			},
			want: []Package{
				{"ansible"},
				{"rhc"},
				{"rhc-worker-playbook"},
				{"subscription-manager"},
				{"subscription-manager-plugin-ostree"},
				{"insights-client"},
				{"test"},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			p := &Packages{
				requiredPackages: tt.fields.requiredPackages,
				packages:         tt.fields.packages,
			}
			p.Remove(tt.args.pkgs...)
			if !package_array_sorted_equal(p.Packages(), tt.want) {
				t.Errorf("Packages.Remove() = %v, want %v", p.Packages(), tt.want)
			}
		})
	}
}

func TestImage_AddPackage(t *testing.T) {
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
		pkg []Package
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Packages
	}{
		{
			name: "test",
			fields: fields{
				packages: Packages{
					requiredPackages: []Package{
						{"ansible"},
						{"rhc"},
						{"rhc-worker-playbook"},
						{"subscription-manager"},
						{"subscription-manager-plugin-ostree"},
						{"insights-client"},
					},
					packages: []Package{
						{"test"},
					},
				},
			},
			args: args{
				pkg: []Package{
					{"test2"},
				},
			},
			want: Packages{
				requiredPackages: []Package{
					{"ansible"},
					{"rhc"},
					{"rhc-worker-playbook"},
					{"subscription-manager"},
					{"subscription-manager-plugin-ostree"},
					{"insights-client"},
				},
				packages: []Package{
					{"test"},
					{"test2"},
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			i := &Image{
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
			i.AddPackage(tt.args.pkg...)
			if !reflect.DeepEqual(i.Packages(), tt.want) {
				t.Errorf("Image.AddPackage() = %v, want %v", i.packages, tt.want)
			}
		})
	}
}

func TestImage_RemovePackage(t *testing.T) {
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
		pkg []Package
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Packages
	}{
		{
			name: "test",
			fields: fields{
				packages: Packages{
					requiredPackages: []Package{
						{"ansible"},
						{"rhc"},
						{"rhc-worker-playbook"},
						{"subscription-manager"},
						{"subscription-manager-plugin-ostree"},
						{"insights-client"},
					},
					packages: []Package{
						{"test"},
					},
				},
			},
			args: args{
				pkg: []Package{
					{"test"},
				},
			},
			want: Packages{
				requiredPackages: []Package{
					{"ansible"},
					{"rhc"},
					{"rhc-worker-playbook"},
					{"subscription-manager"},
					{"subscription-manager-plugin-ostree"},
					{"insights-client"},
				},
			},
		},
		{
			name: "remove required package",
			fields: fields{
				packages: Packages{
					requiredPackages: []Package{
						{"ansible"},
						{"rhc"},
						{"rhc-worker-playbook"},
						{"subscription-manager"},
						{"subscription-manager-plugin-ostree"},
						{"insights-client"},
					},
					packages: []Package{
						{"test"},
					},
				},
			},
			args: args{
				pkg: []Package{
					{"ansible"},
				},
			},
			want: Packages{
				requiredPackages: []Package{
					{"ansible"},
					{"rhc"},
					{"rhc-worker-playbook"},
					{"subscription-manager"},
					{"subscription-manager-plugin-ostree"},
					{"insights-client"},
				},
				packages: []Package{
					{"test"},
				},
			},
		},
		{
			name: "remove non existing package",
			fields: fields{
				packages: Packages{
					requiredPackages: []Package{
						{"ansible"},
						{"rhc"},
						{"rhc-worker-playbook"},
						{"subscription-manager"},
						{"subscription-manager-plugin-ostree"},
						{"insights-client"},
					},
					packages: []Package{
						{"test"},
					},
				},
			},
			args: args{
				pkg: []Package{
					{"test2"},
				},
			},
			want: Packages{
				requiredPackages: []Package{
					{"ansible"},
					{"rhc"},
					{"rhc-worker-playbook"},
					{"subscription-manager"},
					{"subscription-manager-plugin-ostree"},
					{"insights-client"},
				},
				packages: []Package{
					{"test"},
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			i := &Image{
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
			i.RemovePackage(tt.args.pkg...)
			if !reflect.DeepEqual(i.Packages().StringArray(), tt.want.StringArray()) {
				t.Errorf("Image.RemovePackage() = %v, want %v", i.packages, tt.want)
			}
		})
	}
}

func TestPackages_StringArray(t *testing.T) {
	type fields struct {
		requiredPackages []Package
		packages         []Package
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "test",
			fields: fields{
				requiredPackages: []Package{
					{"ansible"},
					{"rhc"},
					{"rhc-worker-playbook"},
					{"subscription-manager"},
					{"subscription-manager-plugin-ostree"},
					{"insights-client"},
				},
				packages: []Package{
					{"test"},
				},
			},
			want: []string{"ansible", "rhc", "rhc-worker-playbook", "subscription-manager", "subscription-manager-plugin-ostree", "insights-client", "test"},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			p := Packages{
				requiredPackages: tt.fields.requiredPackages,
				packages:         tt.fields.packages,
			}
			if got := p.StringArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Packages.StringArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPackages_MarshalJSON(t *testing.T) {
	type fields struct {
		requiredPackages []Package
		packages         []Package
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "test",
			fields: fields{
				requiredPackages: []Package{
					{"ansible"},
					{"rhc"},
					{"rhc-worker-playbook"},
					{"subscription-manager"},
					{"subscription-manager-plugin-ostree"},
					{"insights-client"},
				},
				packages: []Package{
					{"test"},
				},
			},
			want:    []byte(`["ansible","rhc","rhc-worker-playbook","subscription-manager","subscription-manager-plugin-ostree","insights-client","test"]`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			p := Packages{
				requiredPackages: tt.fields.requiredPackages,
				packages:         tt.fields.packages,
			}
			got, err := p.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Packages.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Packages.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPackages_UnmarshalJSON(t *testing.T) {
	type fields struct {
		requiredPackages []Package
		packages         []Package
	}
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test",
			fields: fields{
				requiredPackages: []Package{
					{"ansible"},
					{"rhc"},
					{"rhc-worker-playbook"},
					{"subscription-manager"},
					{"subscription-manager-plugin-ostree"},
					{"insights-client"},
				},
				packages: []Package{
					{"test"},
				},
			},
			args: args{
				b: []byte(`["ansible","rhc","rhc-worker-playbook","subscription-manager","subscription-manager-plugin-ostree","insights-client","test"]`),
			},
			wantErr: false,
		},
		{
			name: "empty",
			fields: fields{
				requiredPackages: []Package{},
				packages:         []Package{},
			},
			args: args{
				b: []byte(`[]`),
			},
			wantErr: false,
		},
		{
			name: "invalid",
			fields: fields{
				requiredPackages: []Package{},
				packages:         []Package{},
			},
			args: args{
				b: []byte(`"`),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			p := &Packages{
				requiredPackages: tt.fields.requiredPackages,
				packages:         tt.fields.packages,
			}
			if err := p.UnmarshalJSON(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("Packages.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestErrMissingRequiredPackage_Error(t *testing.T) {
	type fields struct {
		pkg Package
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "test",
			fields: fields{
				pkg: Package{"test"},
			},
			want: "missing required package: test",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			e := ErrMissingRequiredPackage{
				pkg: tt.fields.pkg,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("ErrMissingRequiredPackage.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPackages_Has(t *testing.T) {
	type fields struct {
		requiredPackages []Package
		packages         []Package
	}
	type args struct {
		pkg Package
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "test",
			fields: fields{
				requiredPackages: []Package{
					{"ansible"},
					{"rhc"},
					{"rhc-worker-playbook"},
					{"subscription-manager"},
					{"subscription-manager-plugin-ostree"},
					{"insights-client"},
				},
				packages: []Package{
					{"test"},
				},
			},
			args: args{
				pkg: Package{"test"},
			},
			want: true,
		},
		{
			name: "test, non existent",
			fields: fields{
				requiredPackages: []Package{
					{"ansible"},
					{"rhc"},
					{"rhc-worker-playbook"},
					{"subscription-manager"},
					{"subscription-manager-plugin-ostree"},
					{"insights-client"},
				},
				packages: []Package{
					{"test"},
				},
			},
			args: args{
				pkg: Package{"test2"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			p := Packages{
				requiredPackages: tt.fields.requiredPackages,
				packages:         tt.fields.packages,
			}
			if got := p.Has(tt.args.pkg); got != tt.want {
				t.Errorf("Packages.Has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func package_array_sorted_equal(a, b []Package) bool {
	if len(a) != len(b) {
		return false
	}

	a_copy := make([]string, len(a))
	b_copy := make([]string, len(b))

	for i, p := range a {
		a_copy[i] = p.name
	}

	for i, p := range b {
		b_copy[i] = p.name
	}

	sort.Strings(a_copy)
	sort.Strings(b_copy)

	return reflect.DeepEqual(a_copy, b_copy)
}

func TestPackages_MarshalGorm(t *testing.T) {
	type fields struct {
		requiredPackages []Package
		packages         []Package
	}
	type args struct {
		account string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []models.Package
	}{
		{
			name: "test",
			fields: fields{
				requiredPackages: []Package{
					{"ansible"},
					{"rhc"},
					{"rhc-worker-playbook"},
					{"subscription-manager"},
					{"subscription-manager-plugin-ostree"},
					{"insights-client"},
				},
				packages: []Package{
					{"test"},
				},
			},
			args: args{
				account: "test",
			},
			want: []models.Package{
				{Name: "ansible", Account: "test"},
				{Name: "rhc", Account: "test"},
				{Name: "rhc-worker-playbook", Account: "test"},
				{Name: "subscription-manager", Account: "test"},
				{Name: "subscription-manager-plugin-ostree", Account: "test"},
				{Name: "insights-client", Account: "test"},
				{Name: "test", Account: "test"},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			p := Packages{
				requiredPackages: tt.fields.requiredPackages,
				packages:         tt.fields.packages,
			}
			if got := p.MarshalGorm(tt.args.account); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Packages.MarshalGorm() = %v, want %v", got, tt.want)
			}
		})
	}
}
