package source

import (
	feemodeltypes "github.com/CoreumFoundation/coreum/v3/x/feemodel/types"
)

type Source interface {
	GetParams(height int64) (feemodeltypes.Params, error)
}
