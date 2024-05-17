package surveys

import (
	"github.com/carterjackson/ranked-pick-api/internal/common"
	"github.com/carterjackson/ranked-pick-api/internal/db"
	"github.com/carterjackson/ranked-pick-api/internal/resources"
)

type CreateParams struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Options     []string `json:"options"`
}

func Create(ctx *common.Context, tx *db.Queries, iparams interface{}) (interface{}, error) {
	params := iparams.(*CreateParams)

	// TODO: Make survey visibility private by default once invites are implemented
	survey, err := tx.CreateSurvey(ctx, db.CreateSurveyParams{
		UserID:      ctx.UserId,
		Title:       params.Title,
		Description: db.NewNullString(params.Description),
		State:       string(resources.SurveyStatePending),
		Visibility:  string(resources.SurveyVisibilityPublic),
	})
	if err != nil {
		return nil, err
	}

	for _, option := range params.Options {
		_, err = tx.CreateSurveyOption(ctx, db.CreateSurveyOptionParams{
			SurveyID: survey.ID,
			Title:    option,
		})
		if err != nil {
			return nil, err
		}
	}

	return db.NewSurvey(&survey), nil
}
