package main

import (
	"flag"
	"fmt"
	"os"
)

const fallbackDbPath = "$GOPATH/src/github.com/asmarques/miles/airports.csv"

var (
	dbPath = flag.String("db", "airports.csv", "path to airport database")
	verbose = flag.Bool("v", false, "verbose output")
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [-v] [-db file] ap1 ap2 ...\n", os.Args[0])
	flag.PrintDefaults()
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
		os.Exit(2)
	}

	db, err := readAirports(*dbPath)
	if err != nil {
		exit(fmt.Errorf("error reading airport database: %s", err))
	}

	p, err := generatePath(db, route)
	if err != nil {
		exit(fmt.Errorf("error generating path: %s", err))
	}

	printText(p, os.Stdout, *verbose)
}
