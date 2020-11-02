package main

import (
	"fmt"
	link "github.com/Loeka1234/learning-golang/link_parser"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func main() {
	files, _ := filepath.Glob("examples/*.html")
	for i, file := range files {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}

		fmt.Printf("-- Starting test %v --\n", i + 1)
		links, err := link.Parse(strings.NewReader(string(content)))
		if err != nil {
			panic(err)
		}
		for _, l := range links {
			fmt.Println(l.Href + ": " + l.Text)
		}
		fmt.Println("")
	}
}
