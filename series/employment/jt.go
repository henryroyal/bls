package employment

import "github.com/henryroyal/bls/series"

type JobOpeningsAndLaborTurnover struct {
	series.Dataset
}

func NewJobOpeningsAndLaborTurnover() (*JobOpeningsAndLaborTurnover) {
	return &JobOpeningsAndLaborTurnover{
		series.Dataset{
			Name:    "Job Openings and Labor Turnover Survey",
			Symbol:  "jt",
			BaseURL: series.BaseURL,
		},
	}
}
