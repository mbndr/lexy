package lexy

// Style represents a color scheme for a language
type Style struct {
	Foreground string
	Background string

	TokenColors map[TokenType]string // TODO image/color instead of string?
}
