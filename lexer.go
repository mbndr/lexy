package lexy

import (
	"bufio"
	"io"
	"unicode/utf8"
	"bytes"
	"strings"

	"log"

	"github.com/mbndr/lexy/lang"
)

// lexer lexes through a reader
type lexer struct {
	r *bufio.Reader
	la lang.Lang
	sf scanFunc
}

// NewLexer returns a new instance of a lexer
func NewLexer(r io.Reader, la lang.Lang) *lexer {
	l := &lexer{r: bufio.NewReader(r)}
	l.sf = scanIdent
	l.la = la
	return l
}


var (
	// used by scan functions to scan into
	buf bytes.Buffer
)


// Scan returns the next token
func (l lexer) Scan() token {
	buf.Reset()

	c := l.read()
	if c == eof {
		return token{tokenEOF, ""}
	}
	l.unread() // will be reread by scanFunc

	// get the scanFunc
	if isWhitespace(c) {
		l.sf = scanWhitespace
	} else if couldBeComment(c, l.peek()) {
		l.sf = scanComment
	} else if couldBeNumber(c, l.peek()) {
		l.sf = scanNumber
	} else if isLetter(c) || c == '_' {
		l.sf = scanIdent
	} else if strings.ContainsRune(l.la.Operators, c) {
		l.sf = scanOperator
	} else if c == '"' || c == '\'' || c == '`' {
		l.sf = scanString
	}

	return l.sf(&l)
}

// read reads the next rune
func (l *lexer) read() rune {
	c, _, err := l.r.ReadRune()
	if err != nil {
		return eof
	}
	return c
}

// peek returns what the next rune is
func (l *lexer) peek() rune {
	b, err := l.r.Peek(1)
	if err != nil {
		log.Println(err) // TODO
		return eof
	}
	c, _ := utf8.DecodeRune(b)
	return c
}

// unread unreads the last read rune
func (l *lexer) unread() {
	err := l.r.UnreadRune()
	if err != nil {
		log.Fatal(err)
	}
}

// scanFunc is a scanner function for a specific token type
// each function returns
type scanFunc func(*lexer) token

// scanIdent returns a tokenIdent token
func scanIdent(l *lexer) token {
	// can change if we read e.g. a keyword
	var typ tokenType = tokenIdent

	// first char is approved a ident
	// cannot be done in the for loop because then a digit would be a valid first identifier char
	buf.WriteRune(l.read())

	for {
		c := l.read()
		if c == eof {
			break
		}

		if isLetter(c) || isDigit(c) || c == '_' {
			buf.WriteRune(c)
			continue
		}

		// invalid ident char (operator / punctuation etc coming)
		l.unread()
		break
	}

	// check for keywords etc
	s := buf.String()
	
	if l.la.Keywords[s] { typ = tokenKeyword }
	if l.la.Literals[s] { typ = tokenLiteral }
	if l.la.Builtins[s] { typ = tokenBuiltin }

	return token{typ, s}
}

// scanOperator returns a tokenOperator token
func scanOperator(l *lexer) token {

	for {
		c := l.read()
		if c == eof {
			break
		}

		if !strings.ContainsRune(l.la.Operators, c) {
			l.unread()
			break
		}

		buf.WriteRune(c)
	}

	return token{tokenOperator, buf.String()}
}

// scanString returns a tokenString token
// can also be a char
func scanString(l *lexer) token {
	var nextEscaped bool

	// '"' or '\''
	delimiter := l.read()
	buf.WriteRune(delimiter)

	for {
		c := l.read()
		if c == eof {
			break
		}

		buf.WriteRune(c)

		if c == '\\' {
			nextEscaped = true
			continue
		}

		if c == delimiter && !nextEscaped {
			break
		}
		nextEscaped = false
	}

	return token{tokenString, buf.String()}
}

// scanNumber returns a tokenNumber token
// TODO specific number possibilities (prefix, suffix etc)
func scanNumber(l *lexer) token {
	// first char is surely part of a number (maybe '.')
	buf.WriteRune(l.read())

	for {
		c := l.read()
		if c == eof {
			break
		}

		if !isDigit(c) {
			l.unread()
			break
		}

		buf.WriteRune(c)
	}

	return token{tokenNumber, buf.String()}
}

// scanComment returns a tokenComment token
// at this point it's verified that a comment begins
func scanComment(l *lexer) token {

	buf.WriteRune(l.read()) // first comment indicator
	c := l.read()

	buf.WriteRune(c)
	
	if c == '/' {
		// line comment
		for {
			c := l.read()
			if c == eof {
				break
			}

			if c == '\n' {
				l.unread()
				break
			}

			buf.WriteRune(c)
		}
	} else if c == '*' {
		// block comment
		for {
			c := l.read()
			if c == eof {
				break
			}

			if c == '*' && l.peek() == '/' {
				buf.WriteRune(c)        // '*'
				buf.WriteRune(l.read()) // '/'
				break
			}

			buf.WriteRune(c)
		}
	}

	return token{tokenComment, buf.String()}
}

// scanWhitespace returns a tokenWS token
func scanWhitespace(l *lexer) token {
	
	for {
		c := l.read()
		if c == eof {
			break
		}

		if !isWhitespace(c) {
			l.unread()
			break
		}

		buf.WriteRune(c)
	}

	return token{tokenWS, buf.String()}
}
