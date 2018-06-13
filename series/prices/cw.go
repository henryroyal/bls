package prices

import "github.com/henryroyal/bls/series"

type ConsumerPriceIndexUrbanWageWorkers struct {
	series.Dataset
}

func NewConsumerPriceIndexUrbanWageWorkers() (*ConsumerPriceIndexUrbanWageWorkers) {
	return &ConsumerPriceIndexUrbanWageWorkers{
		series.Dataset{
			Name:    "Consumer Price Index - Urban Wage Earners and Clerical Workers",
			Symbol:  "cw",
			BaseURL: series.BaseURL,
		},
	}
}
