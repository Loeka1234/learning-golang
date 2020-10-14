package main

import (
	"fmt"
	"github.com/Loeka1234/learning-golang/templating/05_own_project/templates"
	"net/http"
)

func main() {
	fmt.Println(templating.GetHeaders())
	fmt.Println(templating.GetSections())
	fmt.Println(templating.GetFooters())

	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templating.Render(w)
}
