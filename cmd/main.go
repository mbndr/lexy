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

const (
	tokenLimit = 200
)

func main() {
	src, _ := os.Open("lexer.go")

	var htmlBuf *bytes.Buffer = new(bytes.Buffer)
	var cssBuf *bytes.Buffer = new(bytes.Buffer)

	tokens := lexy.ScanAll(src, lang.Go)
	for i, t := range tokens {
		fmt.Printf("%d %s\n", i, t)
	}

	formatter := format.NewHtmlFormatter(tokens)

	err := formatter.Format(htmlBuf)
	if err != nil {
		log.Fatal(err)
	}

	format.WriteCss(cssBuf, style.AtelierEstuaryLight)

	// write complete html
	f, _ := os.Create("example/example.html")
	t, err := template.New("").Parse(html)
	if err != nil {
		log.Fatal(err)
	}

	t.Execute(f, struct {
		Style template.CSS
		Body  template.HTML
	}{
		template.CSS(cssBuf.String()),
		template.HTML(htmlBuf.String()),
	})

}

var html = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>Page Title</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
	<style>
	body {margin: 0; padding: 0;}
	{{.Style}}
	</style>
</head>
<body>
    {{.Body}}
</body>
</html>
`
