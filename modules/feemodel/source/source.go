package source

import (
	feemodeltypes "github.com/CoreumFoundation/coreum/x/feemodel/types"
)

type Source interface {
	GetParams(height int64) (feemodeltypes.Params, error)
}
