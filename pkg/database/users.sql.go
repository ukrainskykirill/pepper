// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: users.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO users (
    login, name, phone
) VALUES (
    $1, $2, $3
)
`

type CreateUserParams struct {
	Login string
	Name  string
	Phone string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.Exec(ctx, createUser, arg.Login, arg.Name, arg.Phone)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, name, phone FROM users
WHERE id = $1 LIMIT 1
`

type GetUserRow struct {
	ID    uuid.UUID
	Name  string
	Phone string
}

func (q *Queries) GetUser(ctx context.Context, id uuid.UUID) (GetUserRow, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i GetUserRow
	err := row.Scan(&i.ID, &i.Name, &i.Phone)
	return i, err
}

const isExistsById = `-- name: IsExistsById :one
SELECT EXISTS (
    SELECT 1
    FROM users
    WHERE id = $1
)
`

func (q *Queries) IsExistsById(ctx context.Context, id uuid.UUID) (bool, error) {
	row := q.db.QueryRow(ctx, isExistsById, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const isExistsByLogin = `-- name: IsExistsByLogin :one
SELECT EXISTS (
    SELECT 1
    FROM users
    WHERE login = $1
)
`

func (q *Queries) IsExistsByLogin(ctx context.Context, login string) (bool, error) {
	row := q.db.QueryRow(ctx, isExistsByLogin, login)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, created_at, updated_at, login, name, phone FROM users
ORDER BY name
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Login,
			&i.Name,
			&i.Phone,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
