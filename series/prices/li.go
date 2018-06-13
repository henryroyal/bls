package prices

import "github.com/henryroyal/bls/series"

type DepartmentStoreInventoryPriceIndex struct {
	series.Dataset
}

func NewDepartmentStoreInventoryPriceIndex() (*DepartmentStoreInventoryPriceIndex) {
	return &DepartmentStoreInventoryPriceIndex{
		series.Dataset{
			Name:    "Department Store Inventory Price Index",
			Symbol:  "li",
			BaseURL: series.BaseURL,
		},
	}
}
