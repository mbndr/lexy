package format

import (
	"io"
)

type Formatter interface {
	Format(io.Writer) error
}
