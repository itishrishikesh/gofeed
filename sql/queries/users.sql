-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name, password, api_key)
VALUES ($1, $2, $3, $4, $5, encode(sha256(random()::text::bytea), 'hex'))
RETURNING *;

-- name: GetUserByAPIKey :one
SELECT * FROM users WHERE api_key = $1;

-- name: UpdateAPIKey :one
UPDATE users 
SET api_key = encode(sha256(random()::text::bytea) 
WHERE id = $1