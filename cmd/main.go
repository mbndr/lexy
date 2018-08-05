package main

import (
	"fmt"
	"os"
	"log"

	"github.com/mbndr/lexy"
	"github.com/mbndr/lexy/format"
	"github.com/mbndr/lexy/lang"
)

const (
	tokenLimit = 200
)

func main() {
	f, _ := os.Open("example/example.go")

	/*l := lexy.NewLexer(f, lang.Go)

	for i := 0; i < tokenLimit; i++ {
		t := l.Scan()
		fmt.Println(t.String())

		if lexy.IsTokenEOF(t) {
			break
		}
	}*/

	tokens := lexy.ScanAll(f, lang.Go)
	for i, t := range tokens {
		fmt.Printf("%d %s\n", i, t)
	}

	formatter := format.NewHtmlFormatter(tokens)

	err := formatter.Format(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}

}
