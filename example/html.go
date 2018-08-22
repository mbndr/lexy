package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/mbndr/lexy"
	"github.com/mbndr/lexy/format"
	"github.com/mbndr/lexy/lang"
	"github.com/mbndr/lexy/style"
)

var (
	s = style.ByName("AtelierEstuaryLight")
	inputFile = "token.go"
	outputFile = "example.html"
)

// generate a html file with highlighted code
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

	// buffers to store generated code
	htmlBuf := new(bytes.Buffer)
	cssBuf := new(bytes.Buffer)

	// create formatter
	formatter := format.NewHtml(*s, htmlBuf, cssBuf)
	err = formatter.Format(tokens)
	if err != nil {
		log.Fatal(err)
	}

	// write to the template
	t, err := template.New("").Parse(htmlTemplate)
	if err != nil {
		log.Fatal(err)
	}

	t.Execute(dest, struct {
		Style template.CSS
		Body  template.HTML
	}{
		template.CSS(cssBuf.String()),
		template.HTML(htmlBuf.String()),
	})

}

var htmlTemplate = `<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>Lexy example</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <style>
    {{.Style}}
    </style>
</head>
<body>
    {{.Body}}
</body>
</html>`
