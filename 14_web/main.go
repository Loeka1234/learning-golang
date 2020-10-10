package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "About me!")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Server Starting...")
	http.ListenAndServe(":3000", nil)
}
