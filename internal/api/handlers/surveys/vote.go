package surveys

import (
	"github.com/carterjackson/ranked-pick-api/internal/common"
	"github.com/carterjackson/ranked-pick-api/internal/db"
)

type VoteParams struct {
	Options []int64 `json:"options"`
}

func Vote(ctx *common.Context, tx *db.Queries, id int64, iparams interface{}) error {
	params := iparams.(*VoteParams)

	// TODO: Validate survey state

	for i, optionId := range params.Options {
		_, err := tx.ReadSurveyOption(ctx, optionId)
		if err != nil {
			return err
		}

		_, err = tx.UpsertSurveyAnswer(ctx, db.UpsertSurveyAnswerParams{
			SurveyID:       id,
			UserID:         ctx.UserId,
			SurveyOptionID: optionId,
			Rank:           int64(i),
		})
		if err != nil {
			return err
		}
	}

	return nil
}
