package gov

import (
	"github.com/forbole/bdjuno/v4/types"
)

type AuthModule interface {
	UpdateParams(height int64) error
}

type DistrModule interface {
	UpdateParams(height int64) error
}

type MintModule interface {
	UpdateParams(height int64) error
	UpdateInflation() error
}

type SlashingModule interface {
	UpdateParams(height int64) error
}

type StakingModule interface {
	GetStakingPoolSnapshot(height int64) (*types.PoolSnapshot, error)
	UpdateParams(height int64) error
}

type FeeModelModule interface {
	UpdateParams(height int64) error
}

type CustomParamsModule interface {
	UpdateParams(height int64) error
}

type AssetFTModule interface {
	UpdateParams(height int64) error
}

type AssetNFTModule interface {
	UpdateParams(height int64) error
}
