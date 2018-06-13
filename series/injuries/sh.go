package injuries

import "github.com/henryroyal/bls/series"

type NonfatalOccupationalInjuriesAndIllnesses1989to2001 struct {
	series.Dataset
}

func NewNonfatalOccupationalInjuriesAndIllnesses1989to2001() (*NonfatalOccupationalInjuriesAndIllnesses1989to2001) {
	return &NonfatalOccupationalInjuriesAndIllnesses1989to2001{
		series.Dataset{
			Name:    "Occupational injuries and illnesses: industry data (1989 - 2001)",
			Symbol:  "sh",
			BaseURL: series.BaseURL,
		},
	}
}
