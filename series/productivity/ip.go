package productivity

import "github.com/henryroyal/bls/series"

type IndustryProductivity struct {
	series.Dataset
}

func NewIndustryProductivity() (*IndustryProductivity) {
	return &IndustryProductivity{
		series.Dataset{
			Name:    "Industry Productivity",
			Symbol:  "ip",
			BaseURL: series.BaseURL,
		},
	}
}
