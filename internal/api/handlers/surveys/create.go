package surveys

import (
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

	// TODO: Make surveys private by default once invites are implemented
	survey, err := tx.CreateSurvey(ctx, db.CreateSurveyParams{
		UserID:      ctx.UserId,
		Title:       params.Title,
		Description: db.NewNullString(params.Description),
		State:       resources.SurveyStatePending,
		Visibility:  resources.SurveyVisibilityPublic,
	})
	if err != nil {
		return nil, err
	}

	return survey, nil
}
