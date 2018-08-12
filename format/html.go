package format

// TODO option for inline style?

import (
	"fmt"
	"io"
	"strings"

	"github.com/mbndr/lexy"
)

// cssClasses is a list in which the css class for each token is specified
var cssClasses = map[lexy.TokenType]string{
	lexy.TokenKeyword:  "kw",
	lexy.TokenLiteral:  "li",
	lexy.TokenBuiltin:  "bu",
	lexy.TokenOperator: "op",
	lexy.TokenComment:  "co",
	lexy.TokenString:   "st",
	lexy.TokenNumber:   "nu",
}

// HtmlFormatter formats the code into a html <pre><code> construct
type HtmlFormatter struct {
	tokens []lexy.Token
	cssWriter io.Writer
	htmlWriter io.Writer
}

// NewHtmlFormatter returns a new HtmlFormatter.
// For better separation of CSS and HTML there have to be given two writers
func NewHtmlFormatter(ts []lexy.Token, hw, cw io.Writer) HtmlFormatter {
	return HtmlFormatter{tokens: ts, cssWriter: cw, htmlWriter: hw}
}

// Format writes the HTML and CSS
func (f *HtmlFormatter) Format(s lexy.Style) error {
	fmt.Fprint(f.htmlWriter, `<pre><code class="lexy">`)

	for _, t := range f.tokens {

		if t.Typ == lexy.TokenWS || t.Typ == lexy.TokenIdent {
			fmt.Fprint(f.htmlWriter, t.Val)
			continue
		}

		fmt.Fprintf(f.htmlWriter, `<span class="%s">%s</span>`, cssClasses[t.Typ], t.Val)

	}

	fmt.Fprint(f.htmlWriter, `</code></pre>`)

	f.writeCss(s)

	return nil
}

// writeCss writes the css of a style
func (f *HtmlFormatter) writeCss(s lexy.Style) {
	// "body"
	fmt.Fprintf(f.cssWriter,
		".lexy {background-color: %s; color: %s; display: block; padding: 10px;}\n",
		s.Background, s.Foreground,
	)
	// code
	for typ, class := range cssClasses {
		fmt.Fprintf(f.cssWriter, ".%s {color: %s;}", class, s.TokenColors[typ])
	}
}

// htmlWhitespace transforms whitespaces to html entities
// TODO configurable spaces for tabs
func htmlWhitespace(s string) string {
	s = strings.Replace(s, " ", "&nbsp;", -1)
	s = strings.Replace(s, "\t", "&nbsp;&nbsp;&nbsp;&nbsp;", -1)
	return s
}
