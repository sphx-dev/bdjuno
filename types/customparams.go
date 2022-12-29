package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// CustomParamsStakingParams contains the staking parameters of the x/customparams module
type CustomParamsStakingParams struct {
	MinSelfDelegation sdk.Int `json:"min_self_delegation,omitempty" yaml:"min_self_delegation"`
}

// CustomParamsParams contains the data of the x/customparams module parameters
type CustomParamsParams struct {
	StakingParams CustomParamsStakingParams `json:"staking_params" yaml:"staking_params"`
	Height        int64                     `json:"height" ymal:"height"`
}

// NewCustomParamsParams allows to build a new CustomParamsParams instance
func NewCustomParamsParams(stakingParams CustomParamsStakingParams, height int64) *CustomParamsParams {
	return &CustomParamsParams{
		StakingParams: stakingParams,
		Height:        height,
	}
}
