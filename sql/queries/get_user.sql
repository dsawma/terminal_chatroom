-- name: GetUserByUsername :one
SELECT id, username, created_at, updated_at, hashed_password
FROM users
WHERE username = $1;