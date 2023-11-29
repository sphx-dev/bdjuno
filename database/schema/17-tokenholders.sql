CREATE TABLE account_denom_balance
(
    address         TEXT        NOT NULL,
    denom           TEXT        NOT NULL,
    amount          TEXT        NOT NULL,
    height          BIGINT  NOT NULL,
    PRIMARY KEY(address, denom)
);

CREATE VIEW token_holder_count AS
    SELECT denom, COUNT(*) AS holders
    FROM account_denom_balance
    WHERE amount != '0'
    GROUP BY denom;
