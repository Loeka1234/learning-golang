package main

import (
	"fmt"
	"github.com/Loeka1234/learning-golang/templating/04_an_application/views"
	"net/http"
)

var index *views.View
var contact *views.View

func main() {
	fmt.Println("--- Index ---")
	index = views.NewView("bootstrap", "views/index.gohtml")
	fmt.Println("--- Contact ---")
	contact = views.NewView("bootstrap", "views/contact.gohtml")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/contact", contactHandler)
	http.ListenAndServe(":3000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := index.Render(w, nil)
	if err != nil {
		panic(err)
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	err := contact.Render(w, nil)
	if err != nil {
		panic(err)
	}
}
