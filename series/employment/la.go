package employment

import "github.com/henryroyal/bls/series"

type LocalAreaUnemploymentStatistics struct {
	series.Dataset
}

func NewLocalAreaUnemploymentStatistics() (*LocalAreaUnemploymentStatistics) {
	return &LocalAreaUnemploymentStatistics{
		series.Dataset{
			Name:    "Local Area Unemployment Statistics",
			Symbol:  "la",
			BaseURL: series.BaseURL,
		},
	}
}
