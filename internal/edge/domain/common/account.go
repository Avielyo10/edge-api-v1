package common

import (
	"context"
	"errors"
	"net/http"

	"github.com/redhatinsights/edge-api/config"
	"github.com/redhatinsights/platform-go-middlewares/identity"
)

// Account is the account of an user in Red Hat Hybrid Console.
type Account struct {
	id string
}

// DefaultAccount returns the default account.
var DefaultAccount = Account{"0000000"}

// ErrNoAccount is returned when no account is found in the context
var ErrNoAccount = errors.New("no account found in context")

// IsZero returns true if the account is empty.
func (a Account) IsZero() bool {
	return a == Account{}
}

// String get account id as string
func (a Account) String() string {
	return a.id
}

// GetAccount from http request header
func GetAccount(r *http.Request) (Account, error) {
	return GetAccountFromContext(r.Context())
}

// GetAccountFromContext determines account number from supplied context
func GetAccountFromContext(ctx context.Context) (Account, error) {
	if config.Get() != nil {
		if !config.Get().Auth {
			return DefaultAccount, nil
		}
		if ctx.Value(identity.Key) != nil {
			ident := identity.Get(ctx)
			if ident.Identity.AccountNumber != "" {
				return Account{ident.Identity.AccountNumber}, nil
			}
		}
	}
	return Account{}, ErrNoAccount
}

// MarshalJSON marshals account to json
func (a Account) MarshalJSON() ([]byte, error) {
	return []byte(`"` + a.id + `"`), nil
}

// UnmarshalJSON unmarshals json to account
func (a *Account) UnmarshalJSON(data []byte) error {
	if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		return ErrInvalidAccount{account: string(data)}
	}
	a.id = string(data[1 : len(data)-1])
	return nil
}

// ErrInvalidAccount is returned when account is invalid
type ErrInvalidAccount struct {
	account string
}

// Error returns error message
func (e ErrInvalidAccount) Error() string {
	return "invalid account: " + e.account
}
