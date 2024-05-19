-- name: CreateUser :one
INSERT INTO users (
  password_hash, username, display_name
) VALUES (
  ?, ?, ?
)
RETURNING *;

-- name: ReadUser :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: ReadUserByUsername :one
SELECT * FROM users
WHERE username = ? LIMIT 1;

-- name: UpdateUser :one
UPDATE users SET
username = coalesce(sqlc.narg('username'), username),
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
