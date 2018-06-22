package series

import (
	"regexp"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"fmt"
	"path"
)

const BaseURL string = `https://download.bls.gov/pub/time.series/`

//TODO: way to map arbitrary 'ignore' offsets (pop off first two rows or ignore fifth column)
type DatasetDescriptor interface {
	String() string
	DatasetURL() string
	DatasetFiles() ([]string, error)
}

type Dataset struct {
	Name         string
	Symbol       string
	BaseURL      string
	Schema       string
	Datafiles    []string
	SeriesFormat regexp.Regexp
}

func (self *Dataset) getHTMLDocument(url string) (*goquery.Document, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf(
			"getHTMLDocument: http response code error: %d %s",
			resp.StatusCode,
			resp.Status,
		)
	}

	index, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	return index, nil
}

func (self *Dataset) DatasetFiles() ([]string, error) {

	var files []string

	index, err := self.getHTMLDocument(self.DatasetURL())
	if err != nil {
		return nil, err
	}

	index.Find("html body pre a").Each(func(_ int, li *goquery.Selection) {
		if val, ok := li.Attr("href"); ok {
			_, file := path.Split(val)
			files = append(files, file)
		}
	})
	return files, nil
}

func (self *Dataset) DatasetURL() string {
	return self.BaseURL + self.Symbol + "/"
}

func (self *Dataset) String() string {
	return self.Symbol
}
