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

	var optionsResp []*resources.SurveyOption
	for _, option := range options {
		optionsResp = append(optionsResp, newSurveyOptionResp(&option))
	}

	return optionsResp, nil
}
