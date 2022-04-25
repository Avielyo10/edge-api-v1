package image

import (
	"reflect"
	"testing"
)

func TestNewDistribution(t *testing.T) {
	type args struct {
		dist string
	}
	tests := []struct {
		name string
		args args
		want Distribution
	}{
		{
			name: "valid",
			args: args{
				dist: "test",
			},
			want: Distribution{
				dist: "test",
			},
		},
		{
			name: "invalid",
			args: args{
				dist: "",
			},
			want: Distribution{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := NewDistribution(tt.args.dist); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDistribution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistribution_String(t *testing.T) {
	type fields struct {
		dist string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "valid",
			fields: fields{
				dist: "test",
			},
			want: "test",
		},
		{
			name: "invalid",
			fields: fields{
				dist: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			d := Distribution{
				dist: tt.fields.dist,
			}
			if got := d.String(); got != tt.want {
				t.Errorf("Distribution.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistribution_IsZero(t *testing.T) {
	type fields struct {
		dist string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "valid",
			fields: fields{
				dist: "test",
			},
			want: false,
		},
		{
			name: "invalid",
			fields: fields{
				dist: "",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			d := Distribution{
				dist: tt.fields.dist,
			}
			if got := d.IsZero(); got != tt.want {
				t.Errorf("Distribution.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidDist(t *testing.T) {
	type args struct {
		dist string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid",
			args: args{
				dist: "test",
			},
			want: true,
		},
		{
			name: "invalid",
			args: args{
				dist: "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := isValidDist(tt.args.dist); got != tt.want {
				t.Errorf("isValidDist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistribution_MarshalJSON(t *testing.T) {
	type fields struct {
		dist string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "valid",
			fields: fields{
				dist: "test",
			},
			want:    []byte(`"test"`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			d := Distribution{
				dist: tt.fields.dist,
			}
			got, err := d.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Distribution.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Distribution.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistribution_UnmarshalJSON(t *testing.T) {
	type fields struct {
		dist string
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
			name: "valid",
			fields: fields{
				dist: "test",
			},
			args: args{
				b: []byte(`"test"`),
			},
			wantErr: false,
		},
		{
			name: "invalid",
			fields: fields{
				dist: "",
			},
			args: args{
				b: []byte(`""`),
			},
			wantErr: true,
		},
		{
			name: "invalid string",
			fields: fields{
				dist: "",
			},
			args: args{
				b: []byte(`"test`),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			d := &Distribution{
				dist: tt.fields.dist,
			}
			if err := d.UnmarshalJSON(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("Distribution.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
