package employment

import "github.com/henryroyal/bls/series"

type NationalEmploymentHoursAndEarningsSIC struct {
	series.Dataset
}

func NewNationalEmploymentHoursAndEarningsSIC() (*NationalEmploymentHoursAndEarningsSIC) {
	return &NationalEmploymentHoursAndEarningsSIC{
		series.Dataset{
			Name:    "National Employment, Hours, and Earnings (SIC basis)",
			Symbol:  "ee",
			BaseURL: series.BaseURL,
		},
	}
}
