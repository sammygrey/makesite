package main

import (
	"flag"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/alexjohnj/caesar"
)

type HTMLcontents struct {
	Encrypted string
	Key       int
	Example   string
}

func main() {
	abc := "a b c d e f g h i j k l m n o p q r s t u v w x y z"
	parsedDir := flag.String("dir", ".", "Read all txt files in directory")
	flag.Parse()

	if *parsedDir != "" {
		println("Converting all txt files in " + *parsedDir + " to encrypted html")
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
					randKey := rand.Intn(26)
					templateStruct := HTMLcontents{Encrypted: caesar.EncryptPlaintext(string(fileContents), randKey), Key: randKey, Example: caesar.EncryptPlaintext(abc, randKey)}

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
