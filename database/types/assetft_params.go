package types

// AssetFTParamsRow represents a single row inside the assetft_params table
type AssetFTParamsRow struct {
	OneRowID bool   `db:"one_row_id"`
	Params   string `db:"params"`
	Height   int64  `db:"height"`
}
