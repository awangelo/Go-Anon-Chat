-- name: GetMessages :many
SELECT content, user_ip FROM messages
ORDER BY id ASC;