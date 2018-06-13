package pay

import "github.com/henryroyal/bls/series"

type EmployeeBenefitsSurvey struct {
	series.Dataset
}

func NewEmployeeBenefitsSurvey() (*EmployeeBenefitsSurvey) {
	return &EmployeeBenefitsSurvey{
		series.Dataset{
			Name:    "Employee Benefits Survey",
			Symbol:  "eb",
			BaseURL: series.BaseURL,
		},
	}
}
