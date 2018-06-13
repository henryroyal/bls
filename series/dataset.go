package bls

// Dataset Interface
type DatasetFetcher interface {
	FetchDocumentation()
	FetchSeriesIndex()
	FetchDataset()
}

// Dataset Base Structure
type Dataset struct {
	Name   string
	Symbol string
	blsUrl string
}
