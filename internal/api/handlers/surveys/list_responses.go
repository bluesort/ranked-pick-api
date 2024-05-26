package surveys

import (
	"github.com/carterjackson/ranked-pick-api/internal/common"
	"github.com/carterjackson/ranked-pick-api/internal/db"
)

type ListResponsesParams struct {
	UserId int64 `json:"user_id,string"`
}

func ListResponses(ctx *common.Context, id int64, iparams interface{}) (interface{}, error) {
	params := iparams.(*ListResponsesParams)

	var err error
	var responses []db.SurveyResponse
	if params.UserId != 0 {
		responses, err = ctx.Queries.ListSurveyResponsesForSurveyUser(ctx, db.ListSurveyResponsesForSurveyUserParams{
			SurveyID: id,
			UserID:   db.NewNullInt64(ctx.UserId),
		})
		if err != nil {
			return nil, err
		}
	} else {
		responses, err = ctx.Queries.ListSurveyResponsesForSurvey(ctx, id)
		if err != nil {
			return nil, err
		}
	}

	return responses, nil
}
