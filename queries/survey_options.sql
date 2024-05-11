-- name: CreateSurveyOption :one
INSERT INTO survey_options (
  survey_id, title
) VALUES (
  ?, ?
)
RETURNING *;

-- name: ReadSurveyOption :one
SELECT * FROM survey_options
WHERE id = ? LIMIT 1;

-- name: ListSurveyOptionsForSurvey :many
SELECT * FROM survey_options
WHERE survey_id = ?
ORDER BY id ASC LIMIT 100;

-- name: UpdateSurveyOption :one
UPDATE survey_options SET
title = ?
WHERE id = ?
RETURNING *;

-- name: DeleteSurveyOption :exec
DELETE FROM survey_options
WHERE id = ?;
