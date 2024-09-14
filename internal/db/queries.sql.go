// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: queries.sql

package db

import (
	"context"
)

const createExampleData = `-- name: CreateExampleData :one
INSERT INTO example_data (content) VALUES ($1) RETURNING id, content
`

func (q *Queries) CreateExampleData(ctx context.Context, content string) (ExampleDatum, error) {
	row := q.db.QueryRowContext(ctx, createExampleData, content)
	var i ExampleDatum
	err := row.Scan(&i.ID, &i.Content)
	return i, err
}

const getExampleData = `-- name: GetExampleData :one
SELECT id, content FROM example_data WHERE id = $1
`

func (q *Queries) GetExampleData(ctx context.Context, id int32) (ExampleDatum, error) {
	row := q.db.QueryRowContext(ctx, getExampleData, id)
	var i ExampleDatum
	err := row.Scan(&i.ID, &i.Content)
	return i, err
}
