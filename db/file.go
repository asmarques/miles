package db

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"

	"github.com/asmarques/miles/flight"
)

const recordSize = 18

const (
	nameField    = 3
	latField     = 4
	longField    = 5
	countryField = 8
	cityField    = 10
	icaoField    = 12
	iataField    = 13
)

// ReadAirports reads the airport database from a given file
func ReadAirports(file string) (Database, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	db := make(database)
	reader := csv.NewReader(f)
	reader.FieldsPerRecord = recordSize
	records := 0

	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		records++

		if records == 1 {
			continue
		}

		lat, err := strconv.ParseFloat(record[latField], 64)
		if err != nil {
			return nil, err
		}

		long, err := strconv.ParseFloat(record[longField], 64)
		if err != nil {
			return nil, err
		}

		airport := &flight.Airport{
			Name:    record[nameField],
			Country: record[countryField],
			City:    record[cityField],
			Icao:    record[icaoField],
			Iata:    record[iataField],
			Lat:     lat,
			Long:    long,
		}

		db[airport.Iata] = airport
	}

	return db, nil
}
