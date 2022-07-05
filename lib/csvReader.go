package lib

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/grahms/geolocationservice/util"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"os"
	"regexp"
)

type GeoCSV struct {
	CSVModel
}

func NewCsv() *GeoCSV {
	return &GeoCSV{}
}

// Parse takes a filename and a callback function
func (g *GeoCSV) Parse(filePath string, callback func(loc GeoCSV)) {
	//ip_address,country_code,country,city,latitude,longitude,mystery_value
	in, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer func(in *os.File) {
		err := in.Close()
		if err != nil {

		}
	}(in)

	// locations: the channel that will receive locations
	locations := make(chan GeoCSV)
	cerr := make(chan error)

	// launch a goroutine to load locations to locations channel
	go func() {
		cerr <- gocsv.UnmarshalToChan(in, locations)
	}()
	for {
		select {
		// panic if chanel returns an error
		case err := <-cerr:
			panic(err)
		default:
			loc := <-locations
			// locations will be loaded to callback function
			callback(loc)
		}

	}

}

// isValidIP takes an ip and validates with a regular expression
func (g *GeoCSV) isValidIP(ip *string) bool {
	matchString := `^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$`
	b, err := regexp.MatchString(matchString, *ip)
	if err != nil {
		return false
	}
	return b
}

// isValidCountryCode takes a country code and validates with a regular expression
func (g *GeoCSV) isValidCountryCode(countryCode *string) bool {
	matchString := `^[A-Z]{2}`
	return regxValidator(matchString, *countryCode)
}

// isValidCityName takes a city name and validates with a regular expression
func (g *GeoCSV) isValidCityName(cityName *string) bool {
	matchString := `^[a-zA-Z]+(?:[\s-][a-zA-Z]+)*$`
	return regxValidator(matchString, *cityName)
}

// isValidCountry takes a country name and validates with a regular expression
func (g *GeoCSV) isValidCountry(country *string) bool {
	matchString := `^[a-zA-Z]+(?:[\s-][a-zA-Z]+)*$`
	return regxValidator(matchString, *country)
}

// isValidCoordinate takes latitude & longitude and validates with a regular expression
func (g *GeoCSV) isValidCoordinate(lat *string, long *string) bool {
	coord := *lat + "," + *long
	matchString := `^((\-?|\+?)?\d+(\.\d+)?),\s*((\-?|\+?)?\d+(\.\d+)?)$`
	return regxValidator(matchString, coord)

}
func (g *GeoCSV) ValidateRow(row *GeoCSV) bool {
	mapRow := make(map[string]bool)
	mapRow[util.IPADDRKEY] = g.isValidIP(&row.IPAddress)
	mapRow[util.COUNTRYCODEKEY] = g.isValidCountryCode(&row.CountryCode)
	mapRow[util.COUNTRYNAMEKEY] = g.isValidCountry(&row.Country)
	mapRow[util.CITY] = g.isValidCityName(&row.City)
	mapRow[util.COORDS] = g.isValidCoordinate(&row.Latitude, &row.Longitude)

	for k, v := range mapRow {
		if v == false {
			print(util.ColorRed, "invalid row: ")
			_ = ginmetrics.GetMonitor().GetMetric("csv_rows").Inc([]string{k, "INVALID"})
			fmt.Printf("%v", row)
			println()
			return false
		}

	}
	_ = ginmetrics.GetMonitor().GetMetric("csv_rows").Inc([]string{"VALID"})
	print(util.ColorGreen, "valid row: ")
	fmt.Printf("%v", row)
	println()
	return true

}

func regxValidator(matchString string, s string) bool {
	b, err := regexp.MatchString(matchString, s)
	if err != nil {
		return false
	}
	return b
}
