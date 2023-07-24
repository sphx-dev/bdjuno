package database

import (
	"encoding/json"
	"fmt"

	feemodeltypes "github.com/CoreumFoundation/coreum/v2/x/feemodel/types"
	dbtypes "github.com/forbole/bdjuno/v3/database/types"
	"github.com/forbole/bdjuno/v3/types"
)

// SaveFeeModelParams allows to store the given params into the database.
func (db *Db) SaveFeeModelParams(params types.FeeModelParams) error {
	paramsBz, err := json.Marshal(&params.Params)
	if err != nil {
		return fmt.Errorf("error while marshaling feemodel params: %s", err)
	}

	stmt := `
INSERT INTO feemodel_params (params, height) 
VALUES ($1, $2)
ON CONFLICT (one_row_id) DO UPDATE 
    SET params = excluded.params,
        height = excluded.height
WHERE feemodel_params.height <= excluded.height`

	_, err = db.Sql.Exec(stmt, string(paramsBz), params.Height)
	if err != nil {
		return fmt.Errorf("error while storing feemodel params: %s", err)
	}

	return nil
}

// GetFeeModelParams returns the types.FeeModelParams instance containing the current params
func (db *Db) GetFeeModelParams() (*types.FeeModelParams, error) {
	var rows []dbtypes.FeeModelParamsRow
	stmt := `SELECT * FROM feemodel_params LIMIT 1`
	err := db.Sqlx.Select(&rows, stmt)
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return nil, fmt.Errorf("no feemodel params found")
	}

	var feemodelParams feemodeltypes.Params
	err = json.Unmarshal([]byte(rows[0].Params), &feemodelParams)
	if err != nil {
		return nil, err
	}

	return &types.FeeModelParams{
		Params: feemodelParams,
		Height: rows[0].Height,
	}, nil
}
