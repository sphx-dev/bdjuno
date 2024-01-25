package database_test

import (
	"encoding/json"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	dbtypes "github.com/forbole/bdjuno/v4/database/types"
	"github.com/forbole/bdjuno/v4/types"
)

func (suite *DbTestSuite) TestSaveAuthParams() {
	authParams := authtypes.DefaultParams()
	err := suite.database.SaveAuthParams(types.NewAuthParams(authParams, 10))
	suite.Require().NoError(err)

	var rows []dbtypes.AuthParamsRow
	err = suite.database.Sqlx.Select(&rows, `SELECT * FROM auth_params`)
	suite.Require().NoError(err)

	suite.Require().Len(rows, 1)

	var stored authtypes.Params
	err = json.Unmarshal([]byte(rows[0].Params), &stored)
	suite.Require().NoError(err)
	suite.Require().Equal(authParams, stored)
}

func (suite *DbTestSuite) TestGetAuthParams() {
	authParams := authtypes.DefaultParams()

	paramsBz, err := json.Marshal(&authParams)
	suite.Require().NoError(err)

	_, err = suite.database.SQL.Exec(
		`INSERT INTO auth_params (params, height) VALUES ($1, $2)`,
		string(paramsBz), 10,
	)
	suite.Require().NoError(err)

	params, err := suite.database.GetAuthParams()
	suite.Require().NoError(err)

	suite.Require().Equal(&types.AuthParams{
		Params: authParams,
		Height: 10,
	}, params)
}
