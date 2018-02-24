package format

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/asmarques/miles/flight"
)

// JSONFormatter formats a path as plain text
var JSONFormatter = &jsonFormatter{}

type jsonFormatter struct{}

const unitMiles = "miles"

type jsonSegment struct {
	flight.Segment
	Distance float64 `json:"distance"`
}

type jsonPath struct {
	Segments      []jsonSegment `json:"segments"`
	TotalDistance float64       `json:"totalDistance"`
	Unit          string        `json:"unit"`
}

func (jf *jsonFormatter) Write(path *flight.Path, writer io.Writer, verbose bool) error {
	var segments []jsonSegment

	for _, s := range path.Segments {
		segments = append(segments, jsonSegment{
			Segment:  *s,
			Distance: s.Distance(),
		})
	}

	result, err := json.Marshal(jsonPath{
		Segments:      segments,
		TotalDistance: path.Distance(),
		Unit:          unitMiles,
	})
	if err != nil {
		fmt.Printf("error generating json output: %s", err)
	}

	_, err = fmt.Print(result)
	if err != nil {
		return fmt.Errorf("error writing to output: %s", err)
	}

	return nil
}
