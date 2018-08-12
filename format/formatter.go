package format

import "github.com/mbndr/lexy"

// Formatter is an interface for a structure which is responsible
// to format into a specific format (e.g. HTML, Image...)
type Formatter interface {
	// Format formats the tokens
	// 
	// number of writers and the need for a style can vary with each formatter,
	// so only tokens are mandatory in the Format method
	Format([]lexy.Token) error
}
