package templating

import (
	"github.com/Loeka1234/learning-golang/templating/05_own_project/data"
	"html/template"
	"net/http"
)

const TemplatesDir = "templates"

// Returns the filepath for index.gohtml and global-css.gohtml
func getDefaultFiles() []string {
	return []string{TemplatesDir + "/index.gohtml", TemplatesDir + "/global-css.gohtml"}
}

// Converts component type + name to full path
func nameToFolder(componentType string, name string) string {
	return "templates/components/" + componentType + "/" + name
}

type Data struct {
	BackColor string
	Color     string
}

type Header struct {
	Path string
	Data Data
}

type Section struct {
	Path string
	Data Data
}

type Footer struct {
	Path string
	Data Data
}

type ViewData struct {
	Header  Header
	Section Section
	Footer  Footer
}

func Render(w http.ResponseWriter) {
	conf := data.GetWebconfig()

	t, err := template.ParseFiles(
		getDefaultFiles()[0],
		getDefaultFiles()[1],
		nameToFolder("headers", conf.Header.Selected),
		nameToFolder("sections", conf.Section.Selected),
		nameToFolder("footers", conf.Footer.Selected),
	)
	if err != nil {
		panic(err)
	}

	err = t.ExecuteTemplate(w, "index.gohtml", conf)
	if err != nil {
		panic(err)
	}
}
