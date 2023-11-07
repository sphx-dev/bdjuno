package source

import (
	assetnfttypes "github.com/CoreumFoundation/coreum/v3/x/asset/nft/types"
)

type Source interface {
	GetParams(height int64) (assetnfttypes.Params, error)
}
