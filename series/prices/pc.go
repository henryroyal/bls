package prices

import "github.com/henryroyal/bls/series"

type ProducerPriceIndexIndustryData struct {
	series.Dataset
}

func NewProducerPriceIndexIndustryData() (*ProducerPriceIndexIndustryData) {
	return &ProducerPriceIndexIndustryData{
		series.Dataset{
			Name:    "Producer Price Index Industry Data - Current Series",
			Symbol:  "pc",
			BaseURL: series.BaseURL,
		},
	}
}
