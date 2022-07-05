package api

type Model struct {
	CountryCode *string      `json:"countryCode"`
	CityName    *string      `json:"cityName"`
	IpAddress   *string      `json:"ipAddress"`
	Coordinates *Coordinates `json:"coordinates"`
}

type Coordinates struct {
	Latitude  *string `json:"latitude"`
	Longitude *string `json:"longitude"`
}
