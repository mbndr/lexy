package lexy

// WordList is a map for better access to search a keyword etc
type WordList map[string]bool

// NewWordList returns a new WordList
func NewWordList(s []string) WordList {
	wl := make(WordList)

	for _, w := range s {
		wl[w] = true
	}

	return wl
}

// Lang represents the data of a language
type Lang struct {
	// TODO comment indicators
	Keywords  WordList
	Literals  WordList
	Builtins  WordList
	Operators string
}
