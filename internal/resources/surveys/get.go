package surveys

import "github.com/carterjackson/ranked-pick-api/internal/common"

func Get(ctx *common.Context, id int64) (interface{}, error) {
	survey, err := ctx.Queries.GetSurvey(ctx, id)
	if err != nil {
		return nil, err
	}

	return survey, nil
}
