package db

import (
	"database/sql"
	"time"

	"github.com/carterjackson/ranked-pick-api/internal/resources"
)

//
// Helpers for converting db objects to resources
//

type SurveyRow struct {
	ID            int64
	UserID        int64
	Title         string
	State         string
	Visibility    string
	Description   sql.NullString
	CreatedAt     time.Time
	UpdatedAt     time.Time
	ResponseCount int64
}

func NewSurvey(survey *Survey) *resources.Survey {
	return &resources.Survey{
		Id:          survey.ID,
		Title:       survey.Title,
		UserId:      survey.UserID,
		State:       resources.SurveyState(survey.State),
		Visibility:  resources.SurveyVisibility(survey.Visibility),
		Description: survey.Description.String,
		CreatedAt:   survey.CreatedAt,
		UpdatedAt:   survey.UpdatedAt,
	}
}

func NewSurveyFromRow(survey *SurveyRow) *resources.Survey {
	return &resources.Survey{
		Id:            survey.ID,
		Title:         survey.Title,
		UserId:        survey.UserID,
		State:         resources.SurveyState(survey.State),
		Visibility:    resources.SurveyVisibility(survey.Visibility),
		Description:   survey.Description.String,
		ResponseCount: survey.ResponseCount,
		CreatedAt:     survey.CreatedAt,
		UpdatedAt:     survey.UpdatedAt,
	}
}

func NewSurveyOption(option *SurveyOption) *resources.SurveyOption {
	return &resources.SurveyOption{
		Id:    option.ID,
		Title: option.Title,
	}
}

func NewUser(user *User) *resources.User {
	return &resources.User{
		Id:          user.ID,
		Username:    user.Username,
		DisplayName: user.DisplayName.String,
	}
}
