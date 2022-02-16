package main

import (
	"io/ioutil"
	"os"
	"text/template"
)

type HTMLtemplate struct {
	Content string
}

func main() {
	fileContents, _ := ioutil.ReadFile("./first-post.txt")

	templateStruct := HTMLtemplate{Content: string(fileContents)}

	parsedTemplate, _ := template.ParseFiles("template.tmpl")

	newFile, _ := os.Create("first-post.html")

	err := parsedTemplate.Execute(newFile, templateStruct)

	if err != nil {
		panic(err)
	}
}
