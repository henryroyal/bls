package employment

import (
	"github.com/henryroyal/bls/series"
	"regexp"
)

/*
	                      1         2
	             1234567890123456789012345678
	Series ID    BDS0000000000000000110101LQ5
	Positions    Value        Field Name
	1-2          BD           Prefix
	3            S            Seasonal Adjustment Code
	4-13         0000000000	  Area Code
	14-19        000000       Industry Code
	20           0            Unit Analysis Code
	21           0            Data Element Code
	22-23        01           Size Class Code
	24-25        01           Data Class Code
	26           L            Rate/Level Code
	27           Q            Record Type Code
	28           5            Ownership Code
 */

var (
	seriesIdFormat *regexp.Regexp = regexp.MustCompile(
		`(?P<p>\w{2})(?P<sac>\w{1})(?P<ac>\d{10})(?P<ic>\d{6})(?P<uac>\d{1})(?P<dec>\d{1})(?P<scc>\w{2})(?P<dcc>\w{2})(?P<rlc>\w{1})(?P<rtc>\w{1})(?P<oc>\w{1})`,
	)
)

type BusinessEmploymentDynamics struct {
	series.Dataset
}

func NewBusinessEmploymentDynamics() (*BusinessEmploymentDynamics) {
	return &BusinessEmploymentDynamics{
		series.Dataset{
			Name:         "Business Employment Dynamics",
			Symbol:       "bd",
			BaseURL:      series.BaseURL,
			SeriesFormat: *seriesIdFormat,
			Schema: `
				CREATE SCHEMA IF NOT EXISTS bd;
				CREATE TABLE IF NOT EXISTS bd.county (
					county_code smallint PRIMARY KEY,
					county_name varchar(12)
				);
				CREATE TABLE IF NOT EXISTS bd.dataclass (
					dataclass_code char(2) PRIMARY KEY,
					dataclass_name varchar(50),
					display_level  smallint,
					selectable     boolean,
					sort_sequence  smallint
				);
				CREATE TABLE IF NOT EXISTS bd.dataelement (
					dataelement_code smallint PRIMARY KEY,
					dataelement_name varchar(50)
				);
				CREATE TABLE IF NOT EXISTS bd.footnote (
					footnote_code smallint PRIMARY KEY,
					footnote_text text
				);
				CREATE TABLE IF NOT EXISTS bd.industry (
					industry_code int PRIMARY KEY,
					industry_name varchar(100),
					display_level smallint,
					selectable    char(1),
					sort_sequence smallint
				);
				CREATE TABLE IF NOT EXISTS bd.msa (
					msa_code smallint PRIMARY KEY,
					msa_name varchar(50)
				);
				CREATE TABLE IF NOT EXISTS bd.ownership (
					ownership_code smallint PRIMARY KEY,
					ownership_name varchar(5)
				);
				CREATE TABLE IF NOT EXISTS bd.periodicity (
					periodicity_code char(1) PRIMARY KEY,
					periodicity_name varchar(50)
				);
				CREATE TABLE IF NOT EXISTS bd.ratelevel (
					ratelevel_code char(1) PRIMARY KEY,
					ratelevel_name varchar(50)
				);
				CREATE TABLE IF NOT EXISTS bd.seasonal (
					seasonal_code char(1) PRIMARY KEY,
					seasonal_text varchar(50)
				);
				CREATE TABLE IF NOT EXISTS bd.sizeclass (
					sizeclass_code smallint PRIMARY KEY,
					sizeclass_name varchar(50)
				);
				CREATE TABLE IF NOT EXISTS bd.state (
					state_code smallint PRIMARY KEY,
					state_name varchar(50)
				);
				CREATE TABLE IF NOT EXISTS bd.unitanalysis (
					unitanalysis_code smallint PRIMARY KEY,
					unitanalysis_name varchar(5)
				);
				CREATE TABLE IF NOT EXISTS bd.series (
					series_id         char(28) PRIMARY KEY,
					seasonal          char(1),
					msa_code          smallint,
					state_code        smallint,
					county_code       smallint,
					industry_code     int,
					unitanalysis_code smallint,
					dataelement_code  smallint,
					sizeclass_code    smallint,
					dataclass_code    char(2),
					ratelevel_code    char(1),
					periodicity_code  char(1),
					ownership_code    smallint,
					series_title      text,
					footnote_codes    smallint,
					begin_year        smallint,
					begin_period      char(3),
					end_year          smallint,
					end_period        char(3)
				);
				CREATE TABLE IF NOT EXISTS bd.observations (
					id             SERIAL PRIMARY KEY,
					series_id      char(28) REFERENCES bd.series (series_id),
					year           int,
					period         char(3),
					value          int,
					footnote_codes int
				);`,
		},
	}
}
