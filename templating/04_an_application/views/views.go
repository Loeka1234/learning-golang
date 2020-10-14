package views

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

const LayoutDir = "views/layouts"

func NewView(layout string, files ...string) *View {
	files = append(layoutFiles(), files...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	fmt.Println("Files in View.Template: ", files)

	return &View{
		Template: t,
		Layout:   layout,
	}
}

type View struct {
	Template *template.Template
	Layout   string
}

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "/*.gohtml")
	if err != nil {
		panic(err)
	}
	fmt.Println(files)
	return files
}

var flashRotator int = 0

func flashes() map[string]string {
	flashRotator++
	if flashRotator%3 == 0 {
		return map[string]string{
			"warning": "You are about to exceed your plan limits!",
		}
	} else {
		return map[string]string{}
	}
}

type ViewData struct {
	Flashes map[string]string
	Data    interface{}
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	vd := ViewData{
		Flashes: flashes(),
		Data: data,
	}
	fmt.Println("ViewData: ", vd)
	fmt.Println("Layout: ", v.Layout)
	return v.Template.ExecuteTemplate(w, v.Layout, vd)
}
