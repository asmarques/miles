package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/asmarques/geodist"
)

const kmToMiles = 0.62137

func generatePath(db database, codes []string) (*path, error) {
	var last *airport
	segments := make([]*segment, len(codes)-1)
	total := 0.0

	for i, code := range codes {
		code = strings.ToUpper(strings.TrimSpace(code))

		apt, ok := db[code]
		if !ok {
			return nil, fmt.Errorf("airport not found: %s", code)
		}

		if last != nil {
			distance, err := distance(last, apt)
			if err != nil {
				return nil, fmt.Errorf("error calculating distance between %s and %s\n", last.iata, apt.iata)
			}

			s := new(segment)
			s.origin = last
			s.destination = apt
			s.distance = distance

			segments[i-1] = s

			total += distance
		}
		last = apt
	}

	return &path{segments, total}, nil
}

func distance(ap1 *airport, ap2 *airport) (float64, error) {
	p1 := geodist.Point{ap1.lat, ap1.long}
	p2 := geodist.Point{ap2.lat, ap2.long}

	d, err := geodist.VincentyDistance(p1, p2)
	if err != nil {
		return math.NaN(), err
	}

	return d * kmToMiles, nil
}
