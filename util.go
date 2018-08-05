package lexy

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
	return ( isDigit(c) || (c == '.' && isDigit(n)) )
}

func IsTokenEOF(t token) bool {
	return t.typ == tokenEOF
}