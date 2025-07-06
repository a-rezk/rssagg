-- name: CreateUser :execresult
INSERT INTO users (id, created_at, updated_at, name, api_key)
VALUES (?, ?, ?, ?, SHA2(UUID(), 256));

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = ?;

-- name: GetUserByAPIkey :one
SELECT * FROM users
WHERE api_key = ?;