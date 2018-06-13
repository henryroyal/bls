package pay

import "github.com/henryroyal/bls/series"

type BenefitsCompensationSurvey struct {
	series.Dataset
}

func NewBenefitsCompensationSurvey() (*BenefitsCompensationSurvey) {
	return &BenefitsCompensationSurvey{
		series.Dataset{
			Name:    "National Compensation Survey - Benefits (Beginning with 2010 data)",
			Symbol:  "nb",
			BaseURL: series.BaseURL,
		},
	}
}
