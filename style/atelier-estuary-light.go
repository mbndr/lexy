package style

import "github.com/mbndr/lexy"

var AtelierEstuaryLight = lexy.Style{
	Foreground: "#5f5e4e",
	Background: "#f4f3ec",

	TokenColors: map[lexy.TokenType]string{
		lexy.TokenKeyword:  "#5f9182",
		lexy.TokenLiteral:  "#ae7313",
		lexy.TokenBuiltin:  "#ae7313",
		lexy.TokenOperator: "#5f5e4e",
		lexy.TokenComment:  "#6c6b5a",
		lexy.TokenString:   "#7d9726",
		lexy.TokenNumber:   "#ba6236",
	},
}
