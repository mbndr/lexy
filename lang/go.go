package lang

import "github.com/mbndr/lexy"

var Go lexy.Lang

func init() {
	Go.Keywords = lexy.NewWordList([]string{
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
	Go.Literals = lexy.NewWordList([]string{
		"true",
		"false",
		"iota",
		"nil",
	})
	Go.Builtins = lexy.NewWordList([]string{
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
	Go.Operators = "+-*/&=!()[]{}|<>^:.,;%"
}
