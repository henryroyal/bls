package spending

import "github.com/henryroyal/bls/series"

type AmericanTimeUseSurvey struct {
	series.Dataset
}

func NewAmericanTimeUseSurvey() (*AmericanTimeUseSurvey) {
	return &AmericanTimeUseSurvey{
		series.Dataset{
			Name:    "American Time Use Survey",
			Symbol:  "tu",
			BaseURL: series.BaseURL,
		},
	}
}
