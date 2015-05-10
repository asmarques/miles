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
		apt.Name = record[nameField]
		apt.Country = record[countryField]
		apt.City = record[cityField]
		apt.Icao = record[icaoField]
		apt.Iata = record[iataField]

		lat, err := strconv.ParseFloat(record[latField], 64)
		if err != nil {
			return nil, err
		}
		apt.Lat = lat

		long, err := strconv.ParseFloat(record[longField], 64)
		if err != nil {
			return nil, err
		}
		apt.Long = long

		db[apt.Iata] = apt
	}

	return db, nil
}
