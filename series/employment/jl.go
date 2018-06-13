package employment

import "github.com/henryroyal/bls/series"

type JobOpeningsAndLaborTurnoverSIC struct {
	series.Dataset
}

func NewJobOpeningsAndLaborTurnoverSIC() (*JobOpeningsAndLaborTurnoverSIC) {
	return &JobOpeningsAndLaborTurnoverSIC{
		series.Dataset{
			Name:    "Job Openings and Labor Turnover Survey (SIC basis)",
			Symbol:  "jl",
			BaseURL: series.BaseURL,
		},
	}
}
