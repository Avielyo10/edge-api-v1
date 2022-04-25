package common

import (
	"reflect"
	"testing"
)

func TestNewName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    Name
		wantErr bool
	}{
		{
			name: "valid name",
			args: args{
				name: "1234567",
			},
			want: Name{
				name: "1234567",
			},
			wantErr: false,
		},
		{
			name: "invalid name",
			args: args{
				name: "",
			},
			want:    Name{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := NewName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestName_IsZero(t *testing.T) {
	tests := []struct {
		name string
		n    Name
		want bool
	}{
		{
			name: "zero name",
			n:    Name{},
			want: true,
		},
		{
			name: "non-zero name",
			n:    Name{name: "1234567"},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.n.IsZero(); got != tt.want {
				t.Errorf("Name.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid name",
			args: args{
				name: "1234567",
			},
			want: true,
		},
		{
			name: "invalid name",
			args: args{
				name: "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := isValidName(tt.args.name); got != tt.want {
				t.Errorf("isValidName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestName_String(t *testing.T) {
	tests := []struct {
		name string
		n    Name
		want string
	}{
		{
			name: "valid name",
			n:    Name{name: "1234567"},
			want: "1234567",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.n.String(); got != tt.want {
				t.Errorf("Name.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestName_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		n       Name
		want    []byte
		wantErr bool
	}{
		{
			name: "valid name",
			n:    Name{name: "1234567"},
			want: []byte(`"1234567"`),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := tt.n.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Name.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Name.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestName_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		n       *Name
		args    args
		wantErr bool
	}{
		{
			name: "valid name",
			n:    &Name{},
			args: args{
				data: []byte(`"1234567"`),
			},
			wantErr: false,
		},
		{
			name: "invalid name",
			n:    &Name{},
			args: args{
				data: []byte(``),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := tt.n.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Name.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
