package injuries

import "github.com/henryroyal/bls/series"

type NonfatalOccupationalInjuriesAndIllnessesPre1989 struct {
	series.Dataset
}

func NewNonfatalOccupationalInjuriesAndIllnessesPre1989() (*NonfatalOccupationalInjuriesAndIllnessesPre1989) {
	return &NonfatalOccupationalInjuriesAndIllnessesPre1989{
		series.Dataset{
			Name:    "Occupational injuries and illnesses: industry data (pre-1989)",
			Symbol:  "hs",
			BaseURL: series.BaseURL,
		},
	}
}
