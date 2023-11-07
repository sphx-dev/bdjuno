package database

import (
	"encoding/json"
	"fmt"

	"github.com/forbole/bdjuno/v4/types"

	dbtypes "github.com/forbole/bdjuno/v4/database/types"
)

// SaveCustomParamsParams saves the given x/customparams parameters inside the database
func (db *Db) SaveCustomParamsParams(params *types.CustomParamsParams) error {
	stakingParamsBz, err := json.Marshal(&params.StakingParams)
	if err != nil {
		return fmt.Errorf("error while marshaling staking params: %s", err)
	}

	stmt := `
INSERT INTO customparams_params(staking_params, height) 
VALUES ($1, $2)
ON CONFLICT (one_row_id) DO UPDATE 
	SET staking_params = excluded.staking_params,
		height = excluded.height
WHERE customparams_params.height <= excluded.height`
	_, err = db.SQL.Exec(stmt, string(stakingParamsBz), params.Height)
	if err != nil {
		return fmt.Errorf("error while storing customparams params: %s", err)
	}

	return nil
}

// GetCustomParamsParams returns the most recent customparamsernance parameters
func (db *Db) GetCustomParamsParams() (*types.CustomParamsParams, error) {
	var rows []dbtypes.CustomParamsParamsRow
	err := db.Sqlx.Select(&rows, `SELECT * FROM customparams_params`)
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return nil, nil
	}

	row := rows[0]

	var stakingParams types.CustomParamsStakingParams
	err = json.Unmarshal([]byte(row.StakingParams), &stakingParams)
	if err != nil {
		return nil, err
	}

	return types.NewCustomParamsParams(stakingParams, row.Height), nil
}
