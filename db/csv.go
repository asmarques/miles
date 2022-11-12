package db

import (
	_ "embed"
	"encoding/csv"
	"io"
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

// Read reads the airport database from a CSV
func Read(reader io.Reader) (Database, error) {
	csvReader := csv.NewReader(reader)
	csvReader.FieldsPerRecord = recordSize
	records := 0

	db := make(database)

	for {
		record, err := csvReader.Read()
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
