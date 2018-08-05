package lang

// TODO in seperate file
type WordList map[string]bool

func NewWordList(s []string) WordList {
	wl := make(WordList)

	for _, w := range s {
		wl[w] = true
	}

	return wl
}

type Lang struct {
	// TODO comment indicators
	Keywords WordList
	Literals WordList
	Builtins WordList
	Operators string // TODO how to implement?
}

// END TODO

var Go Lang

func init() {
	Go.Keywords = NewWordList([]string{
		"break",
		"default",
		"func",
		"interface",
		"select",
		"case",
		"defer",
		"go",
		"map",
		"struct",
		"chan",
		"else",
		"goto",
		"package",
		"switch",
		"const",
		"fallthrough",
		"if",
		"range",
		"type",
		"continue",
		"for",
		"import",
		"return",
		"var",
		// types
		"bool",
		"byte",
		"complex64",
		"complex128",
		"float32",
		"float64",
		"int8",
		"int16",
		"int32",
		"int64",
		"string",
		"uint8",
		"uint16",
		"uint32",
		"uint64",
		"int",
		"uint",
		"uintptr",
		"rune",
		"error",
	})
	Go.Literals = NewWordList([]string{
		"true",
		"false",
		"iota",
		"nil",
	})
	Go.Operators = "+-*/&=!()[]{}|<>^:.,;%"
	/*
	Go.Operators = NewWordList([]string{
		"+",
		"&",
		"+=",
		"&=",
		"&&",
		"==",
		"!=",
		"(",
		")",
		"-",
		"|",
		"-=",
		"|=",
		"||",
		"<",
		"<=",
		"[",
		"]",
		"*",
		"^",
		"*=",
		"^=",
		"<-",
		">",
		">=",
		"{",
		"}",
		"/",
		"<<",
		"/=",
		"<<=",
		"++",
		"=",
		":=",
		",",
		";",
		"%",
		">>",
		"%=",
		">>=",
		"--",
		"!",
		"...",
		".",
		":",
		"&^",
		"&^=",
	})
	*/
	Go.Builtins = NewWordList([]string{
		"append",
		"cap",
		"close",
		"complex",
		"copy",
		"imag",
		"len",
		"make",
		"new",
		"panic",
		"print",
		"println",
		"real",
		"recover",
		"delete",
	})
}
