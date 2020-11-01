package main

import (
	"flag"
	"fmt"
	cyao "github.com/Loeka1234/learning-golang/web_interactive_book"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the CYAO web application on")
	fileName := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *fileName)

	f, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}

	story, err := cyao.JsonStory(f)
	if err != nil {
		panic(err)
	}

	tpl := template.Must(template.New("").Parse(storyTmpl))

	h := cyao.NewHandler(story,
		cyao.WithTemplate(tpl),
		cyao.WithPathFunc(pathFn),
	)
	mux := http.NewServeMux()
	mux.Handle("/story/", h)
	fmt.Printf("Starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}

func pathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}
	if strings.Contains(path, "/story/") {
		return path[len("/story/"):]
	}
	return path
}

var storyTmpl = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
</head>
<body>
<h1>{{.Title}}</h1>
{{range .Story}}
    <p>{{.}}</p>
{{end}}
<ul>
    {{range .Options}}
    <li>
        <a href="/story/{{.Chapter}}">{{.Text}}</a>
    </li>
    {{end}}
</ul>
</body>
</html>`
