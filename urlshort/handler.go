package urlshort

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
	"net/http"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

func YAMLHandler(yaml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathUrls, err := parseYaml(yaml)
	if err != nil {
		return nil, err
	}
	pathToUrls := buildMap(pathUrls)
	return MapHandler(pathToUrls, fallback), nil
}

func JSONHandler(json []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathUrls, err := parseJSON(json)
	if err != nil {
		return nil, err
	}
	pathToUrls := buildMap(pathUrls)
	return MapHandler(pathToUrls, fallback), nil
}

func MySQLHandler(connectionString string, fallback http.Handler) (http.HandlerFunc, error) {
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	results, err := db.Query("SELECT path, url FROM golang_test_urlshortener;")
	if err != nil {
		return nil, err
	}

	pathToUrls := make(map[string]string)
	for results.Next() {
		var (
			path string
			url  string
		)
		err := results.Scan(&path, &url)
		if err != nil {
			return nil, err
		}

		pathToUrls[path] = url
	}

	return MapHandler(pathToUrls, fallback), nil
}

func buildMap(pathUrls []pathUrl) map[string]string {
	returnValue := make(map[string]string)
	for _, pathUrl := range pathUrls {
		returnValue[pathUrl.Path] = pathUrl.URL
	}
	return returnValue
}

func parseYaml(data []byte) ([]pathUrl, error) {
	var pathUrls []pathUrl
	err := yaml.Unmarshal(data, &pathUrls)
	if err != nil {
		return nil, err
	}
	return pathUrls, nil
}

func parseJSON(data []byte) ([]pathUrl, error) {
	var pathUrls []pathUrl
	err := yaml.Unmarshal(data, &pathUrls)
	if err != nil {
		return nil, err
	}
	return pathUrls, nil
}

type pathUrl struct {
	Path string `yaml:"path" json:"path"`
	URL  string `yaml:"url" json:"url"`
}
