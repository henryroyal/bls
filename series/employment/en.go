package employment

import "github.com/henryroyal/bls/series"

type StateAndCountyEmploymentAndWagesQuarterly struct {
	series.Dataset
}

func NewStateAndCountyEmploymentAndWagesQuarterly() (*StateAndCountyEmploymentAndWagesQuarterly) {
	return &StateAndCountyEmploymentAndWagesQuarterly{
		series.Dataset{
			Name:    "State and County Employment and Wages from Quarterly Census of Employment and Wages",
			Symbol:  "en",
			BaseURL: series.BaseURL,
		},
	}
}
