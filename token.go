package lexy

import "fmt"

type tokenType int

const (
	// e.g. "x color val"
	// variable name or other identifier
	tokenIdent tokenType = iota
	// e.g. "if for return"
	// checked after scanned tokenIdent
	tokenKeyword
	// e.g. "true nil"
	// checked after scanned tokenIdent
	tokenLiteral
	// e.g. "append"
	// checked after scanned tokenIdent
	tokenBuiltin // check from ident
	// e.g. "+ - = < / . ( }"
	// operators and punctuation
	tokenOperator
	// e.g. "//x /*x*/
	tokenComment
	// e.g. `"string" 'c'`
	// highlighted with apostrophys
	tokenString
	// e.g. "0x664"
	// TODO
	tokenNumber
	// e.g. "\t ' ' \n"
	tokenWS
	// EOF
	// returned if the file reached the end
	tokenEOF
	// invalid token
	tokenInvalid
)

type token struct {
	typ tokenType
	val string
}

func (t token) String() string {
	return fmt.Sprintf("token{%s, %q}", tokenTypeNames[t.typ], t.val)
}

var tokenTypeNames = map[tokenType]string{
	tokenKeyword:  "keywo",
	tokenLiteral:  "liter",
	tokenBuiltin:  "built",
	tokenOperator: "opera",
	tokenComment:  "comnt",
	tokenString:   "strng",
	tokenIdent:    "ident",
	tokenNumber:   "numbr",
	tokenEOF:      "__eof",
	tokenWS:       "whtsp",
	tokenInvalid:  "inval",
}

var eof = rune(0)
