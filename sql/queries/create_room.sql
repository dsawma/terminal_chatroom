INSERT INTO room(id, created_at, updated_at,room_name, owner_id)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2
)
RETURNING *;