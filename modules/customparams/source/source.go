package source

import (
	customparamstypes "github.com/CoreumFoundation/coreum/x/customparams/types"
)

type Source interface {
	GetParams(height int64) (customparamstypes.StakingParams, error)
}
