CREATE TABLE position_event (
    id            TEXT PRIMARY KEY,
    margin_account    TEXT NOT NULL,
    market_id      TEXT NOT NULL,
    size          TEXT NOT NULL,
    entry_price    TEXT NOT NULL,
    leverage      TEXT NOT NULL,
    entry_time     TEXT NOT NULL,
    side          TEXT NOT NULL,
    tp_order_id     TEXT,
    sl_order_id     TEXT,
    status        TEXT NOT NULL
);

