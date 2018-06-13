package prices

import "github.com/henryroyal/bls/series"

type AveragePriceData struct {
	series.Dataset
}

func NewAveragePriceData() (*AveragePriceData) {
	return &AveragePriceData{
		series.Dataset{
			Name:    "Average Prices",
			Symbol:  "ap",
			BaseURL: series.BaseURL,
		},
	}
}
