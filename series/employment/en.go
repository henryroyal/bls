package employment

import "github.com/henryroyal/bls/series"

// FIXME - this set is missing, figure out where it went or remove it
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
