CREATE TABLE wallet_balance (
    id uuid,
    balance float8,

    PRIMARY KEY (id)
);

CREATE TABLE transaction_history (
    "time" time,
    "from" uuid,
    "to" uuid,
    amount float8,

    FOREIGN KEY ("from") REFERENCES wallet_balance(id),
    FOREIGN KEY ("to") REFERENCES wallet_balance(id)
);
