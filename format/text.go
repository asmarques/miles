package format

import (
	"fmt"
	"io"
	"math"

	"github.com/asmarques/miles/flight"
)

// TextFormatter formats a path as plain text
var TextFormatter = &textFormatter{}

type textFormatter struct{}

func (tf *textFormatter) Write(path *flight.Path, writer io.Writer, verbose bool) error {
	for i, pair := range path.Segments {
		o := pair.Origin
		d := pair.Destination

		distance := int(math.Floor(pair.Distance() + 0.5))

		var line string
		if verbose {
			line = fmt.Sprintf("%d:\t%s - %s, %s (%f, %f)\n\t%s - %s, %s (%f, %f)\n\t%d miles\n\n",
				i, o.Iata, o.Name, o.Country, o.Lat, o.Long,
				d.Iata, d.Name, d.Country, d.Lat, d.Long, distance)
		} else {
			line = fmt.Sprintf("%d: %s\t%s\t%d\tmiles\n", i,
				o.Iata, d.Iata, distance)
		}

		_, err := fmt.Print(line)
		if err != nil {
			return fmt.Errorf("error writing output: %s", err)
		}
	}

	if len(path.Segments) > 1 {
		distance := int(math.Floor(path.Distance() + 0.5))

		var line string
		if verbose {
			line = fmt.Sprintf("total:\t%d miles\n", distance)
		} else {
			line = fmt.Sprintf("\n\ttotal:\t%d\tmiles\n", distance)
		}
		_, err := fmt.Print(line)
		if err != nil {
			return fmt.Errorf("error writing output: %s", err)
		}
	}

	return nil
}
