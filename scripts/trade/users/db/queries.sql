-- name: Insert :one
INSERT INTO users.users (
    id, privy_id, address
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: GetByID :one
SELECT * FROM users.users
WHERE id = $1 LIMIT 1;
-- No index required for GetByID as it is a primary key lookup.

-- name: GetByPrivyID :one
SELECT * FROM users.users
WHERE privy_id = $1 LIMIT 1;
-- No index required for GetByPrivyID as it is a unique lookup (auto indexed).

-- name: GetByWalletAddress :one
SELECT * FROM users.users
WHERE address = $1 LIMIT 1;
-- No index required for GetByWalletAddress as it is a unique lookup (auto indexed).

-- name: ListAll :many
SELECT * FROM users.users;
-- No index required for ListAll as it retrieves all records.
