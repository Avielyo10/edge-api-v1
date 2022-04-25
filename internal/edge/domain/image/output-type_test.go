package image

import (
	"reflect"
	"testing"
)

func TestNewOutputTypeFromString(t *testing.T) {
	type args struct {
		out string
	}
	tests := []struct {
		name    string
		args    args
		want    OutputType
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				out: "rhel-edge-installer",
			},
			want: OutputType{
				out: "rhel-edge-installer",
			},
			wantErr: false,
		},
		{
			name: "invalid",
			args: args{
				out: "invalid",
			},
			want:    OutputType{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := NewOutputTypeFromString(tt.args.out)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewOutputTypeFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOutputTypeFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewOutputType(t *testing.T) {
	type args struct {
		out []string
	}
	tests := []struct {
		name string
		args args
		want []OutputType
	}{
		{
			name: "valid",
			args: args{
				out: []string{"rhel-edge-commit"},
			},
			want: []OutputType{
				{
					out: "rhel-edge-commit",
				},
			},
		},
		{
			name: "invalid",
			args: args{
				out: []string{"invalid"},
			},
			want: []OutputType{{"rhel-edge-commit"}},
		},
		{
			name: "nothing, expacting rhel-edge-commit",
			args: args{
				out: []string{},
			},
			want: []OutputType{{"rhel-edge-commit"}},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := NewOutputType(tt.args.out...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOutputType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOutputType_IsZero(t *testing.T) {
	type fields struct {
		out string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "empty",
			want: true,
		},
		{
			name: "not empty",
			fields: fields{
				out: "rhel-edge-commit",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			o := OutputType{
				out: tt.fields.out,
			}
			if got := o.IsZero(); got != tt.want {
				t.Errorf("OutputType.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOutputType_String(t *testing.T) {
	type fields struct {
		out string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "empty",
			want: "",
		},
		{
			name: "not empty",
			fields: fields{
				out: "rhel-edge-commit",
			},
			want: "rhel-edge-commit",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			o := OutputType{
				out: tt.fields.out,
			}
			if got := o.String(); got != tt.want {
				t.Errorf("OutputType.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOutputType_IsISO(t *testing.T) {
	type fields struct {
		out string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "empty",
			want: false,
		},
		{
			name: "not empty",
			fields: fields{
				out: "rhel-edge-commit",
			},
			want: false,
		},
		{
			name: "rhel-edge-installer",
			fields: fields{
				out: "rhel-edge-installer",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			o := OutputType{
				out: tt.fields.out,
			}
			if got := o.IsISO(); got != tt.want {
				t.Errorf("OutputType.IsISO() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOutputType_IsTAR(t *testing.T) {
	type fields struct {
		out string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "empty",
			want: false,
		},
		{
			name: "not empty",
			fields: fields{
				out: "rhel-edge-installer",
			},
			want: false,
		},
		{
			name: "rhel-edge-commit",
			fields: fields{
				out: "rhel-edge-commit",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			o := OutputType{
				out: tt.fields.out,
			}
			if got := o.IsTAR(); got != tt.want {
				t.Errorf("OutputType.IsTAR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAvailableOutputType(t *testing.T) {
	type args struct {
		out string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			want: false,
		},
		{
			name: "not empty",
			args: args{
				out: "rhel-edge-commit",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := IsAvailableOutputType(tt.args.out); got != tt.want {
				t.Errorf("IsAvailableOutputType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOutputType_MarshalJSON(t *testing.T) {
	type fields struct {
		out string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "empty",
			want: []byte(`""`),
		},
		{
			name: "not empty",
			fields: fields{
				out: "rhel-edge-commit",
			},
			want: []byte(`"rhel-edge-commit"`),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			o := OutputType{
				out: tt.fields.out,
			}
			got, err := o.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("OutputType.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OutputType.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOutputType_UnmarshalJSON(t *testing.T) {
	type fields struct {
		out string
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
			name: "empty",
			fields: fields{
				out: "",
			},
			args: args{
				data: []byte(`""`),
			},
			wantErr: false,
		},
		{
			name: "not empty",
			fields: fields{
				out: "rhel-edge-commit",
			},
			args: args{
				data: []byte(`"rhel-edge-commit"`),
			},
			wantErr: false,
		},
		{
			name: "invalid",
			fields: fields{
				out: "rhel-edge-commit",
			},
			args: args{
				data: []byte(`"`),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			o := &OutputType{
				out: tt.fields.out,
			}
			if err := o.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("OutputType.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
