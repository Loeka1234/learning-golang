package main

import (
	"html/template"
	"net/http"
)

var testTemplate *template.Template

type User struct {
	ID            int
	Email         string
	HasPermission func(string) bool
}

type ViewData struct {
	User User
}

//func (u User) HasPermission(feature string) bool {
//	if feature == "feature-a" {
//		return true
//	} else {
//		return false
//	}
//}

func main() {
	var err error
	testTemplate, err = template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	vd := ViewData{
		User: User{
			ID:    1,
			Email: "loeka@loeka.com",
			HasPermission: func(feature string) bool {
				if feature == "feature-b" {
					return true
				}
				return false
			},
		},
	}
	err := testTemplate.Execute(w, vd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
