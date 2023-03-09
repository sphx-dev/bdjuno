package types

// AssetNFTParamsRow represents a single row inside the assetnft_params table
type AssetNFTParamsRow struct {
	OneRowID bool   `db:"one_row_id"`
	Params   string `db:"params"`
	Height   int64  `db:"height"`
}
