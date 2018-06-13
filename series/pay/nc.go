package pay

import "github.com/henryroyal/bls/series"

type NationalCompensationSurvey struct {
	series.Dataset
}

func NewNationalCompensationSurvey() (*NationalCompensationSurvey) {
	return &NationalCompensationSurvey{
		series.Dataset{
			Name:    "National Compensation Survey",
			Symbol:  "nc",
			BaseURL: series.BaseURL,
		},
	}
}
