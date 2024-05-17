package surveys

import (
	"github.com/carterjackson/ranked-pick-api/internal/common"
	"github.com/carterjackson/ranked-pick-api/internal/db"
)

func Read(ctx *common.Context, id int64) (interface{}, error) {
	survey, err := ctx.Queries.ReadSurvey(ctx, id)
	if err != nil {
		return nil, err
	}

	return db.NewSurvey(&survey), nil
}
