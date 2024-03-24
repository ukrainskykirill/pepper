-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY name;

-- name: CreateUser :exec
INSERT INTO users (
    login, name, phone
) VALUES (
    $1, $2, $3
);