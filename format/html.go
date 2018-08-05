package format

// TODO option for inline style

import (
	"io"
	"fmt"
	"strings"

	"github.com/mbndr/lexy"
	"github.com/mbndr/lexy/style"
)

// cssClasses is a list in which the css class for each token is specified
var cssClasses = map[lexy.TokenType]string{
	lexy.TokenKeyword: "kw",
	lexy.TokenLiteral: "li",
	lexy.TokenBuiltin: "bu",
	lexy.TokenOperator: "op",
	lexy.TokenComment: "co",
	lexy.TokenString: "st",
	lexy.TokenNumber: "nu",
}

// WriteCss returns the css data of a style
func WriteCss(w io.Writer, s style.Style) {
	// "body"
	fmt.Fprintf(w,
		".lexy {background-color: %s; color: %s; display: block;}\n",
		s.Background, s.Foreground,
	)
	for typ, class := range cssClasses {
		fmt.Fprintf(w, ".%s {color: %s;}", class, s.TokenColors[typ])
	}
}

// HtmlTableFormatter formats the code into a HTML table
// With line numbers:
// <tr><td class="ln">1</td><td><span class="co">// My code</span></td></tr>
type HtmlFormatter struct {
	tokens []lexy.Token
}

func NewHtmlFormatter(ts []lexy.Token) HtmlFormatter {
	return HtmlFormatter{tokens: ts}
}

func (f *HtmlFormatter) Format(w io.Writer) error {
	fmt.Fprint(w, `<pre><code class="lexy">`)

	for _, t := range f.tokens {

		if t.Typ == lexy.TokenWS || t.Typ == lexy.TokenIdent {
			fmt.Fprint(w, t.Val)
			continue
		}

		fmt.Fprintf(w, `<span class="%s">%s</span>`, cssClasses[t.Typ], t.Val)

	}

	fmt.Fprint(w, `</code></pre>`)

	return nil
}



// htmlWhitespace transforms whitespaces to html entities
func htmlWhitespace(s string) string {
	s = strings.Replace(s, " ", "&nbsp;", -1)
	s = strings.Replace(s, "\t", "&nbsp;&nbsp;&nbsp;&nbsp;", -1)
	return s
}
