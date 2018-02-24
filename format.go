package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
)

func printText(p *path, w io.Writer, v bool) {
	for i, pair := range p.Segments {
		o := pair.Origin
		d := pair.Destination

		if v {
			fmt.Fprintf(w, "%d:\t%s - %s, %s (%f, %f)\n\t%s - %s, %s (%f, %f)\n\t", i,
				o.Iata, o.Name, o.Country, o.Lat, o.Long,
				d.Iata, d.Name, d.Country, d.Lat, d.Long)
		} else {
			fmt.Fprintf(w, "%d: %s\t%s\t", i,
				o.Iata, d.Iata)
		}

		distance := int(math.Floor(pair.Distance + 0.5))

		if v {
			fmt.Fprintf(w, "%d miles\n\n", distance)
		} else {
			fmt.Fprintf(w, "%d\tmiles\n", distance)
		}
	}

	if len(p.Segments) > 1 {
		distance := int(math.Floor(p.Distance + 0.5))
		if v {
			fmt.Fprintf(w, "total:\t%d miles\n", distance)
		} else {
			fmt.Fprintf(w, "\n\ttotal:\t%d\tmiles\n", distance)
		}
	}
}

func printJSON(p *path, w io.Writer, v bool) {
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("error generating json output")
	}

	fmt.Printf("%s", b)
}
