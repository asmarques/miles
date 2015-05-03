package main

import (
	"fmt"
	"io"
	"math"
)

func printText(p *path, w io.Writer, v bool) {
	for i, pair := range p.segments {
		o := pair.origin
		d := pair.destination

		if v {
			fmt.Fprintf(w, "%d:\t%s - %s, %s (%f, %f)\n\t%s - %s, %s (%f, %f)\n\t", i,
				o.iata, o.name, o.country, o.lat, o.long,
				d.iata, d.name, d.country, d.lat, d.long)
		} else {
			fmt.Fprintf(w, "%d: %s\t%s\t", i,
				o.iata, d.iata)
		}

		distance := int(math.Floor(pair.distance + 0.5))

		if v {
			fmt.Fprintf(w, "%d miles\n\n", distance)
		} else {
			fmt.Fprintf(w, "%d\tmiles\n", distance)
		}
	}

	if len(p.segments) > 1 {
		distance := int(math.Floor(p.distance + 0.5))
		if v {
			fmt.Fprintf(w, "total:\t%d miles\n", distance)
		} else {
			fmt.Fprintf(w, "\n\ttotal:\t%d\tmiles\n", distance)
		}
	}
}
