package main

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

const recordSize = 18

const (
	nameField    = 3
	latField     = 4
	longField    = 5
	countryField = 8
	cityField    = 9
	icaoField    = 12
	iataField    = 13
)

func readAirports(file string) (database, error) {
	f, err := os.Open(file)
	defer f.Close()

	if err != nil {
		return nil, err
	}

	db := make(map[string]*airport)
	reader := csv.NewReader(f)
	reader.FieldsPerRecord = recordSize
	header := true

	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		if header {
			header = false
			continue
		}

		apt := new(airport)
		apt.name = record[nameField]
		apt.country = record[countryField]
		apt.city = record[cityField]
		apt.icao = record[icaoField]
		apt.iata = record[iataField]

		lat, err := strconv.ParseFloat(record[latField], 64)
		if err != nil {
			return nil, err
		}
		apt.lat = lat

		long, err := strconv.ParseFloat(record[longField], 64)
		if err != nil {
			return nil, err
		}
		apt.long = long

		db[apt.iata] = apt
	}

	return db, nil
}
