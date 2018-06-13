package employment

import "github.com/henryroyal/bls/series"

type MassLayoffStatistics struct {
	series.Dataset
}

func NewMassLayoffStatistics() (*MassLayoffStatistics) {
	return &MassLayoffStatistics{
		series.Dataset{
			Name:    "Mass Layoff Statistics",
			Symbol:  "ml",
			BaseURL: series.BaseURL,
		},
	}
}
