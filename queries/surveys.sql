-- name: CreateSurvey :one
INSERT INTO surveys (
  user_id, title, state, visibility, description
) VALUES (
  ?, ?, ?, ?, ?
)
RETURNING *;

-- name: ReadSurvey :one
SELECT * FROM surveys
WHERE id = ? LIMIT 1;

-- name: ListSurveys :many
SELECT * FROM surveys
ORDER BY id DESC LIMIT 100;

-- name: ListSurveysForUser :many
SELECT * FROM surveys
WHERE user_id = ?
ORDER BY id DESC LIMIT 100;

-- name: UpdateSurvey :one
UPDATE surveys SET
user_id = coalesce(sqlc.narg('user_id'), user_id),
title = coalesce(sqlc.narg('title'), title),
state = coalesce(sqlc.narg('state'), state),
visibility = coalesce(sqlc.narg('visibility'), visibility),
description = coalesce(sqlc.narg('description'), description),
updated_at = CURRENT_TIMESTAMP
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: DeleteSurvey :exec
DELETE FROM surveys
WHERE id = ?;
