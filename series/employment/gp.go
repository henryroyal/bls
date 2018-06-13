package employment

import "github.com/henryroyal/bls/series"

type GeographicProfile struct {
	series.Dataset
}

func NewGeographicProfile() (*GeographicProfile) {
	return &GeographicProfile{
		series.Dataset{
			Name:    "Geographic Profile",
			Symbol:  "gp",
			BaseURL: series.BaseURL,
		},
	}
}
