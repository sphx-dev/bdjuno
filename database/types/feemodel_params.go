package types

// FeeModelParamsRow represents a single row inside the feemodel_params table
type FeeModelParamsRow struct {
	OneRowID bool   `db:"one_row_id"`
	Params   string `db:"params"`
	Height   int64  `db:"height"`
}
