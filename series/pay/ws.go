package pay

import "github.com/henryroyal/bls/series"

type WorkStoppageData struct {
	series.Dataset
}

func NewWorkStoppageData() (*WorkStoppageData) {
	return &WorkStoppageData{
		series.Dataset{
			Name:    "Work Stoppage Data",
			Symbol:  "ws",
			BaseURL: series.BaseURL,
		},
	}
}
