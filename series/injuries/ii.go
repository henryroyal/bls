package injuries

import "github.com/henryroyal/bls/series"

type NonfatalOccupationalInjuriesAndIllnesses2003to2013 struct {
	series.Dataset
}

func NewNonfatalOccupationalInjuriesAndIllnesses2003to2013() (*NonfatalOccupationalInjuriesAndIllnesses2003to2013) {
	return &NonfatalOccupationalInjuriesAndIllnesses2003to2013{
		series.Dataset{
			Name:    "Occupational Injuries and Illnesses - industry data (2003 - 2013)",
			Symbol:  "ii",
			BaseURL: series.BaseURL,
		},
	}
}
