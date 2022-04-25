package common

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/redhatinsights/edge-api/config"
	"github.com/redhatinsights/platform-go-middlewares/identity"
)

func TestAccount_IsZero(t *testing.T) {
	tests := []struct {
		name string
		a    Account
		want bool
	}{
		{
			name: "empty",
			a:    Account{},
			want: true,
		},
		{
			name: "not empty",
			a:    DefaultAccount,
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.a.IsZero(); got != tt.want {
				t.Errorf("Account.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccount_String(t *testing.T) {
	tests := []struct {
		name string
		a    Account
		want string
	}{
		{
			name: "empty",
			a:    Account{},
			want: "",
		},
		{
			name: "not empty",
			a:    DefaultAccount,
			want: "0000000",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.a.String(); got != tt.want {
				t.Errorf("Account.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAccount(t *testing.T) {
	config.Init()

	validReq, _ := http.NewRequest("GET", "/", nil)
	validReq = validReq.WithContext(context.WithValue(validReq.Context(), identity.Key, identity.XRHID{
		Identity: identity.Identity{
			AccountNumber: "0000000",
		},
	}))

	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		args    args
		want    Account
		auth    bool
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				r: validReq,
			},
			want:    Account{id: "0000000"},
			wantErr: false,
			auth:    false,
		},
		{
			name: "ok with auth",
			args: args{
				r: validReq,
			},
			want:    Account{id: "0000000"},
			wantErr: false,
			auth:    true,
		},
		{
			name: "no valid account, with auth",
			args: args{
				r: &http.Request{},
			},
			want:    Account{},
			wantErr: true,
			auth:    true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			config.Get().Auth = tt.auth
			got, err := GetAccount(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAccountFromContext(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    Account
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				ctx: context.WithValue(context.Background(), identity.Key, identity.XRHID{
					Identity: identity.Identity{
						AccountNumber: "0000000",
					},
				}),
			},
			want:    Account{id: "0000000"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := GetAccountFromContext(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccountFromContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccountFromContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccount_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		a       Account
		want    []byte
		wantErr bool
	}{
		{
			name:    "empty",
			a:       Account{},
			want:    []byte(`""`),
			wantErr: false,
		},
		{
			name:    "not empty",
			a:       DefaultAccount,
			want:    []byte(`"0000000"`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := tt.a.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Account.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Account.MarshalJSON() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestAccount_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		a       *Account
		args    args
		wantErr bool
	}{
		{
			name:    "empty",
			a:       &Account{},
			args:    args{data: []byte(`""`)},
			wantErr: false,
		},
		{
			name:    "not empty",
			a:       &DefaultAccount,
			args:    args{data: []byte(`"0000000"`)},
			wantErr: false,
		},
		{
			name:    "invalid",
			a:       &Account{},
			args:    args{data: []byte(`"`)},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := tt.a.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Account.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestErrInvalidAccount_Error(t *testing.T) {
	tests := []struct {
		name string
		e    ErrInvalidAccount
		want string
	}{
		{
			name: "ok",
			e: ErrInvalidAccount{
				account: "0000000",
			},
			want: `invalid account: 0000000`,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.e.Error(); got != tt.want {
				t.Errorf("ErrInvalidAccount.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
