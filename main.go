package main

import (
	"flag"
	"fmt"
	"os"
)

const fallbackDbPath = "$GOPATH/src/github.com/asmarques/miles/airports.csv"

var (
	dbPath       = flag.String("d", "airports.csv", "path to airport database")
	outputFormat = flag.String("o", "text", "output format (text, json)")
	verbose      = flag.Bool("v", false, "verbose output")
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [-d file] [-o text|json] [-v] ap1 ap2 ...\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}

func exit(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	_, err := os.Stat(*dbPath)
	if os.IsNotExist(err) {
		*dbPath = os.ExpandEnv(fallbackDbPath)
	}

	route := flag.Args()
	if len(route) < 2 {
		flag.Usage()
	}

	var f formatter

	switch *outputFormat {
	case "text":
		f = printText
	case "json":
		f = printJson
	default:
		flag.Usage()
	}

	db, err := readAirports(*dbPath)
	if err != nil {
		exit(fmt.Errorf("error reading airport database: %s", err))
	}

	p, err := generatePath(db, route)
	if err != nil {
		exit(fmt.Errorf("error generating path: %s", err))
	}

	f(p, os.Stdout, *verbose)
}
