INSERT INTO room_member(room_id, user_id)
VALUES (
    $1,
    $2
)
RETURNING *;