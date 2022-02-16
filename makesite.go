package main

import (
	"flag"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

type HTMLcontents struct {
	Content string
}

func main() {
	fileInput := flag.String("file", "first-post.txt", "txt file to pass in")
	flag.Parse()

	newHTML := strings.Split(*fileInput, ".")[0] + ".html"

	fileContents, _ := ioutil.ReadFile(*fileInput)

	templateStruct := HTMLcontents{Content: string(fileContents)}

	parsedTemplate, _ := template.ParseFiles("template.tmpl")

	newFile, _ := os.Create(newHTML)

	err := parsedTemplate.Execute(newFile, templateStruct)
	if err != nil {
		panic(err)
	}
}
