package surveys

import (
	"github.com/carterjackson/ranked-pick-api/internal/common"
	"github.com/carterjackson/ranked-pick-api/internal/resources"
)

func List(ctx *common.Context) (interface{}, error) {
	surveys, err := ctx.Queries.ListSurveys(ctx)
	if err != nil {
		return nil, err
	}

	var surveysResp []*resources.Survey
	for _, survey := range surveys {
		surveysResp = append(surveysResp, newSurveyResp(&survey))
	}

	return surveysResp, nil
}
