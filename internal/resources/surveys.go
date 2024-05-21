package resources

import "time"

type SurveyState string
type SurveyVisibility string

type Survey struct {
	Id            int64            `json:"id"`
	Title         string           `json:"title"`
	UserId        int64            `json:"user_id"`
	State         SurveyState      `json:"state"`
	Visibility    SurveyVisibility `json:"visibility"`
	Description   string           `json:"description,omitempty"`
	ResponseCount int64            `json:"response_count"`
	CreatedAt     time.Time        `json:"created_at"`
	UpdatedAt     time.Time        `json:"updated_at"`
}

const (
	SurveyStatePending          = SurveyState("pending")
	SurveyStateGatheringOptions = SurveyState("gathering_options")
	SurveyStateVoting           = SurveyState("voting")
	SurveyStateClosed           = SurveyState("closed")

	SurveyVisibilityPublic  = SurveyVisibility("public")
	SurveyVisibilityPrivate = SurveyVisibility("private")
	SurveyVisibilityLink    = SurveyVisibility("link")
)
