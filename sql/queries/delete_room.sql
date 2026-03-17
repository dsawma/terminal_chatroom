-- name: DeleteRoom :one
DELETE FROM room
WHERE id = $1
RETURNING *;