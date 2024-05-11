package surveys

import (
	"github.com/carterjackson/ranked-pick-api/internal/db"
	"github.com/carterjackson/ranked-pick-api/internal/resources"
)

func newSurveyResp(survey *db.Survey) *resources.Survey {
	return &resources.Survey{
		Id:          survey.ID,
		Title:       survey.Title,
		Description: survey.Description.String,
	}
}
