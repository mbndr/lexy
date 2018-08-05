package lexy

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"unicode/utf8"

	"log"
)

// Lexer lexes through a reader
type Lexer struct {
	r  *bufio.Reader
	la Lang
	sf scanFunc
}

// NewLexer returns a new instance of a Lexer
func NewLexer(r io.Reader, la Lang) *Lexer {
	l := &Lexer{r: bufio.NewReader(r)}
	l.sf = scanIdent
	l.la = la
	return l
}

var (
	// used by scan functions to scan into
	buf bytes.Buffer
)

// Scan returns the next token
func (l Lexer) Scan() Token {
	buf.Reset()

	c := l.read()
	if c == eof {
		return Token{TokenEOF, ""}
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
func (l *Lexer) read() rune {
	c, _, err := l.r.ReadRune()
	if err != nil {
		return eof
	}
	return c
}

// peek returns what the next rune is
func (l *Lexer) peek() rune {
	b, err := l.r.Peek(1)
	if err != nil {
		return eof
	}
	c, _ := utf8.DecodeRune(b)
	return c
}

// unread unreads the last read rune
func (l *Lexer) unread() {
	err := l.r.UnreadRune()
	if err != nil {
		log.Fatal(err)
	}
}

// scanFunc is a scanner function for a specific token type
// each function returns
type scanFunc func(*Lexer) Token

// scanIdent returns a TokenIdent token
// another TokenType can be returned if the value matches a keyword etc
func scanIdent(l *Lexer) Token {
	// can change if we read e.g. a keyword
	var typ TokenType = TokenIdent

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

	if l.la.Keywords[s] {
		typ = TokenKeyword
	}
	if l.la.Literals[s] {
		typ = TokenLiteral
	}
	if l.la.Builtins[s] {
		typ = TokenBuiltin
	}

	return Token{typ, s}
}

// scanOperator returns a TokenOperator token
func scanOperator(l *Lexer) Token {

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

	return Token{TokenOperator, buf.String()}
}

// scanString returns a TokenString token
// can also be a char
func scanString(l *Lexer) Token {
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

		if c == '\\' && !nextEscaped {
			nextEscaped = true
			continue
		}

		if c == delimiter && !nextEscaped {
			break
		}
		nextEscaped = false
	}

	return Token{TokenString, buf.String()}
}

// scanNumber returns a TokenNumber token
// TODO specific number possibilities (prefix, suffix etc)?
func scanNumber(l *Lexer) Token {
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

	return Token{TokenNumber, buf.String()}
}

// scanComment returns a TokenComment token
// at this point it's verified that a comment begins
func scanComment(l *Lexer) Token {

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

	return Token{TokenComment, buf.String()}
}

// scanWhitespace returns a TokenWS token
func scanWhitespace(l *Lexer) Token {

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

	return Token{TokenWS, buf.String()}
}
