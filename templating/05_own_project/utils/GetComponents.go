package utils

import "path/filepath"

const ComponentsDir = "templates/components"

// Returns all filenames of a specific foler in components
func GetComponents(componentType string) []string {
	files, err := filepath.Glob(ComponentsDir + "/" + componentType + "/*.gohtml")
	if err != nil {
		panic(err)
	}
	return files
}

// Get all header filenames
func GetHeaders() []string {
	return FilePathsToFileNames(GetComponents("headers"))
}

// Get all section filenames
func GetSections() []string {
	return FilePathsToFileNames(GetComponents("sections"))
}

// Get all footers filenames
func GetFooters() []string {
	return FilePathsToFileNames(GetComponents("footers"))
}
