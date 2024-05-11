package surveys

import (
	"github.com/carterjackson/ranked-pick-api/internal/common"
	"github.com/carterjackson/ranked-pick-api/internal/db"
)

type UpdateParams struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func Update(ctx *common.Context, tx *db.Queries, id int64, iparams interface{}) (interface{}, error) {
	params := iparams.(*UpdateParams)

	survey, err := tx.UpdateSurvey(ctx, db.UpdateSurveyParams{
		ID:          id,
		Title:       db.NewNullString(params.Title),
		Description: db.NewNullString(params.Description),
	})
	if err != nil {
		return nil, err
	}

	return newSurveyResp(&survey), nil
}
