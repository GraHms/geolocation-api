package lib

type CSVModel struct {
	IPAddress    string `csv:"ip_address"` // .csv column headers
	CountryCode  string `csv:"country_code"`
	Country      string `csv:"country"`
	City         string `csv:"city"`
	Latitude     string `csv:"latitude"`
	Longitude    string `csv:"longitude"`
	MysteryValue string `csv:"mystery_value"`
}
