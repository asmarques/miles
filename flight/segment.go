package flight

import (
	"math"

	"github.com/asmarques/geodist"
)

const kmToMiles = 0.62137

// Segment represents a non-stop flight between origin and destination airports
type Segment struct {
	Origin      *Airport `json:"origin"`
	Destination *Airport `json:"destination"`
}

// Distance returns the distance (in miles) between the origin and destination airports
func (s *Segment) Distance() float64 {
	origin := geodist.Point{Lat: s.Origin.Lat, Long: s.Origin.Long}
	destination := geodist.Point{Lat: s.Destination.Lat, Long: s.Destination.Long}

	distance, err := geodist.VincentyDistance(origin, destination)
	if err != nil {
		return math.NaN()
	}

	return distance * kmToMiles
}
