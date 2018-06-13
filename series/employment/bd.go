package employment

import "github.com/henryroyal/bls/series"

type BusinessEmploymentDynamics struct {
	series.Dataset
}

func NewBusinessEmploymentDynamics() (*BusinessEmploymentDynamics) {
	return &BusinessEmploymentDynamics{
		series.Dataset{
			Name:    "Business Employment Dynamics",
			Symbol:  "bd",
			BaseURL: series.BaseURL,
		},
	}
}
