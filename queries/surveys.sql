-- name: CreateSurvey :one
INSERT INTO surveys (
  user_id, title, state, visibility, description
) VALUES (
  ?, ?, ?, ?, ?
)
RETURNING *;

-- name: ReadSurvey :one
SELECT surveys.*, COUNT(survey_responses.id) AS response_count FROM surveys
LEFT JOIN survey_responses
ON survey_responses.survey_id = surveys.id
AND survey_responses.rank = 0
WHERE surveys.id = ?
GROUP BY surveys.id LIMIT 1;

-- name: ListSurveys :many
SELECT surveys.*, COUNT(survey_responses.id) AS response_count FROM surveys
LEFT JOIN survey_responses
ON survey_responses.survey_id = surveys.id
AND survey_responses.rank = 0
ORDER BY surveys.id DESC LIMIT 100;

-- name: ListSurveysForUser :many
SELECT surveys.*, COUNT(survey_responses.id) AS response_count FROM surveys
LEFT JOIN survey_responses
ON survey_responses.survey_id = surveys.id
AND survey_responses.rank = 0
WHERE surveys.user_id = ?
GROUP BY surveys.id
ORDER BY surveys.id DESC LIMIT 100;

-- name: ListSurveysForUserResponded :many
SELECT surveys.*, COUNT(sr.id) AS response_count FROM surveys
JOIN survey_responses sr
ON sr.survey_id = surveys.id
AND sr.rank = 0
WHERE EXISTS (
  SELECT 1 FROM survey_responses WHERE survey_responses.survey_id = surveys.id AND survey_responses.user_id = ?
)
GROUP BY surveys.id
ORDER BY surveys.id DESC LIMIT 100;

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
