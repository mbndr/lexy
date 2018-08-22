package main

// NOT COMPLETE / WORKING
/*
import (
	"fmt"
	"log"
	"os"

	"github.com/mbndr/lexy"
	"github.com/mbndr/lexy/format"
	"github.com/mbndr/lexy/lang"
	"github.com/mbndr/lexy/style"
)

var (
	s = style.ByName("Zenburn")
	inputFile = "lexer.go"
	outputFile = "example.png"
)

// generates a png file with colored lines
func main() {

	// input file
	src, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	// output file
	dest, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer dest.Close()

	// scan tokens (and print them)
	tokens := lexy.ScanAll(src, lang.Go)
	for i, t := range tokens {
		fmt.Printf("%d %s\n", i, t)
	}

	// create formatter
	formatter := format.NewLineImage(*s, dest)
	err = formatter.Format(tokens)
	if err != nil {
		log.Fatal(err)
	}

}
*/