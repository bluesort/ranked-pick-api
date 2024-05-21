package users

import (
	"github.com/carterjackson/ranked-pick-api/internal/api/handlers/surveys"
	"github.com/carterjackson/ranked-pick-api/internal/common"
	"github.com/carterjackson/ranked-pick-api/internal/db"
	"github.com/carterjackson/ranked-pick-api/internal/errors"
)

func Delete(ctx *common.Context, tx *db.Queries, id int64) error {
	if ctx.UserId != id {
		return errors.NewForbiddenError()
	}

	dbSurveys, err := tx.ListSurveysForUser(ctx, id)
	if err != nil {
		return err
	}
	for _, survey := range dbSurveys {
		err = surveys.Delete(ctx, tx, survey.ID)
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
