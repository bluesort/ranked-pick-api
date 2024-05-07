-- name: CreateTokenHash :one
INSERT INTO token_hashes (
	user_id, hash, expires_at
) VALUES (
  ?, ?, ?
)
RETURNING *;

-- name: ReadTokenHashByHash :one
SELECT * FROM token_hashes
WHERE hash = ? LIMIT 1;

-- name: DeleteTokenHash :exec
DELETE FROM token_hashes
WHERE hash = ?;
