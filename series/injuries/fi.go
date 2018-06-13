package injuries

import "github.com/henryroyal/bls/series"

type FatalOccupationalInjuries2003to2010 struct {
	series.Dataset
}

func NewFatalOccupationalInjuries2003to2010() (*FatalOccupationalInjuries2003to2010) {
	return &FatalOccupationalInjuries2003to2010{
		series.Dataset{
			Name:    "Census of Fatal Occupational Injuries (2003 - 2010)",
			Symbol:  "fi",
			BaseURL: series.BaseURL,
		},
	}
}
