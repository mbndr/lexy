package lexy

import "fmt"

type TokenType int

const (
	// e.g. "x color val"
	// variable name or other identifier
	TokenIdent TokenType = iota
	// e.g. "if for return"
	// checked after scanned tokenIdent
	TokenKeyword
	// e.g. "true nil"
	// checked after scanned tokenIdent
	TokenLiteral
	// e.g. "append"
	// checked after scanned tokenIdent
	TokenBuiltin // check from ident
	// e.g. "+ - = < / . ( }"
	// operators and punctuation
	TokenOperator
	// e.g. "//x /*x*/
	TokenComment
	// e.g. `"string" 'c'`
	// highlighted with apostrophys
	TokenString
	// e.g. "0x664"
	TokenNumber
	// e.g. "\t ' ' \n"
	TokenWS
	// EOF
	// returned if the file reached the end
	TokenEOF
	// invalid token
	TokenInvalid
)

type Token struct {
	Typ TokenType
	Val string
}

func (t Token) String() string {
	return fmt.Sprintf("Token{%s, %q}", tokenTypeNames[t.Typ], t.Val)
}

var tokenTypeNames = map[TokenType]string{
	TokenKeyword:  "keywo",
	TokenLiteral:  "liter",
	TokenBuiltin:  "built",
	TokenOperator: "opera",
	TokenComment:  "comnt",
	TokenString:   "strng",
	TokenIdent:    "ident",
	TokenNumber:   "numbr",
	TokenEOF:      "__eof",
	TokenWS:       "whtsp",
	TokenInvalid:  "inval",
}

var eof = rune(0)
