# miles
[![Build Status](https://travis-ci.org/asmarques/miles.svg)](https://travis-ci.org/asmarques/miles)

Calculate the distance flown for a given itinerary.

## Installation

```bash
go get github.com/asmarques/miles
```

## Usage

To calculate the distance for a given itinerary, specify the [IATA code](http://en.wikipedia.org/wiki/International_Air_Transport_Association_airport_code) of each airport along the route as an argument to `miles`:

```bash
$ miles LIS EWR SFO EWR LIS
0: LIS	EWR	3384	miles
1: EWR	SFO	2565	miles
2: SFO	EWR	2565	miles
3: EWR	LIS	3384	miles

	total:	11899	miles
```

The `miles` command supports the following options:
  - `-d` to specify the path to the airport database. If not provided, the following locations will be searched:
    - ./airports.csv
    - $GOPATH/src/github.com/asmarques/miles/airports.csv
  - `-u` download an updated copy of the airport database
  - `-o` choose the output format. Available formats are: `text` (default) and `json`.
  - `-v` to enable verbose output when using the `text` output format. Includes additional airport information such as airport name, country, latitute and longitude.

## License

[MIT](LICENSE)

The airport database (`airports.csv`) is obtained from [OurAirports.com](http://ourairports.com/data) which is available in the [public domain](http://en.wikipedia.org/wiki/Public_domain).