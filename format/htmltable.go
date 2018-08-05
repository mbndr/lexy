package format

// NOT COMPLETE / WORKING

// TODO customizable tabwidth
/*
import (
	"io"
	"fmt"
	"bytes"
	"strings"

	"github.com/mbndr/lexy"
)


// HtmlTableFormatter formats the code into a HTML table
// With line numbers:
// <tr><td class="ln">1</td><td><span class="co">// My code</span></td></tr>
type HtmlTableFormatter struct {
	tokens []lexy.Token
}

func NewHtmlTableFormatter(ts []lexy.Token) HtmlTableFormatter {
	return HtmlTableFormatter{tokens: ts}
}

func (f *HtmlTableFormatter) Format(w io.Writer) error {
	fmt.Fprint(w, `<table border="1" class="lexy">`)
	line := 1

	row := newTr(line)

	for _, t := range f.tokens {

		// check for new lines
		if t.Typ == lexy.TokenWS {
			for _, c := range t.Val {
				if c == '\n' {
					// a new line -> new row
					fmt.Fprint(w, row.String())
					row = newTr(line)
					line++
				}
			}
		}

		// non whitespace tokens are added to the row
		row.Add(t)
	}

	fmt.Fprint(w, `</table>`)

	return nil
}

// table row helper
type tr struct {
	line int
	tokens []lexy.Token
}

func newTr(l int) tr {
	return tr{line: l, tokens: []lexy.Token{}}
}

func (t *tr) Add(to lexy.Token) {
	t.tokens = append(t.tokens, to)
}

func (t *tr) String() string {
	var buf bytes.Buffer

	buf.WriteString("<tr>")
	buf.WriteString(fmt.Sprintf(`<td class="line">%d</td>`, t.line))
	// TODO line number
	buf.WriteString(`<td class="code">`)

	for _, to := range t.tokens {
		s := to.Val

		if to.Typ == lexy.TokenWS {
			s = htmlWhitespace(s)
		}

		// TODO multi line comments
		if strings.ContainsRune(s, '\n') {

		}

		buf.WriteString(
			fmt.Sprintf(`<span class="%s">%s</span>`, cssClasses[to.Typ], s),
		)
	}

	buf.WriteString("</td></tr>\n")

	return buf.String()
}

/*

fmt.Fprint(w, "<tr>")

		// line number
		if f.LineNumbers {
			fmt.Fprintf(w, "<td>%d</td>", i+1)
		}
		// code
		fmt.Fprint(w, "<td>")

		// line break
		if t.Typ == lexy.TokenWS {
			for _, c := range t.Val {
				if c == '\n' {
					fmt.Fprint(w, "</td></tr>")
					continue
				}
			}
		}
*/
