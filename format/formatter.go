package format

import "github.com/mbndr/lexy"

// Formatter is an interface for a structure which is responsible
// to format into a specific format (e.g. HTML, Terminal...)
// The reader where to write has to be given in a construct func
type Formatter interface {
	Format(lexy.Style) error
}
