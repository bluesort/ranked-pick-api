package surveys

import (
	"database/sql"

	"github.com/carterjackson/ranked-pick-api/internal/common"
	"github.com/carterjackson/ranked-pick-api/internal/db"
	"github.com/carterjackson/ranked-pick-api/internal/errors"
)

func Delete(ctx *common.Context, tx *db.Queries, id int64) error {
	survey, err := ctx.Queries.ReadSurvey(ctx, id)
	if err == sql.ErrNoRows {
		return errors.NewNotFoundError()
	} else if err != nil {
		return err
	}

	if survey.UserID != ctx.UserId {
		return errors.NewAuthError()
	}

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

	return nil
}
