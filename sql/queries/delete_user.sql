-- name: DeleteUserByUsername :one
DELETE FROM users
WHERE username = $1
RETURNING *;