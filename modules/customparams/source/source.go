package source

import (
	customparamstypes "github.com/CoreumFoundation/coreum/v2/x/customparams/types"
)

type Source interface {
	GetParams(height int64) (customparamstypes.StakingParams, error)
}
