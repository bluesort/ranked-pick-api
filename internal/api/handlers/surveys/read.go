package surveys

import (
	"database/sql"

	"github.com/carterjackson/ranked-pick-api/internal/common"
	"github.com/carterjackson/ranked-pick-api/internal/db"
	"github.com/carterjackson/ranked-pick-api/internal/errors"
)

func Read(ctx *common.Context, id int64) (interface{}, error) {
	dbSurvey, err := ctx.Queries.ReadSurvey(ctx, id)
	if err == sql.ErrNoRows {
		return nil, errors.NewNotFoundError()
	} else if err != nil {
		return nil, err
	}
	surveyRow := db.SurveyRow(dbSurvey)
	survey := db.NewSurveyFromRow(&surveyRow)

	count, err := ctx.Queries.CountSurveyResponsesForSurvey(ctx, survey.Id)
	if err != nil {
		return nil, err
	}
	survey.ResponseCount = count

	return survey, nil
}
