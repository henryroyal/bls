package injuries

import "github.com/henryroyal/bls/series"

type NonfatalOccupationalInjuriesAndIllnesses2002 struct {
	series.Dataset
}

func NewNonfatalOccupationalInjuriesAndIllnesses2002() (*NonfatalOccupationalInjuriesAndIllnesses2002) {
	return &NonfatalOccupationalInjuriesAndIllnesses2002{
		series.Dataset{
			Name:    "Occupational injuries and illnesses: industry data (2002)",
			Symbol:  "si",
			BaseURL: series.BaseURL,
		},
	}
}
