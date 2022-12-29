package types

// CustomParamsParamsRow represents a single row inside the customparams_params table
type CustomParamsParamsRow struct {
	OneRowID      bool   `db:"one_row_id"`
	StakingParams string `db:"staking_params"`
	Height        int64  `db:"height"`
}
