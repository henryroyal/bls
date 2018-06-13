package injuries

import "github.com/henryroyal/bls/series"

type FatalOccupationalInjuriesPost2011 struct {
	series.Dataset
}

func NewFatalOccupationalInjuriesPost2011() (*FatalOccupationalInjuriesPost2011) {
	return &FatalOccupationalInjuriesPost2011{
		series.Dataset{
			Name:    "Census of Fatal Occupational Injuries (2011 forward)",
			Symbol:  "fw",
			BaseURL: series.BaseURL,
		},
	}
}
