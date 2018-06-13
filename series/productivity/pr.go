package productivity

import "github.com/henryroyal/bls/series"

type MajorSectorProductivityAndCosts struct {
	series.Dataset
}

func NewMajorSectorProductivityAndCosts() (*MajorSectorProductivityAndCosts) {
	return &MajorSectorProductivityAndCosts{
		series.Dataset{
			Name:    "Major Sector Productivity and Costs",
			Symbol:  "pr",
			BaseURL: series.BaseURL,
		},
	}
}
