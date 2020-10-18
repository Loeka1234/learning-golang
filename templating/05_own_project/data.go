package main

import (
	"encoding/json"
	"errors"
	templating "github.com/Loeka1234/learning-golang/templating/05_own_project/templates"
	"io/ioutil"
)

type JSONData struct {
	Header  string
	Section string
	Footer  string
}

const (
	DataDirectory = "data"
	WebconfigPath = "/webconfig.json"
)

func getWebconfig() JSONData {
	data, err := ioutil.ReadFile(DataDirectory + WebconfigPath)
	if err != nil {
		panic(err)
	}

	var toReturn JSONData

	err = json.Unmarshal(data, &toReturn)
	if err != nil {
		panic(err)
	}

	return toReturn
}

func editWebconfig(componentType string, newValue string) error {
	data := getWebconfig()

	switch componentType {
	case "header":
		if !Contains(templating.GetHeaders(), newValue) {
			return errors.New("component does not exists")
		}
		data.Header = newValue
	case "section":
		if !Contains(templating.GetSections(), newValue) {
			return errors.New("component does not exists")
		}
		data.Section = newValue
	case "footer":
		if !Contains(templating.GetFooters(), newValue) {
			return errors.New("component does not exists")
		}
		data.Footer = newValue
	default:
		panic(errors.New("component type not found in editWebconfig"))
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(DataDirectory + WebconfigPath, bytes, 0644)
	return err
}
