package types

// AuthParamsRow represents a single row inside the auth_params table
type AuthParamsRow struct {
	OneRowID bool   `db:"one_row_id"`
	Params   string `db:"params"`
	Height   int64  `db:"height"`
}
