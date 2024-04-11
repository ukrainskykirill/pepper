-- name: GetUser :one
SELECT id, name, phone FROM users
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

-- name: IsExistsByLogin :one
SELECT EXISTS (
    SELECT 1
    FROM users
    WHERE login = $1
);

-- name: IsExistsById :one
SELECT EXISTS (
    SELECT 1
    FROM users
    WHERE id = $1
);

-- name: UpadateUser :exec
UPDATE users
SET name = $1, discription = $2
WHERE id = $3;