-- name: GetRoomByRoomName :one
SELECT *
FROM rooms
WHERE room_name = $1;