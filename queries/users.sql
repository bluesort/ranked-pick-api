-- name: ReadUser :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: ReadUserByEmail :one
SELECT * FROM users
WHERE email = ? LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id DESC LIMIT 100;

-- name: CreateUser :one
INSERT INTO users (
  password_hash, email, display_name
) VALUES (
  ?, ?, ?
)
RETURNING *;

-- name: UpdateUser :one
UPDATE users SET
email = coalesce(sqlc.narg('email'), email),
display_name = coalesce(sqlc.narg('display_name'), display_name),
updated_at = CURRENT_TIMESTAMP
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: UpdateUserPassword :one
UPDATE users SET
password_hash = ?,
updated_at = CURRENT_TIMESTAMP
WHERE id = ?
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;
