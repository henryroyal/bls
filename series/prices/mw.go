package prices

import "github.com/henryroyal/bls/series"

type LegacyConsumerPriceIndexUrbanWageWorkers struct {
	series.Dataset
}

func NewLegacyConsumerPriceIndexUrbanWageWorkers() (*LegacyConsumerPriceIndexUrbanWageWorkers) {
	return &LegacyConsumerPriceIndexUrbanWageWorkers{
		series.Dataset{
			Name:    "Consumer Price Index - Urban Wage Earners and Clerical Workers (Old Series)",
			Symbol:  "mw",
			BaseURL: series.BaseURL,
		},
	}
}
