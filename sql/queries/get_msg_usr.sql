-- name: GetMessagesUser :one
SELECT * FROM message
WHERE user_id = $1;