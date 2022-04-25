package image

import (
	"reflect"
	"testing"
)

func TestNewStatusFromString(t *testing.T) {
	type args struct {
		state string
	}
	tests := []struct {
		name    string
		args    args
		want    Status
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				state: "success",
			},
			want: Status{
				state: "success",
			},
			wantErr: false,
		},
		{
			name: "invalid",
			args: args{
				state: "invalid",
			},
			want:    Status{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := NewStatusFromString(tt.args.state)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewStatusFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStatusFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewStatus(t *testing.T) {
	tests := []struct {
		name string
		want Status
	}{
		{
			name: "new status is building",
			want: Status{
				state: "building",
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := NewStatus(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatus_IsZero(t *testing.T) {
	type fields struct {
		state string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "zero status",
			fields: fields{
				state: "",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := Status{
				state: tt.fields.state,
			}
			if got := s.IsZero(); got != tt.want {
				t.Errorf("Status.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatus_String(t *testing.T) {
	type fields struct {
		state string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success",
			fields: fields{
				state: "success",
			},
			want: "success",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := Status{
				state: tt.fields.state,
			}
			if got := s.String(); got != tt.want {
				t.Errorf("Status.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatus_IsSuccess(t *testing.T) {
	type fields struct {
		state string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "success",
			fields: fields{
				state: "success",
			},
			want: true,
		},
		{
			name: "failure",
			fields: fields{
				state: "building",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := Status{
				state: tt.fields.state,
			}
			if got := s.IsSuccess(); got != tt.want {
				t.Errorf("Status.IsSuccess() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatus_IsBuilding(t *testing.T) {
	type fields struct {
		state string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "success",
			fields: fields{
				state: "building",
			},
			want: true,
		},
		{
			name: "failure",
			fields: fields{
				state: "success",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := Status{
				state: tt.fields.state,
			}
			if got := s.IsBuilding(); got != tt.want {
				t.Errorf("Status.IsBuilding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatus_IsError(t *testing.T) {
	type fields struct {
		state string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "success",
			fields: fields{
				state: "error",
			},
			want: true,
		},
		{
			name: "failure",
			fields: fields{
				state: "success",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := Status{
				state: tt.fields.state,
			}
			if got := s.IsError(); got != tt.want {
				t.Errorf("Status.IsError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatus_MarshalJSON(t *testing.T) {
	type fields struct {
		state string
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
				state: "success",
			},
			want:    []byte(`"success"`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := Status{
				state: tt.fields.state,
			}
			got, err := s.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Status.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Status.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatus_UnmarshalJSON(t *testing.T) {
	type fields struct {
		state string
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
				state: "success",
			},
			args: args{
				data: []byte(`"success"`),
			},
			wantErr: false,
		},
		{
			name: "failure",
			fields: fields{
				state: "success",
			},
			args: args{
				data: []byte(`"failure"`),
			},
			wantErr: true,
		},
		{
			name: "invalid",
			fields: fields{
				state: "success",
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
			s := &Status{
				state: tt.fields.state,
			}
			if err := s.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Status.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
