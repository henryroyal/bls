package series

import (
	"fmt"
	"path"
)

const BaseURL string = "https://download.bls.gov/pub/time.series/"

type DatasetFetcher interface {
	DocumentationURL()
	SeriesIndexURL()
}

type Dataset struct {
	Name    string
	Symbol  string
	BaseURL string
}

func (self *Dataset) DocumentationURL() string {
	return path.Join(BaseURL, self.Symbol, fmt.Sprintf("%v.txt", self.Symbol))
}

func (self *Dataset) SeriesIndexURL() string {
	return path.Join(BaseURL, self.Symbol, fmt.Sprintf("%v.series", self.Symbol))
}

func (self *Dataset) String() string {
	return self.Name
}
