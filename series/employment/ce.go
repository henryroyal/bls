package employment

import "github.com/henryroyal/bls/series"

type NationalEmploymentHoursAndEarnings struct {
	series.Dataset
}

func NewNationalEmploymentHoursAndEarnings() (*NationalEmploymentHoursAndEarnings) {
	return &NationalEmploymentHoursAndEarnings{
		series.Dataset{
			Name:    "National Employment, Hours, and Earnings (SIC basis)",
			Symbol:  "ce",
			BaseURL: series.BaseURL,
		},
	}
}
