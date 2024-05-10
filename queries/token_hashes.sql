-- name: UpsertTokenHash :one
INSERT INTO token_hashes (
	user_id, hash, expires_at
) VALUES (
  ?, ?, ?
)
ON CONFLICT (user_id) DO UPDATE SET
	hash = EXCLUDED.hash,
	expires_at = EXCLUDED.expires_at
RETURNING *;

-- name: ReadTokenHashByUserId :one
SELECT * FROM token_hashes
WHERE user_id = ? LIMIT 1;

-- name: DeleteTokenHash :exec
DELETE FROM token_hashes
WHERE hash = ?;
