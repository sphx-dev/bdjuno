package database

import (
	"encoding/json"
	"fmt"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	dbtypes "github.com/forbole/bdjuno/v4/database/types"
	"github.com/forbole/bdjuno/v4/types"
)

// SaveAuthParams allows to store the given params into the database.
func (db *Db) SaveAuthParams(params types.AuthParams) error {
	paramsBz, err := json.Marshal(&params.Params)
	if err != nil {
		return fmt.Errorf("error while marshaling auth params: %s", err)
	}

	stmt := `
INSERT INTO auth_params (params, height) 
VALUES ($1, $2)
ON CONFLICT (one_row_id) DO UPDATE 
    SET params = excluded.params,
        height = excluded.height
WHERE auth_params.height <= excluded.height`

	_, err = db.SQL.Exec(stmt, string(paramsBz), params.Height)
	if err != nil {
		return fmt.Errorf("error while storing auth params: %s", err)
	}

	return nil
}

// GetAuthParams returns the types.AuthParams instance containing the current params
func (db *Db) GetAuthParams() (*types.AuthParams, error) {
	var rows []dbtypes.AuthParamsRow
	stmt := `SELECT * FROM auth_params LIMIT 1`
	err := db.Sqlx.Select(&rows, stmt)
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return nil, fmt.Errorf("no auth params found")
	}

	var authParams authtypes.Params
	err = json.Unmarshal([]byte(rows[0].Params), &authParams)
	if err != nil {
		return nil, err
	}

	return &types.AuthParams{
		Params: authParams,
		Height: rows[0].Height,
	}, nil
}
