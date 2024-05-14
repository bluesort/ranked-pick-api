package surveys

import (
	"github.com/carterjackson/ranked-pick-api/internal/common"
	"github.com/carterjackson/ranked-pick-api/internal/db"
	"github.com/carterjackson/ranked-pick-api/internal/resources"
)

func ListOptions(ctx *common.Context, tx *db.Queries, id int64) (interface{}, error) {
	options, err := tx.ListSurveyOptionsForSurvey(ctx, id)
	if err != nil {
		return nil, err
	}

	// TODO: Use previous vote order if user has voted

	optionsResp := make([]*resources.SurveyOption, len(options))
	for i, option := range options {
		optionsResp[i] = newSurveyOptionResp(&option)
	}

	return optionsResp, nil
}
