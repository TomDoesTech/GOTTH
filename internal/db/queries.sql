-- name: CreateUser :one
INSERT INTO users (firstname, lastname)
VALUES ($1, $2) RETURNING *;

-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: CreateTodo :one
INSERT INTO todos (user_id, task, done)
VALUES ($1, $2, $3) RETURNING *;

-- name: GetTodos :many
SELECT * FROM todos
WHERE user_id = $1;