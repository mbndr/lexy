package main

import (
	"fmt"
	"log"
	"os"
	"flag"
	"html/template"
	"bytes"

	"github.com/mbndr/lexy"
	"github.com/mbndr/lexy/format"
	"github.com/mbndr/lexy/lang"
	"github.com/mbndr/lexy/style"
)

var (
	outputFile = flag.String("o", "example.html", "file to store html")
	inputFile = flag.String("i", "token.go", "file to lex and highlight")
	styleName = flag.String("s", "AtelierEstuaryLight", "style to use")
)

// generates a html file with highlighted code
// currently always with go language and one specific style
func main() {
	flag.Parse()

	s := style.ByName(*styleName)

	// input file
	src, err := os.Open(*inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	// output file
	dest, err := os.Create(*outputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer dest.Close()

	// scan tokens
	tokens := lexy.ScanAll(src, lang.Go)
	for i, t := range tokens {
		fmt.Printf("%d %s\n", i, t)
	}

	// write to template
	htmlBuf := new(bytes.Buffer)
	cssBuf := new(bytes.Buffer)

	formatter := format.NewHtmlFormatter(tokens)
	err = formatter.Format(htmlBuf)
	if err != nil {
		log.Fatal(err)
	}

	format.WriteCss(cssBuf, *s)

	t, err := template.New("").Parse(htmlTemplate)
	if err != nil {
		log.Fatal(err)
	}

	t.Execute(dest, struct{
		Style template.CSS
		Body template.HTML
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