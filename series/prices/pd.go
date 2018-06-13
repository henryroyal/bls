package prices

import "github.com/henryroyal/bls/series"

type LegacyProducerPriceIndexIndustryDataSIC struct {
	series.Dataset
}

func NewLegacyProducerPriceIndexIndustryDataSIC() (*LegacyProducerPriceIndexIndustryDataSIC) {
	return &LegacyProducerPriceIndexIndustryDataSIC{
		series.Dataset{
			Name:    "Producer Price Index Industry Data - Discontinued Series (SIC basis)",
			Symbol:  "pd",
			BaseURL: series.BaseURL,
		},
	}
}
