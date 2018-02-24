package db

import (
	"encoding/csv"
	"io"
	"net/http"
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

const (
	fallbackPath = "$GOPATH/src/github.com/asmarques/miles/airports.csv"
	srcURL       = "http://ourairports.com/data/airports.csv"
)

// Update updates the file containing the airport database
func Update(file string) error {
	updatefile := file + ".part"
	out, err := os.Create(updatefile)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(srcURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	os.Rename(updatefile, file)
	if err != nil {
		return err
	}

	return nil
}

// Read reads the airport database from a given file
func Read(file string) (Database, error) {
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		file = os.ExpandEnv(fallbackPath)
	}

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
