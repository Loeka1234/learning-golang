package main

import (
	"flag"
	"fmt"
	"github.com/Loeka1234/learning-golang/urlshort"
	"io/ioutil"
	"net/http"
)

func main() {
	YAMLFileLocation := flag.String("yaml", "routes.yaml", "use a yaml file for your routes")
	JSONFileLocation := flag.String("json", "routes.json", "use a json file for your routes")
	MySQLConnectionString := flag.String("sql", "", "use a mySQL database table")
	flag.Parse()

	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/loeka":    "https://www.loeka.me",
		"/artoeroe": "https://prototype.xiler.net",
	}
	handler := urlshort.MapHandler(pathsToUrls, mux)

	yaml, err := ioutil.ReadFile(*YAMLFileLocation)
	if err == nil {
		handler, err = urlshort.YAMLHandler([]byte(yaml), handler)
		if err != nil {
			panic(err)
		}
	}

	json, err := ioutil.ReadFile(*JSONFileLocation)
	if err == nil {
		handler, err = urlshort.JSONHandler([]byte(json), handler)
		if err != nil {
			panic(err)
		}
	}

	if *MySQLConnectionString != "" {
		handler, err = urlshort.MySQLHandler(*MySQLConnectionString, handler)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Starting the server on :8080")
	panic(http.ListenAndServe(":8080", handler))
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
