package users

import (
	"github.com/carterjackson/ranked-pick-api/internal/common"
	"github.com/carterjackson/ranked-pick-api/internal/db"
	"github.com/carterjackson/ranked-pick-api/internal/errors"
	"github.com/carterjackson/ranked-pick-api/internal/resources"
)

func ListRespondedSurveys(ctx *common.Context, id int64) (interface{}, error) {
	if ctx.UserId != id {
		return nil, errors.NewForbiddenError()
	}

	surveys, err := ctx.Queries.ListSurveysForUserResponded(ctx, db.NewNullInt64(id))
	if err != nil {
		return nil, err
	}

	var surveysResp []*resources.Survey
	for _, survey := range surveys {
		surveyRow := db.SurveyRow(survey)
		surveysResp = append(surveysResp, db.NewSurveyFromRow(&surveyRow))
	}

	return surveysResp, nil
}
