// hljs2lexystyle converts highlightjs styles to lexy styles
// because hljs has more styling options and lexy only a few
// the exact styles are lost after conversion
//
// TODO download all files if no style given
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/aymerick/douceur/parser"
)

const (
	dlUrl = "https://raw.githubusercontent.com/highlightjs/highlight.js/master/src/styles/%s.css"
)

var (
	// name of the package of the output files
	packageName = flag.String("p", "style", "package of the output files")
	// name of the hljs file to download
	styleName = flag.String("s", "github-gist", "hljs style to download")
	// output file
	outputFile = flag.String("o", "", "output file for the generated lexy style")
	// rulemap
	//      .hljs      color  #ff9900
	r = map[string]map[string]string{}
)

func main() {
	flag.Parse()

	// get stylesheet
	resp, err := http.Get(fmt.Sprintf(dlUrl, *styleName))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// parse stylesheet
	stylesheet, err := parser.Parse(string(body))
	if err != nil {
		log.Fatal(err)
	}

	// process for easier access
	for _, rule := range stylesheet.Rules {
		for _, s := range rule.Selectors {
			r[s] = make(map[string]string)
			for _, d := range rule.Declarations {
				r[s][d.Property] = d.Value
			}
		}
	}

	// name of the exported style variable
	styleVarName := toStyleVarName(*styleName)

	if *outputFile == "" {
		*outputFile = styleVarName + ".go"
	}

	// build file
	funcMap := map[string]interface{}{
		"Property": getDeclarationValue,
	}

	t, err := template.New("").Funcs(funcMap).Parse(fileTemplate)
	if err != nil {
		log.Fatal(err)
	}

	outFile, err := os.Create(*outputFile)
	if err != nil {
		log.Fatal(err)
	}

	t.Execute(outFile, struct {
		Package string // TODO
		Name    string // Language name (how it's exported)
	}{
		Package: *packageName,
		Name:    styleVarName,
	})
}

func toStyleVarName(s string) string {
	s = strings.Title(s)
	s = strings.Replace(s, "-", "", -1)
	s = strings.Replace(s, "_", "", -1)
	return s
}

// getDeclarationValue returns a value of a selector property
// checks if '"' are needed and is flexible with background vs background-color
// because we check if the property is "background", in the template must be searched with "background" and not "background-color"
func getDeclarationValue(selector, property string) string {
	val := r[selector][property]

	if val == "" && property == "background" {
		val = r[selector]["background-color"]
	}

	if val == "" {
		log.Printf("empty value: %s -> %s", selector, property)
	}

	// TODO default color if not specified or handles that css later on itself?

	return val
}

var fileTemplate = `package {{.Package}}

import "github.com/mbndr/lexy"

var {{.Name}} = lexy.Style{
	Foreground: "{{Property ".hljs" "color"}}",
	Background: "{{Property ".hljs" "background"}}",

	TokenColors: map[lexy.TokenType]string{
		lexy.TokenKeyword:  "{{Property ".hljs-keyword" "color"}}",
		lexy.TokenLiteral:  "{{Property ".hljs-literal" "color"}}",
		lexy.TokenBuiltin:  "{{Property ".hljs-built_in" "color"}}",
		lexy.TokenOperator: "inherit", // TODO
		lexy.TokenComment:  "{{Property ".hljs-comment" "color"}}",
		lexy.TokenString:   "{{Property ".hljs-string" "color"}}",
		lexy.TokenNumber:   "{{Property ".hljs-number" "color"}}",
	},
}
`
