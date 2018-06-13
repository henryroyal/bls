package pay

import "github.com/henryroyal/bls/series"

type EmployerCostForEmployeeCompensationQuarterly struct {
	series.Dataset
}

func NewEmployerCostForEmployeeCompensationQuarterly() (*EmployerCostForEmployeeCompensationQuarterly) {
	return &EmployerCostForEmployeeCompensationQuarterly{
		series.Dataset{
			Name:    "Employer Costs for Employee Compensation by Quarter",
			Symbol:  "ci",
			BaseURL: series.BaseURL,
		},
	}
}
