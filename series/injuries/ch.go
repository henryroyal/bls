package injuries

import "github.com/henryroyal/bls/series"

type NonfatalOccupationalInjuries2003to2010 struct {
	series.Dataset
}

func NewNonfatalOccupationalInjuries2003to2010() (*NonfatalOccupationalInjuries2003to2010) {
	return &NonfatalOccupationalInjuries2003to2010{
		series.Dataset{
			Name:    "Nonfatal cases involving days away from work: selected characteristics (2003 - 2010)",
			Symbol:  "ch",
			BaseURL: series.BaseURL,
		},
	}
}
