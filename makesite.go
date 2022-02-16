package main

import (
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type HTMLcontents struct {
	Content string
}

func main() {
	parsedDir := flag.String("dir", ".", "Read all txt files in directory")
	flag.Parse()

	if *parsedDir != "" {
		println("Converting all txt files in " + *parsedDir + " to html")
		files, fileError := ioutil.ReadDir(*parsedDir)
		if fileError != nil {
			panic(fileError)
		}

		for _, file := range files {
			if file.Mode().IsRegular() {
				extension := filepath.Ext(file.Name())
				if extension == ".txt" {
					HTMLName := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name())) + ".html"
					fileContents, _ := ioutil.ReadFile(file.Name())

					templateStruct := HTMLcontents{Content: string(fileContents)}

					// Use a defined template
					parsedTemplate, _ := template.ParseFiles("template.tmpl")

					// Create a file to write to
					newFile, _ := os.Create(HTMLName)

					// Write to new file using template and data
					err := parsedTemplate.Execute(newFile, templateStruct)
					if err != nil {
						panic(err)
					}
				}
			}
		}
	} else {
		println("A directory was not passed, no HTML files have been generated.")
	}
}
