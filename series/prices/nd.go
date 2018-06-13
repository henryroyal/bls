package prices

import "github.com/henryroyal/bls/series"

type LegacyProducerPriceIndexIndustryDataNAICS struct {
	series.Dataset
}

func NewLegacyProducerPriceIndexIndustryDataNAICS() (*LegacyProducerPriceIndexIndustryDataNAICS) {
	return &LegacyProducerPriceIndexIndustryDataNAICS{
		series.Dataset{
			Name:    "Producer Price Index Industry Data - Discontinued Series (NAICS basis)",
			Symbol:  "nd",
			BaseURL: series.BaseURL,
		},
	}
}
