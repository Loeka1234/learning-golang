package templating

import (
	"html/template"
	"net/http"
	"path/filepath"
)

const TemplatesDir = "templates"
const ComponentsDir = TemplatesDir + "/components"

// Returns all filenames of a specific foler in components
func getComponents(componentType string) []string {
	files, err := filepath.Glob(ComponentsDir + "/" + componentType + "/*.gohtml")
	if err != nil {
		panic(err)
	}
	return files
}

// Returns the filepath for index.gohtml and global-css.gohtml
func getDefaultFiles() []string {
	return []string{TemplatesDir + "/index.gohtml", TemplatesDir + "/global-css.gohtml"}
}

// Converts a slice of paths to a slice of filenames
func filePathsToFileName(filePaths []string) []string {
	for i, filePath := range filePaths {
		filePaths[i] = filepath.Base(filePath)
	}
	return filePaths
}

// Get all header filenames
func GetHeaders() []string {
	return filePathsToFileName(getComponents("headers"))
}

// Get all section filenames
func GetSections() []string {
	return filePathsToFileName(getComponents("sections"))
}

// Get all footers filenames
func GetFooters() []string {
	return filePathsToFileName(getComponents("footers"))
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

func Render(w http.ResponseWriter, header string, section string, footer string) {
	vw := ViewData{
		Header: Header{
			Path: nameToFolder("headers", header),
			Data: Data{
				BackColor: "#2e2e30",
				Color:     "white",
			},
		},
		Section: Section{
			Path: nameToFolder("sections", section),
			Data: Data{
				BackColor: "white",
				Color:     "black",
			},
		},
		Footer: Footer{
			Path: nameToFolder("footers", footer),
			Data: Data{
				BackColor: "#2e2e30",
				Color:     "white",
			},
		},
	}

	t, err := template.ParseFiles(
		getDefaultFiles()[0],
		getDefaultFiles()[1],
		vw.Header.Path,
		vw.Section.Path,
		vw.Footer.Path,
	)
	if err != nil {
		panic(err)
	}

	err = t.ExecuteTemplate(w, "index.gohtml", vw)
	if err != nil {
		panic(err)
	}
}
