package utils

import "path/filepath"

// Converts a slice of paths to a slice of filenames
func FilePathsToFileNames(filePaths []string) []string {
	for i, filePath := range filePaths {
		filePaths[i] = filepath.Base(filePath)
	}
	return filePaths
}
