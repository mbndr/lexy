package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mbndr/lexy"
	"github.com/mbndr/lexy/format"
	"github.com/mbndr/lexy/lang"
	"github.com/mbndr/lexy/style"
)

// prints all tokens, the html and css
func main() {
	if len(os.Args) < 2 {
		log.Fatal("no input file given")
	}

	src, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(os.Stdout, "\nTokens:\n")
	tokens := lexy.ScanAll(src, lang.Go)
	for i, t := range tokens {
		fmt.Printf("%d %s\n", i, t)
	}

	formatter := format.NewHtmlFormatter(tokens)

	fmt.Fprintln(os.Stdout, "\nHTML:\n\n")
	err = formatter.Format(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(os.Stdout, "\nCSS:\n")
	format.WriteCss(os.Stdout, style.AtelierEstuaryLight)

}
