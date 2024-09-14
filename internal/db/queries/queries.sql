-- name: CreateExampleData :one
INSERT INTO example_data (content) VALUES ($1) RETURNING id, content;

-- name: GetExampleData :one
SELECT id, content FROM example_data WHERE id = $1;
