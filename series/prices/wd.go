package prices

import "github.com/henryroyal/bls/series"

type LegacyProducerPriceIndexCommodityData struct {
	series.Dataset
}

func NewLegacyProducerPriceIndexCommodityData() (*LegacyProducerPriceIndexCommodityData) {
	return &LegacyProducerPriceIndexCommodityData{
		series.Dataset{
			Name:    "Producer Price Index Commodity Data - Discontinued Series",
			Symbol:  "wd",
			BaseURL: series.BaseURL,
		},
	}
}
