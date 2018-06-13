package employment

import "github.com/henryroyal/bls/series"

type StateAndAreaEmploymentHoursAndEarningsSIC struct {
	series.Dataset
}

func NewStateAndAreaEmploymentHoursAndEarningsSIC() (*StateAndAreaEmploymentHoursAndEarningsSIC) {
	return &StateAndAreaEmploymentHoursAndEarningsSIC{
		series.Dataset{
			Name:    "State and Area Employment, Hours, and Earnings (SIC basis)",
			Symbol:  "sa",
			BaseURL: series.BaseURL,
		},
	}
}
