package source

import (
	customparamstypes "github.com/CoreumFoundation/coreum/v4/x/customparams/types"
)

type Source interface {
	GetParams(height int64) (customparamstypes.StakingParams, error)
}
