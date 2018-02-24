package format

import (
	"io"

	"github.com/asmarques/miles/flight"
)

// Formatter is capable of formatting a path for a given output
type Formatter interface {
	Write(path *flight.Path, writer io.Writer, verbose bool) error
}
