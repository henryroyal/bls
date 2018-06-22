package employment

import (
	"github.com/henryroyal/bls/series"
	"regexp"
)

/*
TODO: ee.series file is messed up, with the fifth field being junk
TODO: there will definitely be duplicates if you try to load all these data files
 */

/*
	                      1         2
	             12345678901234567890
	Series ID    EES10140001
	Positions       Value           Field Name
	1-2             EE              Prefix
	3               S               Seasonal Adjustment Code
	4-9             101400          Industry Code
	10-11           01              Data Type Code
 */

var (
	eeSeriesIdFormat = regexp.MustCompile(
		`(?P<p>\w{2})(?P<sac>\w{1})(?P<ic>\d{6})(?P<dtc>\w{2})`,
	)
)

type NationalEmploymentHoursAndEarningsSIC struct {
	series.Dataset
}

func NewNationalEmploymentHoursAndEarningsSIC() (*NationalEmploymentHoursAndEarningsSIC) {
	return &NationalEmploymentHoursAndEarningsSIC{
		series.Dataset{
			Name:         "National Employment, Hours, and Earnings (SIC basis)",
			Symbol:       "ee",
			BaseURL:      series.BaseURL,
			SeriesFormat: *eeSeriesIdFormat,
			Datafiles: []string{
				"ee.data.1.CurrentSeasAE",
				"ee.data.2.CurrentSeasPWWW",
				"ee.data.3.CurrentSeasOther",
				"ee.data.4.HistorySeasAE",
				"ee.data.5.HistorySeasPWWW",
				"ee.data.6.HistorySeasOther",
				"ee.data.7.TotsAECurr",
				"ee.data.8.ManufactureAECurr",
				"ee.data.9.ServiceProdTPUAECurr",
				"ee.data.10.TradeAECurr",
				"ee.data.11.FireAECurr",
				"ee.data.12.ServicesAECurr",
				"ee.data.13.GovtAECurr",
				"ee.data.14.TotsWWPWCurr",
				"ee.data.15.ManufactureWWPWCurr",
				"ee.data.16.ServiceProdTPUWWPWCurr",
				"ee.data.17.TradeWWPWCurr",
				"ee.data.18.FireWWPWCurr",
				"ee.data.19.ServicesWWPWCurr",
				"ee.data.20.GovtWWCurr",
				"ee.data.21.TotsAHECurr",
				"ee.data.22.ManufactureAHECurr",
				"ee.data.23.ServiceProdTPUAHECurr",
				"ee.data.24.TradeAHECurr",
				"ee.data.25.FireAHECurr",
				"ee.data.26.ServicesAHECurr",
				"ee.data.27.TotsAWCurr",
				"ee.data.28.ManufactureAWCurr",
				"ee.data.29.ServiceProdTPUAWCurr",
				"ee.data.30.TradeAWCurr",
				"ee.data.31.FireAWCurr",
				"ee.data.32.ServicesAWCurr",
				"ee.data.33.TotsOtherCurr",
				"ee.data.34.ManufactureOtherCurr",
				"ee.data.35.ServiceProdTPUOtherCurr",
				"ee.data.36.TradeOtherCurr",
				"ee.data.37.FireOtherCurr",
				"ee.data.38.ServicesOtherCurr",
				"ee.data.39.TotsAEHist",
				"ee.data.40.ManufactureAEHist",
				//"ee.data.40a.ManufactureAEHist",
				//"ee.data.40b.ManufactureAEHist",
				"ee.data.41.ServiceProdTPUAEHist",
				"ee.data.42.TradeAEHist",
				"ee.data.43.FireAEHist",
				"ee.data.44.ServicesAEHist",
				"ee.data.45.GovtAEHist",
				"ee.data.46.TotsWWPWHist",
				"ee.data.47.ManufactureWWPWHist",
				//"ee.data.47a.ManufactureWWPWHist",
				//"ee.data.47b.ManufactureWWPWHist",
				"ee.data.48.ServiceProdTPUWWPWHist",
				"ee.data.49.TradeWWPWHist",
				"ee.data.50.FireWWPWHist",
				"ee.data.51.ServicesWWPWHist",
				"ee.data.52.GovtWWHist",
				"ee.data.53.TotsAHEHist",
				"ee.data.54.ManufactureAHEHist",
				//"ee.data.54a.ManufactureAHEHist",
				//"ee.data.54b.ManufactureAHEHist",
				"ee.data.55.ServiceProdTPUAHEHist",
				"ee.data.56.TradeAHEHist",
				"ee.data.57.FireAHEHist",
				"ee.data.58.ServicesAHEHist",
				"ee.data.59.TotsAWHist",
				"ee.data.60.ManufactureAWHist",
				//"ee.data.60a.ManufactureAWHist",
				//"ee.data.60b.ManufactureAWHist",
				//"ee.data.60c.ManufactureAWHist",
				"ee.data.61.ServiceProdTPUAWHist",
				"ee.data.62.TradeAWHist",
				"ee.data.63.FireAWHist",
				"ee.data.64.ServicesAWHist",
				"ee.data.65.TotsOtherHist",
				"ee.data.66.ManufactureOtherHist",
				"ee.data.67.ServiceProdTPUOtherHist",
				"ee.data.68.TradeOtherHist",
				"ee.data.69.FireOtherHist",
				"ee.data.70.ServicesOtherHist",
				"ee.data.71.WeeklyEarnings",
				"ee.data.72.WeeklyEarningsHist",
			},
			Schema: `
				CREATE SCHEMA IF NOT EXISTS ee;
				CREATE TABLE IF NOT EXISTS ee.datatype (
					data_type_code smallint PRIMARY KEY,
					data_type_text varchar(250)
				);
				CREATE TABLE IF NOT EXISTS ee.footnote (
					footnote_code char(1) PRIMARY KEY,
					footnote_text varchar(250)
				);
				CREATE TABLE IF NOT EXISTS ee.industry (
					industry_code     int PRIMARY KEY,
					SIC_code          smallint,
					publishing_status varchar(3),
					industry_name     varchar(250)
				);
				CREATE TABLE IF NOT EXISTS ee.period (
					period_code char(3) PRIMARY KEY,
					period_abbr varchar(5),
					period_name varchar(25)
				);
				CREATE TABLE IF NOT EXISTS ee.series (
					id             SERIAL PRIMARY KEY,
					series_id      char(11),
					industry_code  int REFERENCES ee.industry (industry_code),
					data_type_code smallint REFERENCES ee.datatype (data_type_code),
					seasonal       char(1),
					begin_year     int,
					begin_period   char(3) REFERENCES ee.period (period_code),
					end_year       int,
					end_period     char(3) REFERENCES ee.period (period_code)
				);
				CREATE TABLE IF NOT EXISTS ee.observation (
					id             SERIAL PRIMARY KEY,
					series_id      char(11),
					year           int,
					period         char(3) REFERENCES ee.period (period_code),
					value          int,
					footnote_codes char(1) REFERENCES ce.footnote (footnote_code),
					UNIQUE (series_id, year, period)
				);`,
		},
	}
}
