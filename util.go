package lexy

import "io"

func isWhitespace(c rune) bool {
	return c == ' ' || c == '\t' || c == '\n'
}

func isLetter(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func isDigit(c rune) bool {
	return (c >= '0' && c <= '9')
}

func couldBeComment(c, n rune) bool {
	return (c == '/' && (n == '*' || n == '/'))
}

func couldBeNumber(c, n rune) bool {
	return (isDigit(c) || (c == '.' && isDigit(n)))
}

func IsTokenEOF(t Token) bool {
	return t.Typ == TokenEOF
}

func ScanAll(r io.Reader, la Lang) []Token {
	l := NewLexer(r, la)

	var tokens []Token

	for {
		t := l.Scan()
		tokens = append(tokens, t)
		if IsTokenEOF(t) {
			break
		}
	}

	return tokens
}
