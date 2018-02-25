package format

import (
	"encoding/xml"
	"fmt"
	"io"

	"github.com/asmarques/miles/flight"
)

// KMLFormatter formats a path as a KML document
var KMLFormatter = &kmlFormatter{}

type kmlFormatter struct{}

type kmlDocument struct {
	XMLName    xml.Name `xml:"Document"`
	Name       string   `xml:"name"`
	Styles     []kmlStyle
	Placemarks []kmlPlacemark
}

type kmlStyle struct {
	XMLName   xml.Name `xml:"Style"`
	ID        string   `xml:"id,attr"`
	LineColor string   `xml:"LineStyle>color,omitempty"`
	LineWidth int      `xml:"LineStyle>width,omitempty"`
}

type kmlPlacemark struct {
	XMLName    xml.Name `xml:"Placemark,omitempty"`
	Name       string   `xml:"name,omitempty"`
	Style      string   `xml:"styleUrl,omitempty"`
	LineString kmlLineString
}

type kmlLineString struct {
	Tessellate  int    `xml:"tessellate,omitempty"`
	Coordinates string `xml:"coordinates,omitempty"`
}

const (
	kmlFormat = xml.Header + `<kml xmlns="http://www.opengis.net/kml/2.2" xmlns:gx="http://www.google.com/kml/ext/2.2" xmlns:kml="http://www.opengis.net/kml/2.2" xmlns:atom="http://www.w3.org/2005/Atom">` + "\n%s\n</kml>"
	xmlIndent = "    "
	styleID   = "style"
	lineColor = "ff0000ff"
	lineWidth = 5
)

func (kf *kmlFormatter) Write(path *flight.Path, writer io.Writer, verbose bool) error {
	var placemarks []kmlPlacemark

	styles := []kmlStyle{
		kmlStyle{
			ID:        styleID,
			LineColor: lineColor,
			LineWidth: lineWidth,
		},
	}

	for _, segment := range path.Segments {
		coordinates := fmt.Sprintf("%f,%f,0 %f,%f,0",
			segment.Origin.Long, segment.Origin.Lat,
			segment.Destination.Long, segment.Destination.Lat)
		placemarks = append(placemarks, kmlPlacemark{
			Name:  fmt.Sprintf("%s-%s", segment.Origin.Iata, segment.Destination.Iata),
			Style: styleID,
			LineString: kmlLineString{
				Tessellate:  1,
				Coordinates: coordinates,
			},
		})
	}

	doc := &kmlDocument{
		Styles:     styles,
		Placemarks: placemarks,
	}

	output, err := xml.MarshalIndent(doc, xmlIndent, xmlIndent)
	if err != nil {
		return err
	}

	fmt.Printf(kmlFormat, output)

	return nil
}
