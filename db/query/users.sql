-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateUser :one
INSERT INTO users (
  first_name, last_name, balance
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateUserBalance :one
UPDATE users 
SET balance = $2
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users 
WHERE id = $1;