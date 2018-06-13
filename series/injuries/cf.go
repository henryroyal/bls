package injuries

import "github.com/henryroyal/bls/series"

type FatalOccupationalInjuries1992to2002 struct {
	series.Dataset
}

func NewFatalOccupationalInjuries1992to2002() (*FatalOccupationalInjuries1992to2002) {
	return &FatalOccupationalInjuries1992to2002{
		series.Dataset{
			Name:    "Census of Fatal Occupational Injuries (1992 - 2002)",
			Symbol:  "cf",
			BaseURL: series.BaseURL,
		},
	}
}
