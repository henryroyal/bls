package pay

import "github.com/henryroyal/bls/series"

type EmploymentCostIndexSIC struct {
	series.Dataset
}

func NewEmploymentCostIndexSIC() (*EmploymentCostIndexSIC) {
	return &EmploymentCostIndexSIC{
		series.Dataset{
			Name:    "Employment Cost Index (SIC)",
			Symbol:  "ec",
			BaseURL: series.BaseURL,
		},
	}
}
