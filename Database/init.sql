CREATE TABLE wallets (
    wallet_address VARCHAR(42) PRIMARY KEY,
    balance BIGINT NOT NULL
);

CREATE TABLE transfers (
    id BIGSERIAL PRIMARY KEY, 
    from_address VARCHAR(42) NOT NULL,
    to_address VARCHAR(42) NOT NULL,
    amount BIGINT NOT NULL
);

INSERT INTO wallets (wallet_address, balance)
VALUES ('0x0000000000000000000000000000000000000000', 1000000);