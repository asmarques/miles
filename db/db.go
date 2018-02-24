package db

import "github.com/asmarques/miles/flight"

// Database represents an airport database
type Database interface {
	Count() int
	GetAirport(iataCode string) *flight.Airport
}

type database map[string]*flight.Airport

func (d database) Count() int {
	return len(d)
}

func (d database) GetAirport(iataCode string) *flight.Airport {
	airport, _ := d[iataCode]
	return airport
}
