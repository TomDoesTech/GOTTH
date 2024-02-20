-- name: CreateUser :one
INSERT INTO Users (email, "password")
VALUES ($1, $2) RETURNING *;

-- name: GetUsers :many
SELECT * FROM Users;

-- name: GetUser :one
SELECT * FROM Users
WHERE id = $1 LIMIT 1;

-- name: DeleteUser :exec
DELETE FROM Users
WHERE id = $1;
