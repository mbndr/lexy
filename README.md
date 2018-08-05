# Lexy
[![GoDoc](https://godoc.org/github.com/mbndr/lexy?status.svg)](https://godoc.org/github.com/mbndr/lexy)
[![Go Report Card](https://goreportcard.com/badge/github.com/mbndr/lexy)](https://goreportcard.com/report/github.com/mbndr/lexy)

This is a lexer written just for learning purose. My goal is it to write a syntax highlighter with e.g. HTML output. I also want to be the languages more customizable. Currently it only supports Golang. There will be significant API changes in the future!

## Styles
Currently all styles are imported from Highlight.js.

Unfortunately lexy can't handle the styles as detailed as Highlight.js and can't handle the styles with background image.

## Usage
```go
import (
    "github.com/mbndr/lexy"
    "github.com/mbndr/lexy/format"
    "github.com/mbndr/lexy/lang"
    "github.com/mbndr/lexy/style"
)

f, _ := os.Open("example.go")
htmlBuf = new(bytes.Buffer)
cssBuf = new(bytes.Buffer)

// returns a slice of all Tokens
tokens := lexy.ScanAll(f, lang.Go)

// get the html
formatter := format.NewHtmlFormatter(tokens)
err := formatter.Format(htmlBuf)

// get the css
format.WriteCss(cssBuf, style.AtelierEstuaryLight)
```