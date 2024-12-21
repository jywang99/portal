// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package db

import (
	"context"
)

const getUser = `-- name: GetUser :one
SELECT user_id, name, email, password FROM users WHERE email = $1
`

func (q *Queries) GetUser(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUser, email)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Name,
		&i.Email,
		&i.Password,
	)
	return i, err
}
