-- name: DeleteRoom :one
DELETE FROM rooms
WHERE id = $1
RETURNING *;