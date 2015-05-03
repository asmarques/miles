package main

import (
	"fmt"
	"io"
)

type airport struct {
	name    string
	city    string
	country string
	iata    string
	icao    string
	lat     float64
	long    float64
}

func (a *airport) String() string {
	return fmt.Sprintf("%s", a.iata)
}

type segment struct {
	origin, destination *airport
	distance            float64
}

func (s *segment) String() string {
	return fmt.Sprintf("<%s -> %s (%f)>", s.origin, s.destination, s.distance)
}

type path struct {
	segments []*segment
	distance float64
}

func (p *path) String() string {
	return fmt.Sprintf("<%s (%f)>", p.segments, p.distance)
}

type database map[string]*airport

type formatter func(*path, io.Writer)
