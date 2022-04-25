package image

import (
	"reflect"
	"testing"
)

func TestVersion_IsZero(t *testing.T) {
	type fields struct {
		number uint
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "zero",
			fields: fields{
				number: 0,
			},
			want: true,
		},
		{
			name: "non-zero",
			fields: fields{
				number: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			v := Version{
				number: tt.fields.number,
			}
			if got := v.IsZero(); got != tt.want {
				t.Errorf("Version.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVersion_Update(t *testing.T) {
	type fields struct {
		number uint
	}
	tests := []struct {
		name   string
		fields fields
		want   Version
	}{
		{
			name: "update",
			fields: fields{
				number: 1,
			},
			want: Version{
				number: 2,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			v := &Version{
				number: tt.fields.number,
			}
			v.Update()
			if !reflect.DeepEqual(v, &tt.want) {
				t.Errorf("Version.Update() = %v, want %v", v, &tt.want)
			}
		})
	}
}

func TestVersion_Rollback(t *testing.T) {
	type fields struct {
		number uint
	}
	tests := []struct {
		name   string
		fields fields
		want   Version
	}{
		{
			name: "rollback",
			fields: fields{
				number: 2,
			},
			want: Version{
				number: 1,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			v := &Version{
				number: tt.fields.number,
			}
			v.Rollback()
			if !reflect.DeepEqual(v, &tt.want) {
				t.Errorf("Version.Rollback() = %v, want %v", v, &tt.want)
			}
		})
	}
}

func TestNewVersion(t *testing.T) {
	type args struct {
		number uint
	}
	tests := []struct {
		name    string
		args    args
		want    Version
		wantErr bool
	}{
		{
			name: "invalid",
			args: args{
				number: 0,
			},
			wantErr: true,
		},
		{
			name: "valid",
			args: args{
				number: 1,
			},
			want: Version{
				number: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := NewVersion(tt.args.number)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVersion_Uint(t *testing.T) {
	type fields struct {
		number uint
	}
	tests := []struct {
		name   string
		fields fields
		want   uint
	}{
		{
			name: "valid",
			fields: fields{
				number: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			v := Version{
				number: tt.fields.number,
			}
			if got := v.Uint(); got != tt.want {
				t.Errorf("Version.Uint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVersion_MarshalJSON(t *testing.T) {
	type fields struct {
		number uint
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
				number: 1,
			},
			want:    []byte(`1`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			v := Version{
				number: tt.fields.number,
			}
			got, err := v.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Version.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Version.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVersion_UnmarshalJSON(t *testing.T) {
	type fields struct {
		number uint
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
			name: "valid",
			fields: fields{
				number: 1,
			},
			args: args{
				data: []byte(`1`),
			},
			wantErr: false,
		},
		{
			name: "invalid",
			fields: fields{
				number: 1,
			},
			args: args{
				data: []byte(`0`),
			},
			wantErr: true,
		},
		{
			name: "invalid string",
			fields: fields{
				number: 1,
			},
			args: args{
				data: []byte(`test`),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			v := &Version{
				number: tt.fields.number,
			}
			if err := v.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Version.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
