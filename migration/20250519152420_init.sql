-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    login VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS orders (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id),
    number VARCHAR(255) NOT NULL UNIQUE,
    status VARCHAR(50) NOT NULL,
    accrual DECIMAL(10,2) DEFAULT 0,
    uploaded_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS user_balances (
    user_id UUID PRIMARY KEY REFERENCES users(id),
    current DECIMAL(10,2) NOT NULL DEFAULT 0,
    withdrawn DECIMAL(10,2) NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS withdrawals (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id),
    order_number VARCHAR(255) NOT NULL,
    sum DECIMAL(10,2) NOT NULL,
    processed_at TIMESTAMP NOT NULL
); 
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS withdrawals;
DROP TABLE IF EXISTS user_balances;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS users; 
-- +goose StatementEnd
