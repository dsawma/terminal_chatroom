-- name: DeleteMessage :one
DELETE FROM message
WHERE id = $1
RETURNING *;