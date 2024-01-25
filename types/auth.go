package types

import authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

// AuthParams represents the parameters of the x/auth module
type AuthParams struct {
	Params authtypes.Params
	Height int64
}

// NewAuthParams returns a new AuthParams instance
func NewAuthParams(params authtypes.Params, height int64) AuthParams {
	return AuthParams{
		Params: params,
		Height: height,
	}
}

// Account represents a chain account
type Account struct {
	Address string
}

// NewAccount builds a new Account instance
func NewAccount(address string) Account {
	return Account{
		Address: address,
	}
}
