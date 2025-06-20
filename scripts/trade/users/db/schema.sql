CREATE SCHEMA IF NOT EXISTS users;

CREATE TABLE users.users (
    -- The auto incrementing primary key.
    id uuid PRIMARY KEY,

    -- Privy ID for the account, which is unique and used for authentication.
    privy_id TEXT NOT NULL UNIQUE,

    -- Wallet address associated with the account, which is also unique.
    address TEXT NOT NULL UNIQUE
);
