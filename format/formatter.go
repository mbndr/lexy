package format

import (
	"io"
)

// Formatter is an interface for a structure which is responsible
// to format into a specific format (e.g. HTML, Terminal...)
type Formatter interface {
	Format(io.Writer) error
}
