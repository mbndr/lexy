# Lexy
[![GoDoc](https://godoc.org/github.com/mbndr/lexy?status.svg)](https://godoc.org/github.com/mbndr/lexy)
[![Go Report Card](https://goreportcard.com/badge/github.com/mbndr/lexy)](https://goreportcard.com/report/github.com/mbndr/lexy)

This is a lexer written primarily for learning purpose. My goal is to write a syntax highlighter with e.g. HTML output and other formatters. I also want the languages to be  more customizable (currently only Golang is supported). There may be significant API changes in the future!

Lexy currently is able to tokenize Golang code and highlight it with HtmlFormatter 

## Styles
Currently all styles are imported from Highlight.js.

Unfortunately Lexy can't handle the styles as detailed as Highlight.js and can't handle the styles with background image.

[Highlight.js LICENSE](https://github.com/highlightjs/highlight.js/blob/master/LICENSE)

## Usage
```go
import (
    "github.com/mbndr/lexy"
    "github.com/mbndr/lexy/format"
    "github.com/mbndr/lexy/lang"
    "github.com/mbndr/lexy/style"
)

// open the file to highlight
f, _ := os.Open("example.go")

// create io.Writers to get the html and css
htmlBuf = new(bytes.Buffer)
cssBuf = new(bytes.Buffer)

// get all tokens
tokens := lexy.ScanAll(f, lang.Go)

// create a formatter to get a highlighted code
formatter := format.NewHtml(style.AtelierEstuaryLight, htmlBuf, cssBuf)

// write to the buffers
err = formatter.Format(tokens)
if err != nil {
    log.Fatal(err)
}
```

## TODO
- image formatter
- static lexy.WriteCSS function
- remove IsTokenEOF()
- html preview page for all styles
- no need for TokenInvalid?
