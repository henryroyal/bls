package employment

import (
	"github.com/henryroyal/bls/series"
	"regexp"
)

/*
	                      1         2
	             12345678901234567890
	Series ID    CEU0800000003
	Positions       Value           Field Name
	1-2             CE              Prefix
	3               U               Seasonal Adjustment Code
	4-11		08000000	Supersector and Industry Codes
	12-13           03              Data Type Code

 */

var (
	ceSeriesIdFormat *regexp.Regexp = regexp.MustCompile(
		`(?P<p>\w{2})(?P<sac>\w{1})(?P<sic>\w{8})(?P<dtc>\w{2})`,
	)
)

type NationalEmploymentHoursAndEarnings struct {
	series.Dataset
}

func NewNationalEmploymentHoursAndEarnings() (*NationalEmploymentHoursAndEarnings) {
	return &NationalEmploymentHoursAndEarnings{
		series.Dataset{
			Name:         "National Employment, Hours, and Earnings (SIC basis)",
			Symbol:       "ce",
			BaseURL:      series.BaseURL,
			SeriesFormat: *ceSeriesIdFormat,
			Datafiles:    []string{"ce.data.0.AllCESSeries"},
			Schema: `
				CREATE SCHEMA IF NOT EXISTS ce;
				CREATE TABLE IF NOT EXISTS ce.datatype (
	  				data_type_code smallint PRIMARY KEY,
	  				data_type_text varchar(250)
				);
				CREATE TABLE IF NOT EXISTS ce.footnote (
  					footnote_code char(1) PRIMARY KEY,
	  				footnote_text varchar(250)
				);
				CREATE TABLE IF NOT EXISTS ce.industry (
  					industry_code     int PRIMARY KEY,
	  				naics_code        int,
	  				publishing_status varchar(2),
	  				industry_name     varchar(250),
	  				display_level     char(1),
	  				selectable        char(1),
	  				sort_sequence     smallint
				);
				CREATE TABLE IF NOT EXISTS ce.period (
					period_code char(3) PRIMARY KEY,
  					period_abbr char(10),
  					period_name varchar(20)
				);
				CREATE TABLE IF NOT EXISTS ce.seasonal (
	  				industry_code char(1) PRIMARY KEY,
	  				seasonal_text varchar(25)
				);
				CREATE TABLE IF NOT EXISTS ce.supersector (
  					supersector_code smallint PRIMARY KEY,
	  				supersector_name varchar(50)
				);
				CREATE TABLE IF NOT EXISTS ce.series (
	  				series_id        char(13) PRIMARY KEY,
	  				supersector_code smallint REFERENCES ce.supersector (supersector_code),
	  				industry_code    int      REFERENCES ce.industry (industry_code),
	  				data_type_code   smallint REFERENCES ce.datatype (data_type_code),
	  				seasonal         char(1)  REFERENCES ce.seasonal (industry_code),
	  				series_title     varchar(250),
	  				footnote_codes   char(1)  REFERENCES ce.footnote (footnote_code),
	  				begin_year       int,
	  				begin_period     char(3)  REFERENCES ce.period (period_code),
	  				end_year         int,
	  				end_period       char(3)  REFERENCES ce.period (period_code)
				);
				CREATE TABLE IF NOT EXISTS ce.observations (
  					series_id      char(13) REFERENCES ce.series (series_id),
	  				year           int,
	  				period         char(3)  REFERENCES ce.period (period_code),
	  				value          int,
	  				footnote_codes char(1)  REFERENCES ce.footnote (footnote_code)
				);
			`,
		},
	}
}
