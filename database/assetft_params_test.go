package database_test

import (
	"encoding/json"

	assetfttypes "github.com/CoreumFoundation/coreum/x/asset/ft/types"
	dbtypes "github.com/forbole/bdjuno/v3/database/types"
	"github.com/forbole/bdjuno/v3/types"
)

func (suite *DbTestSuite) TestSaveAssetFTParams() {
	assetFTParams := assetfttypes.DefaultParams()
	err := suite.database.SaveAssetFTParams(types.NewAssetFTParams(assetFTParams, 10))
	suite.Require().NoError(err)

	var rows []dbtypes.AssetFTParamsRow
	err = suite.database.Sqlx.Select(&rows, `SELECT * FROM assetft_params`)
	suite.Require().NoError(err)

	suite.Require().Len(rows, 1)

	var stored assetfttypes.Params
	err = json.Unmarshal([]byte(rows[0].Params), &stored)
	suite.Require().NoError(err)
	suite.Require().Equal(assetFTParams, stored)
}

func (suite *DbTestSuite) TestGetAssetFTParams() {
	assetFTParams := assetfttypes.DefaultParams()

	paramsBz, err := json.Marshal(&assetFTParams)
	suite.Require().NoError(err)

	_, err = suite.database.Sql.Exec(
		`INSERT INTO assetft_params (params, height) VALUES ($1, $2)`,
		string(paramsBz), 10,
	)
	suite.Require().NoError(err)

	params, err := suite.database.GetAssetFTParams()
	suite.Require().NoError(err)

	suite.Require().Equal(&types.AssetFTParams{
		Params: assetFTParams,
		Height: 10,
	}, params)
}
