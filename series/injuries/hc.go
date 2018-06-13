package injuries

import "github.com/henryroyal/bls/series"

type NonfatalOccupationalInjuries2002 struct {
	series.Dataset
}

func NewNonfatalOccupationalInjuries2002() (*NonfatalOccupationalInjuries2002) {
	return &NonfatalOccupationalInjuries2002{
		series.Dataset{
			Name:    "Nonfatal cases involving days away from work: selected characteristics",
			Symbol:  "hc",
			BaseURL: series.BaseURL,
		},
	}
}
