package injuries

import "github.com/henryroyal/bls/series"

type OccupationalInjuriesAndIllnessesPost2014 struct {
	series.Dataset
}

func NewOccupationalInjuriesAndIllnessesPost2014() (*OccupationalInjuriesAndIllnessesPost2014) {
	return &OccupationalInjuriesAndIllnessesPost2014{
		series.Dataset{
			Name:    "Occupational Injuries and Illnesses - industry data (2014 forward)",
			Symbol:  "is",
			BaseURL: series.BaseURL,
		},
	}
}
