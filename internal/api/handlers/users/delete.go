package users

import (
	"github.com/carterjackson/ranked-pick-api/internal/common"
	"github.com/carterjackson/ranked-pick-api/internal/db"
	"github.com/carterjackson/ranked-pick-api/internal/errors"
)

func Delete(ctx *common.Context, tx *db.Queries, id int64) error {
	if ctx.UserId != id {
		return errors.NewAuthError()
	}

	surveys, err := tx.ListSurveysForUser(ctx, id)
	if err != nil {
		return err
	}

	for _, survey := range surveys {
		err = tx.DeleteSurveyResponsesForSurvey(ctx, survey.ID)
		if err != nil {
			return err
		}
		err = tx.DeleteSurveyOptionsForSurvey(ctx, survey.ID)
		if err != nil {
			return err
		}
		err = tx.DeleteSurvey(ctx, survey.ID)
		if err != nil {
			return err
		}
	}

	err = tx.AnonymizeSurveyResponsesForUser(ctx, db.NewNullInt64(id))
	if err != nil {
		return err
	}

	err = tx.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
