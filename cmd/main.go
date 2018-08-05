package main

import (
	"fmt"
	"os"

	"github.com/mbndr/lexy"
	"github.com/mbndr/lexy/lang"
)

const (
	tokenLimit = 200
)

func main() {
	f, _ := os.Open("example/example.go")

	l := lexy.NewLexer(f, lang.Go)

	for i := 0; i < tokenLimit; i++ {
		t := l.Scan()
		fmt.Println(t.String())

		if lexy.IsTokenEOF(t) {
			break
		}
	}

}
