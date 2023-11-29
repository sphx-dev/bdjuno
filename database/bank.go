package database

import (
	"fmt"

	"github.com/forbole/bdjuno/v4/database/types"
	dbtypes "github.com/forbole/bdjuno/v4/database/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lib/pq"
)

// SaveSupply allows to save for the given height the given total amount of coins
func (db *Db) SaveSupply(coins sdk.Coins, height int64) error {
	query := `
INSERT INTO supply (coins, height) 
VALUES ($1, $2) 
ON CONFLICT (one_row_id) DO UPDATE 
    SET coins = excluded.coins,
    	height = excluded.height
WHERE supply.height <= excluded.height`

	_, err := db.SQL.Exec(query, pq.Array(dbtypes.NewDbCoins(coins)), height)
	if err != nil {
		return fmt.Errorf("error while storing supply: %s", err)
	}

	return nil
}

// GetAccountDenomBalance allows to save the balance of an account for a given denom.
func (db *Db) GetAccountDenomBalance(account string, denom string) (types.AccountBalance, bool, error) {
	var vals []types.AccountBalance
	query := `
SELECT address, denom, amount, height FROM account_denom_balance  
WHERE account_denom_balance.address=$1 AND account_denom_balance.denom=$2`

	err := db.SQL.Select(&vals, query, account, denom)
	if err != nil {
		return types.AccountBalance{}, false, err
	}
	if len(vals) == 0 {
		return types.AccountBalance{}, false, nil
	}
	return vals[0], true, nil
}

// SaveAccountDenomBalance allows to save the balance of an account for a given denom.
func (db *Db) SaveAccountDenomBalance(account string, coin sdk.Coin, height int64) error {
	query := `
INSERT INTO account_denom_balance (address, denom, amount, height) 
VALUES ($1, $2, $3, $4) 
ON CONFLICT (address,denom) DO UPDATE 
    SET amount = $3, height = $4
WHERE 
	account_denom_balance.address = $1 
	AND account_denom_balance.denom = $2 
	AND account_denom_balance.height < $4`

	_, err := db.SQL.Exec(query, account, coin.Denom, coin.Amount.String(), height)
	if err != nil {
		return fmt.Errorf("error while storing account balance: %s", err)
	}

	return nil
}
