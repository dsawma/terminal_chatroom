-- name: GetMessagesRoom :one
SELECT * FROM message
WHERE room_id = $1;