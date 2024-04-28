package surveys

import (
	"github.com/carterjackson/ranked-pick-api/internal/common"
)

func List(ctx *common.Context) (interface{}, error) {
	surveys, err := ctx.Queries.ListSurveys(ctx)
	if err != nil {
		return nil, err
	}

	return surveys, nil
}
