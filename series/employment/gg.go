package employment

import "github.com/henryroyal/bls/series"

type GreenGoodsAndServices struct {
	series.Dataset // TODO - make sure this is supposed to be in employment package
}

func NewGreenGoodsAndServices() (*GreenGoodsAndServices) {
	return &GreenGoodsAndServices{
		series.Dataset{
			Name:    "Green Goods and Services",
			Symbol:  "gg",
			BaseURL: series.BaseURL,
		},
	}
}
