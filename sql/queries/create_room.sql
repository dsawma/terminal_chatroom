-- name: CreateRoom :one
INSERT INTO room(id, created_at, updated_at,room_name)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1
)
RETURNING *;