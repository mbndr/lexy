package format

import (
	"io"

	"github.com/mbndr/lexy"
)

type HtmlFormatter struct {
	tokens []lexy.Token
}

func NewHtmlFormatter(ts []lexy.Token) HtmlFormatter {
	return HtmlFormatter{tokens: ts}
}

func (f *HtmlFormatter) Format(w io.Writer) error {
	return nil
}