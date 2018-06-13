package occupations

import "github.com/henryroyal/bls/series"

type OccupationalRequirementsSurvey struct {
	series.Dataset
}

func NewOccupationalRequirementsSurvey() (*OccupationalRequirementsSurvey) {
	return &OccupationalRequirementsSurvey{
		series.Dataset{
			Name:    "Occupational Requirements Survey",
			Symbol:  "or",
			BaseURL: series.BaseURL,
		},
	}
}
