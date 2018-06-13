package employment

import "github.com/henryroyal/bls/series"

type OccupationalEmploymentStatistics struct {
	series.Dataset
}

func NewOccupationalEmploymentStatistics() (*OccupationalEmploymentStatistics) {
	return &OccupationalEmploymentStatistics{
		series.Dataset{
			Name:    "Occupational Employment Statistics (OES)",
			Symbol:  "oe",
			BaseURL: series.BaseURL,
		},
	}
}
