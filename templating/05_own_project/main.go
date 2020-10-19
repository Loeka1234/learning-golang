package main

import (
	"encoding/json"
	"fmt"
	"github.com/Loeka1234/learning-golang/templating/05_own_project/data"
	"github.com/Loeka1234/learning-golang/templating/05_own_project/templates"
	"github.com/Loeka1234/learning-golang/templating/05_own_project/utils"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	fmt.Println(utils.GetHeaders())
	fmt.Println(utils.GetSections())
	fmt.Println(utils.GetFooters())

	http.HandleFunc("/", handler)
	http.Handle("/dashboard", http.FileServer(http.Dir("./dashboard/build")))

	PORT := os.Getenv("PORT")
	fmt.Println("PORT: ", PORT)
	err = http.ListenAndServe(":" + PORT, nil)
	if err != nil {
		panic(err)
	}
}

type ReturnError struct {
	Error string
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	// Website
	if r.URL.Path == "/" {
		indexHandler(w, r)
	}

	// API routes
	switch r.Method {
	case "GET":
		switch r.URL.Path {
		case "/api/headers":
			sendAsJSON(w, utils.GetHeaders(), http.StatusOK)
		case "/api/sections":
			sendAsJSON(w, utils.GetSections(), http.StatusOK)
		case "/api/footers":
			sendAsJSON(w, utils.GetFooters(), http.StatusOK)
		case "/api/webconfig":
			sendAsJSON(w, data.GetWebconfig(), http.StatusOK)
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
			err := data.EditWebconfig(componentType[0], component[0])
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
