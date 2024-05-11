-- TODO: Add unique survey_answers survey_id,user_id,rank index

-- name: UpsertSurveyAnswer :one
INSERT INTO survey_answers (
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

-- name: ListSurveyAnswersForSurvey :many
SELECT * FROM survey_answers
WHERE survey_id = ?
ORDER BY id ASC LIMIT 100;

-- name: ListSurveyAnswersForSurveyUser :many
SELECT * FROM survey_answers
WHERE survey_id = ?
AND user_id = ?
ORDER BY rank ASC LIMIT 100;

-- name: UpdateSurveyAnswer :one
UPDATE survey_answers SET
survey_id = coalesce(sqlc.narg('survey_id'), survey_id),
survey_option_id = coalesce(sqlc.narg('survey_option_id'), survey_option_id),
user_id = coalesce(sqlc.narg('user_id'), user_id),
rank = coalesce(sqlc.narg('rank'), rank)
WHERE id = ?
RETURNING *;

-- name: DeleteSurveyAnswer :exec
DELETE FROM survey_answers
WHERE id = ?;
