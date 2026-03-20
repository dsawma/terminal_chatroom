-- name: GetRoomByRoomName :one
SELECT room_name
FROM room
WHERE room_name = $1;