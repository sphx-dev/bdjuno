package source

import (
	assetfttypes "github.com/CoreumFoundation/coreum/v3/x/asset/ft/types"
)

type Source interface {
	GetParams(height int64) (assetfttypes.Params, error)
}
