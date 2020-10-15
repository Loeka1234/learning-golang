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

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/api/headers", headersHandler)
	http.HandleFunc("/api/sections", sectionsHandler)
	http.HandleFunc("/api/footers", footersHandler)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templating.Render(w)
}

func sendAsJSON(w http.ResponseWriter, i interface{}) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(i)
	if err != nil {
		panic(err)
	}
}

func headersHandler(w http.ResponseWriter, r *http.Request) {
	sendAsJSON(w, templating.GetHeaders())
}

func sectionsHandler(w http.ResponseWriter, r *http.Request) {
	sendAsJSON(w, templating.GetSections())
}

func footersHandler(w http.ResponseWriter, r *http.Request) {
	sendAsJSON(w, templating.GetFooters())
}
