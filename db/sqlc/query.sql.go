// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package sqlc

import (
	"context"
)

const getMessages = `-- name: GetMessages :many
SELECT content, user_ip FROM messages
ORDER BY id ASC
`

type GetMessagesRow struct {
	Content string
	UserIp  string
}

func (q *Queries) GetMessages(ctx context.Context) ([]GetMessagesRow, error) {
	rows, err := q.db.QueryContext(ctx, getMessages)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetMessagesRow
	for rows.Next() {
		var i GetMessagesRow
		if err := rows.Scan(&i.Content, &i.UserIp); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const saveMessage = `-- name: SaveMessage :exec
INSERT INTO messages (content, user_ip)
VALUES (?, ?)
`

type SaveMessageParams struct {
	Content string
	UserIp  string
}

func (q *Queries) SaveMessage(ctx context.Context, arg SaveMessageParams) error {
	_, err := q.db.ExecContext(ctx, saveMessage, arg.Content, arg.UserIp)
	return err
}
