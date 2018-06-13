package injuries

import "github.com/henryroyal/bls/series"

type NonfatalOccupationalInjuries1992to2001 struct {
	series.Dataset
}

func NewNonfatalOccupationalInjuries1992to2001() (*NonfatalOccupationalInjuries1992to2001) {
	return &NonfatalOccupationalInjuries1992to2001{
		series.Dataset{
			Name:    "Nonfatal cases involving days away from work: selected characteristics (1992 - 2001)",
			Symbol:  "cd",
			BaseURL: series.BaseURL,
		},
	}
}
