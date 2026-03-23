-- name: GetRoomByRoomName :one
SELECT *
FROM room
WHERE room_name = $1;