-- name: GetMessages :many
SELECT content, user_ip FROM messages
ORDER BY id ASC;

-- name: SaveMessage :exec
INSERT INTO messages (content, user_ip)
VALUES (?, ?);