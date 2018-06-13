package prices

import "github.com/henryroyal/bls/series"

type LegacyConsumerPriceIndexAllUrban struct {
	series.Dataset
}

func NewLegacyConsumerPriceIndexAllUrban() (*LegacyConsumerPriceIndexAllUrban) {
	return &LegacyConsumerPriceIndexAllUrban{
		series.Dataset{
			Name:    "Consumer Price Index - All Urban Consumers (Old Series)",
			Symbol:  "mu",
			BaseURL: series.BaseURL,
		},
	}
}
