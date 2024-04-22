package surveys

import (
	"fmt"

	"github.com/carterjackson/ranked-pick-api/internal/common"
)

func List(ctx *common.Context) (interface{}, error) {
	surveys, err := ctx.Db.ListSurveys(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Println(surveys)

	return surveys, nil
}
