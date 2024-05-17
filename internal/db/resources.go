package db

import "github.com/carterjackson/ranked-pick-api/internal/resources"

//
// Helpers for converting db objects to resources
//

func NewSurvey(survey *Survey) *resources.Survey {
	return &resources.Survey{
		Id:          survey.ID,
		Title:       survey.Title,
		UserId:      survey.UserID,
		State:       resources.SurveyState(survey.State),
		Visibility:  resources.SurveyVisibility(survey.Visibility),
		Description: survey.Description.String,
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
		Email:       user.Email,
		DisplayName: user.DisplayName.String,
	}
}
