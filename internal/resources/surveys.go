package resources

type Survey struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
}
