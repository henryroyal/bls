package pay

import "github.com/henryroyal/bls/series"

type EmployerCostForEmployeeCompensationSIC struct {
	series.Dataset
}

func NewEmployerCostForEmployeeCompensationSIC() (*EmployerCostForEmployeeCompensationSIC) {
	return &EmployerCostForEmployeeCompensationSIC{
		series.Dataset{
			Name:    "Employer Costs for Employee Compensation",
			Symbol:  "cc",
			BaseURL: series.BaseURL,
		},
	}
}
