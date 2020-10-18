package main

import (
	"encoding/json"
	"fmt"
	"github.com/Loeka1234/learning-golang/templating/05_own_project/templates"
	"net/http"
)

func main() {
	fmt.Println(templating.GetHeaders())
	fmt.Println(templating.GetSections())
	fmt.Println(templating.GetFooters())

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}

type ReturnError struct {
	Error string
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Website
	if r.URL.Path == "/" {
		indexHandler(w, r)
	}

	// API routes
	switch r.Method {
	case "GET":
		switch r.URL.Path {
		case "/api/headers":
			sendAsJSON(w, templating.GetHeaders(), http.StatusOK)
		case "/api/sections":
			sendAsJSON(w, templating.GetSections(), http.StatusOK)
		case "/api/footers":
			sendAsJSON(w, templating.GetFooters(), http.StatusOK)
		case "/api/webconfig":
			sendAsJSON(w, getWebconfig(), http.StatusOK)
		}
	case "POST":
		switch r.URL.Path {
		case "/api/edit":
			// Getting query strings
			componentType, ok := r.URL.Query()["type"]
			component, _ok := r.URL.Query()["comp"]

			fmt.Println("Component: ", componentType)

			// Checking for correct input
			if !ok || !_ok || (componentType[0] != "header" && componentType[0] != "section" && componentType[0] != "footer") {
				sendAsJSON(w, ReturnError{"Query string 'type' or 'comp' incorrect or missing."}, http.StatusInternalServerError)
				return
			}

			// Changing in webconfig
			err := editWebconfig(componentType[0], component[0])
			if err != nil {
				sendAsJSON(w, ReturnError{err.Error()}, http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
		}
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templating.Render(w)
}

func sendAsJSON(w http.ResponseWriter, i interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(i)
	if err != nil {
		panic(err)
	}
}
