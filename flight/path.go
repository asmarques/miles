package flight

// Path represents a list of flight segments
type Path struct {
	Segments []*Segment `json:"segments"`
}

// Distance retrieves the sum of the distances of flight segments which make up the path
func (p *Path) Distance() float64 {
	sum := float64(0)
	for _, segment := range p.Segments {
		sum += segment.Distance()
	}
	return sum
}
