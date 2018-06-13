package injuries

import "github.com/henryroyal/bls/series"

type NonfatalOccupationalInjuriesPost2011 struct {
	series.Dataset
}

func NewNonfatalOccupationalInjuriesPost2011() (*NonfatalOccupationalInjuriesPost2011) {
	return &NonfatalOccupationalInjuriesPost2011{
		series.Dataset{
			Name:    "Nonfatal cases involving days away from work: selected characteristics (2011 forward)",
			Symbol:  "cs",
			BaseURL: series.BaseURL,
		},
	}
}
