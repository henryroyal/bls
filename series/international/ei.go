package international

import "github.com/henryroyal/bls/series"

type ImportExportPriceIndexes struct {
	series.Dataset
}

func NewImportExportPriceIndexes() (*ImportExportPriceIndexes) {
	return &ImportExportPriceIndexes{
		series.Dataset{
			Name:    "Import/Export Price Indexes",
			Symbol:  "ei",
			BaseURL: series.BaseURL,
		},
	}
}
