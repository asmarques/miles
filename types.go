package main

import (
	"fmt"
	"io"
)

type airport struct {
	Name    string  `json:"name"`
	City    string  `json:"city"`
	Country string  `json:"countryCode"`
	Iata    string  `json:"iataCode"`
	Icao    string  `json:"icaoCode"`
	Lat     float64 `json:"latitude"`
	Long    float64 `json:"longitude"`
}

func (a *airport) String() string {
	return fmt.Sprintf("%s", a.Iata)
}

type segment struct {
	Origin      *airport `json:"origin"`
	Destination *airport `json:"destination"`
	Distance    float64  `json:"distance"`
}

func (s *segment) String() string {
	return fmt.Sprintf("<%s -> %s (%f)>", s.Origin, s.Destination, s.Distance)
}

type path struct {
	Segments []*segment `json:"segments"`
	Distance float64    `json:"totalDistance"`
	Unit     string     `json:"unit"`
}

func (p *path) String() string {
	return fmt.Sprintf("<%s (%f)>", p.Segments, p.Distance)
}

type database map[string]*airport

type formatter func(*path, io.Writer, bool)
