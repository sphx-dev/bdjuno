package source

import (
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

type Source interface {
	GetParams(height int64) (authtypes.Params, error)
}
