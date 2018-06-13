package prices

import "github.com/henryroyal/bls/series"

type ConsumerPriceIndexAllUrban struct {
	series.Dataset
}

func NewConsumerPriceIndexAllUrban() (*ConsumerPriceIndexAllUrban) {
	return &ConsumerPriceIndexAllUrban{
		series.Dataset{
			Name:    "Consumer Price Index - All Urban Consumers",
			Symbol:  "cu",
			BaseURL: series.BaseURL,
		},
	}
}
