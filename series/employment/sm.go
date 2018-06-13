package employment

import "github.com/henryroyal/bls/series"

type StateAndAreaEmploymentHoursAndEarnings struct {
	series.Dataset
}

func NewStateAndAreaEmploymentHoursAndEarnings() (*StateAndAreaEmploymentHoursAndEarnings) {
	return &StateAndAreaEmploymentHoursAndEarnings{
		series.Dataset{
			Name:    "State and Area Employment, Hours, and Earnings",
			Symbol:  "sm",
			BaseURL: series.BaseURL,
		},
	}
}
