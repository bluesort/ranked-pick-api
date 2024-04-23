package surveys

import (
	"database/sql"

	"github.com/carterjackson/ranked-pick-api/internal/common"
	"github.com/carterjackson/ranked-pick-api/internal/db"
	"github.com/carterjackson/ranked-pick-api/internal/resources"
)

type CreateParams struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func Create(ctx *common.Context, tx *db.Queries, iparams interface{}) (interface{}, error) {
	params := iparams.(*CreateParams)

	description := sql.NullString{}
	if params.Description != "" {
		description = sql.NullString{String: params.Description, Valid: true}
	}

	survey, err := tx.CreateSurvey(ctx, db.CreateSurveyParams{
		UserID:      0,
		Title:       params.Title,
		Description: description,
		State:       resources.SurveyStatePending,
		Visibility:  resources.SurveyVisibilityPublic,
	})
	if err != nil {
		return nil, err
	}

	return survey, nil
}
