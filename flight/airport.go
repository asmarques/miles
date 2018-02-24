package flight

// Airport represents the details of a given airport
type Airport struct {
	Name    string  `json:"name"`
	City    string  `json:"city"`
	Country string  `json:"countryCode"`
	Iata    string  `json:"iataCode"`
	Icao    string  `json:"icaoCode"`
	Lat     float64 `json:"latitude"`
	Long    float64 `json:"longitude"`
}
