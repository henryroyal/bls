package prices

import "github.com/henryroyal/bls/series"

type ChainedConsumerPriceIndexAllUrban struct {
	series.Dataset
}

func NewChainedConsumerPriceIndexAllUrban() (*ChainedConsumerPriceIndexAllUrban) {
	return &ChainedConsumerPriceIndexAllUrban{
		series.Dataset{
			Name:    "Chained CPI-All Urban Consumers",
			Symbol:  "su",
			BaseURL: series.BaseURL,
		},
	}
}
