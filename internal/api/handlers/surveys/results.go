package surveys

import (
	"slices"

	"github.com/carterjackson/ranked-pick-api/internal/common"
	"github.com/carterjackson/ranked-pick-api/internal/db"
)

type OptionResult struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
	Rank  int64  `json:"rank"`
}

type Result struct {
	VoteCount     int             `json:"vote_count"`
	OptionResults []*OptionResult `json:"option_results"`
}

func Results(ctx *common.Context, tx *db.Queries, id int64) (interface{}, error) {
	options, err := tx.ListSurveyOptionsForSurvey(ctx, id)
	if err != nil {
		return nil, err
	}

	answers, err := tx.ListSurveyAnswersForSurvey(ctx, id)
	if err != nil {
		return nil, err
	}

	optionToResult := make(map[int64]*OptionResult, len(options))
	for _, option := range options {
		optionToResult[option.ID] = &OptionResult{
			Id:    option.ID,
			Title: option.Title,
			Rank:  0,
		}
	}
	for _, answer := range answers {
		optionToResult[answer.SurveyOptionID].Rank += answer.Rank
	}

	var optionResults []*OptionResult
	for _, result := range optionToResult {
		optionResults = append(optionResults, result)
	}
	slices.SortFunc(optionResults, func(a *OptionResult, b *OptionResult) int {
		return int(a.Rank) - int(b.Rank)
	})

	return &Result{
		VoteCount:     len(answers),
		OptionResults: optionResults,
	}, nil
}
