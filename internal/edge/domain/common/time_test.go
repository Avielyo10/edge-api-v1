package common

import (
	"reflect"
	"testing"
	"time"
)

func TestNewTime(t *testing.T) {
	now := time.Now()
	type args struct {
		createdAt time.Time
		updatedAt time.Time
		deletedAt time.Time
	}
	tests := []struct {
		name string
		args args
		want Time
	}{
		{
			name: "success",
			args: args{
				createdAt: now,
				updatedAt: now,
				deletedAt: now,
			},
			want: Time{
				createdAt: now,
				updatedAt: now,
				deletedAt: now,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := NewTime(tt.args.createdAt, tt.args.updatedAt, tt.args.deletedAt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTime_CreatedAt(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name string
		tr   Time
		want time.Time
	}{
		{
			name: "success",
			tr: Time{
				createdAt: now,
			},
			want: now,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.tr.CreatedAt(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Time.CreatedAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTime_UpdatedAt(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name string
		tr   Time
		want time.Time
	}{
		{
			name: "success",
			tr: Time{
				updatedAt: now,
			},
			want: now,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.tr.UpdatedAt(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Time.UpdatedAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTime_DeletedAt(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name string
		tr   Time
		want time.Time
	}{
		{
			name: "success",
			tr: Time{
				deletedAt: now,
			},
			want: now,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.tr.DeletedAt(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Time.DeletedAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTime_IsZero(t *testing.T) {
	tests := []struct {
		name string
		tr   Time
		want bool
	}{
		{
			name: "success",
			tr: Time{
				createdAt: time.Time{},
				updatedAt: time.Time{},
				deletedAt: time.Time{},
			},
			want: true,
		},
		{
			name: "success, defaults",
			tr:   Time{},
			want: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.tr.IsZero(); got != tt.want {
				t.Errorf("Time.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTime_MarshalJSON(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name    string
		tr      Time
		want    []byte
		wantErr bool
	}{
		{
			name: "success",
			tr: Time{
				createdAt: now,
				updatedAt: now,
				deletedAt: now,
			},
			want: []byte(`{"created_at":"` + now.Format(time.RFC3339Nano) + `","updated_at":"` + now.Format(time.RFC3339Nano) + `","deleted_at":"` + now.Format(time.RFC3339Nano) + `"}`),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := tt.tr.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Time.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Time.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTime_UnmarshalJSON(t *testing.T) {
	now := time.Now()
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		tr      *Time
		args    args
		wantErr bool
	}{
		{
			name: "success",
			tr: &Time{
				createdAt: now,
				updatedAt: now,
				deletedAt: now,
			},
			args: args{
				data: []byte(`{"created_at":"` + now.Format(time.RFC3339Nano) + `","updated_at":"` + now.Format(time.RFC3339Nano) + `","deleted_at":"` + now.Format(time.RFC3339Nano) + `"}`),
			},
			wantErr: false,
		},
		{
			name: "failure, created_at",
			tr: &Time{
				createdAt: now,
				updatedAt: now,
				deletedAt: now,
			},
			args: args{
				data: []byte(`{"created_at":"1234","updated_at":"` + now.Format(time.RFC3339Nano) + `","deleted_at":"` + now.Format(time.RFC3339Nano) + `"}`),
			},
			wantErr: true,
		},
		{
			name: "failure, updated_at",
			tr: &Time{
				createdAt: now,
				updatedAt: now,
				deletedAt: now,
			},
			args: args{
				data: []byte(`{"created_at":"` + now.Format(time.RFC3339Nano) + `","updated_at":"1234","deleted_at":"` + now.Format(time.RFC3339Nano) + `"}`),
			},
			wantErr: true,
		},
		{
			name: "failure, deleted_at",
			tr: &Time{
				createdAt: now,
				updatedAt: now,
				deletedAt: now,
			},
			args: args{
				data: []byte(`{"created_at":"` + now.Format(time.RFC3339Nano) + `","updated_at":"` + now.Format(time.RFC3339Nano) + `","deleted_at":"1234"}`),
			},
			wantErr: true,
		},
		{
			name: "failure",
			tr: &Time{
				createdAt: now,
				updatedAt: now,
				deletedAt: now,
			},
			args: args{
				data: []byte(`{"}`),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := tt.tr.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Time.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
