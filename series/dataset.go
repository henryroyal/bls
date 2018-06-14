package series

import (
	"fmt"
)

const BaseURL string = `https://download.bls.gov/pub/time.series/`

type DatasetFetcher interface {
	DocumentationURL()
	SeriesIndexURL()
}

type Dataset struct {
	Name    string
	Symbol  string
	BaseURL string
}

func (self *Dataset) DatasetURL() string {
	return (self.BaseURL + self.Symbol)
}

func (self *Dataset) DocumentationURL() string {
	return (self.BaseURL + self.Symbol + fmt.Sprintf("%v.txt", self.Symbol))
}

func (self *Dataset) SeriesIndexURL() string {
	return (self.BaseURL + self.Symbol + fmt.Sprintf("%v.series", self.Symbol))
}
