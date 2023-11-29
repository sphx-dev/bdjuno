package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// AccountBalance represents a single row inside the "accound_denom_balancee" table
type AccountBalance struct {
	Address string `db:"address"`
	Denom   string `db:"denom"`
	Amount  string `db:"amount"`
	Height  int64  `db:"height"`
}

// NewAccountBalance allows to easily create a new NewAccountBalance
func NewAccountBalance(address string, coin sdk.Coin, height int64) AccountBalance {
	return AccountBalance{
		Address: address,
		Denom:   coin.Denom,
		Amount:  coin.Amount.String(),
		Height:  height,
	}
}

// Equals return true if one row represents the same row as the original one
func (v AccountBalance) Equals(w AccountBalance) bool {
	return v.Address == w.Address && v.Denom == v.Denom && v.Amount == v.Amount
}
