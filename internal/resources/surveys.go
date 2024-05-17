package resources

type SurveyState string
type SurveyVisibility string

type Survey struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	UserId      int64  `json:"user_id"`
	State       SurveyState
	Visibility  SurveyVisibility
	Description string `json:"description,omitempty"`
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
