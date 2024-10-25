CREATE TABLE order_event (
    id            TEXT PRIMARY KEY,
    account_id    TEXT NOT NULL,
    side          TEXT NOT NULL,
    quantity      TEXT NOT NULL,
    price         TEXT NOT NULL,
    order_type    TEXT NOT NULL,
    trigger_price TEXT,
    leverage      TEXT,
    timestamp     TEXT NOT NULL,
    market_id     TEXT NOT NULL,
    status        TEXT NOT NULL
);
