package productivity

import "github.com/henryroyal/bls/series"

type MajorSectorMultifactorProductivity struct {
	series.Dataset
}

func NewMajorSectorMultifactorProductivity() (*MajorSectorMultifactorProductivity) {
	return &MajorSectorMultifactorProductivity{
		series.Dataset{
			Name:    "Major Sector Multifactor Productivity",
			Symbol:  "mp",
			BaseURL: series.BaseURL,
		},
	}
}
