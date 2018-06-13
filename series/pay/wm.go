package pay

import "github.com/henryroyal/bls/series"

type ModeledWageEstimates struct {
	series.Dataset
}

func NewModeledWageEstimates() (*ModeledWageEstimates) {
	return &ModeledWageEstimates{
		series.Dataset{
			Name:    "Modeled Wage Estimates",
			Symbol:  "wm",
			BaseURL: series.BaseURL,
		},
	}
}
