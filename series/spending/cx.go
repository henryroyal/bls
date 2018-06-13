package spending

import "github.com/henryroyal/bls/series"

type ConsumerExpenditureSurvey struct {
	series.Dataset
}

func NewConsumerExpenditureSurvey() (*ConsumerExpenditureSurvey) {
	return &ConsumerExpenditureSurvey{
		series.Dataset{
			Name:    "Consumer Expenditure Survey",
			Symbol:  "cx",
			BaseURL: series.BaseURL,
		},
	}
}
