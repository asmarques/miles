package main

import (
	_ "embed"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/asmarques/miles/db"
	"github.com/asmarques/miles/flight"
	"github.com/asmarques/miles/format"
)

var (
	dbPath       = flag.String("d", "", "path to airport database (use embedded database if not supplied)")
	outputFormat = flag.String("o", "text", "output format (text, json)")
	verbose      = flag.Bool("v", false, "verbose output")
)

//go:embed airports.csv
var airportsCsv string

func usage() {
	log.Printf("usage: %s [-d file] [-o text|json|kml] [-v] ap1 ap2 ...\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	route := flag.Args()
	if len(route) < 2 {
		flag.Usage()
	}

	var formatter format.Formatter

	switch *outputFormat {
	case "text":
		formatter = format.TextFormatter
	case "json":
		formatter = format.JSONFormatter
	case "kml":
		formatter = format.KMLFormatter
	default:
		flag.Usage()
	}

	var reader io.Reader
	if *dbPath != "" {
		f, err := os.Open(*dbPath)
		if err != nil {
			log.Fatalf("error reading airport database file: %s", err)
		}
		defer f.Close()
		reader = f
	} else {
		reader = strings.NewReader(airportsCsv)
	}

	db, err := db.Read(reader)
	if err != nil {
		log.Fatalf("error processing airport database: %s", err)
	}

	path, err := generatePath(db, route)
	if err != nil {
		log.Fatalf("error generating path: %s", err)
	}

	err = formatter.Write(path, os.Stdout, *verbose)
	if err != nil {
		log.Fatalf("error formatting path: %s", err)
	}
}

func generatePath(db db.Database, codes []string) (*flight.Path, error) {
	var lastAirport *flight.Airport
	var segments []*flight.Segment

	for _, code := range codes {
		code = strings.ToUpper(strings.TrimSpace(code))

		airport := db.GetAirport(code)
		if airport == nil {
			return nil, fmt.Errorf("airport not found: %s", code)
		}

		if lastAirport != nil {
			segments = append(segments, &flight.Segment{
				Origin:      lastAirport,
				Destination: airport,
			})
		}
		lastAirport = airport
	}

	path := &flight.Path{Segments: segments}
	return path, nil
}
