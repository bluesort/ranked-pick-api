-- TODO: Add unique survey_responses survey_id,user_id,rank index

-- name: UpsertSurveyResponse :one
INSERT INTO survey_responses (
	survey_id, survey_option_id, user_id, rank
) VALUES (
  ?, ?, ?, ?
)
ON CONFLICT (user_id, survey_option_id) DO UPDATE SET
	survey_id = EXCLUDED.survey_id,
	survey_option_id = EXCLUDED.survey_option_id,
	user_id = EXCLUDED.user_id,
	rank = EXCLUDED.rank
RETURNING *;

-- name: ListSurveyResponsesForSurvey :many
SELECT * FROM survey_responses
WHERE survey_id = ?
ORDER BY id ASC LIMIT 100;

-- name: ListSurveyResponsesForSurveyUser :many
SELECT * FROM survey_responses
WHERE survey_id = ?
AND user_id = ?
ORDER BY rank ASC LIMIT 100;

-- name: CountSurveyResponsesForSurvey :one
SELECT COUNT(DISTINCT user_id) FROM survey_responses
WHERE survey_id = ?;

-- name: UpdateSurveyResponse :one
UPDATE survey_responses SET
survey_id = coalesce(sqlc.narg('survey_id'), survey_id),
survey_option_id = coalesce(sqlc.narg('survey_option_id'), survey_option_id),
user_id = coalesce(sqlc.narg('user_id'), user_id),
rank = coalesce(sqlc.narg('rank'), rank)
WHERE id = ?
RETURNING *;

-- name: DeleteSurveyResponse :exec
DELETE FROM survey_responses
WHERE id = ?;

-- name: DeleteSurveyResponsesForSurvey :exec
DELETE FROM survey_responses
WHERE survey_id = ?;

-- name: AnonymizeSurveyResponsesForUser :exec
UPDATE survey_responses SET
user_id = NULL
WHERE user_id = ?;
