package surveys

import (
	"github.com/carterjackson/ranked-pick-api/internal/common"
	"github.com/carterjackson/ranked-pick-api/internal/db"
)

func Read(ctx *common.Context, id int64) (interface{}, error) {
	dbSurvey, err := ctx.Queries.ReadSurvey(ctx, id)
	if err != nil {
		return nil, err
	}
	survey := db.NewSurvey(&dbSurvey)

	count, err := ctx.Queries.CountSurveyResponsesForSurvey(ctx, survey.Id)
	if err != nil {
		return nil, err
	}
	survey.ResponseCount = count

	return survey, nil
}
