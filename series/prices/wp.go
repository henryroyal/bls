package prices

import "github.com/henryroyal/bls/series"

type ProducerPriceIndexCommodityData struct {
	series.Dataset
}

func NewProducerPriceIndexCommodityData() (*ProducerPriceIndexCommodityData) {
	return &ProducerPriceIndexCommodityData{
		series.Dataset{
			Name:    "Producer Price Index Commodity Data - Current Series",
			Symbol:  "wp",
			BaseURL: series.BaseURL,
		},
	}
}
