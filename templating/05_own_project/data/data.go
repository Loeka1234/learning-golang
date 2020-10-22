package data

import (
	"encoding/json"
	"errors"
	"github.com/Loeka1234/learning-golang/templating/05_own_project/utils"
	"io/ioutil"
)

type Options struct {
	Color     string `json:"color"`
	BackColor string `json:"backColor"`
}

type Component struct {
	Selected string  `json:"selected"`
	Options  Options `json:"options"`
}

type JSONData struct {
	Header  Component `json:"header"`
	Section Component `json:"section"`
	Footer  Component `json:"footer"`
}

const (
	DataDirectory = "data"
	WebconfigPath = "/webconfig.json"
)

func GetWebconfig() JSONData {
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

func EditWebconfig(componentType string, newValue string) error {
	data := GetWebconfig()

	switch componentType {
	case "header":
		if !utils.Contains(utils.GetHeaders(), newValue) {
			return errors.New("component does not exists")
		}
		data.Header.Selected = newValue
	case "section":
		if !utils.Contains(utils.GetSections(), newValue) {
			return errors.New("component does not exists")
		}
		data.Section.Selected = newValue
	case "footer":
		if !utils.Contains(utils.GetFooters(), newValue) {
			return errors.New("component does not exists")
		}
		data.Footer.Selected = newValue
	default:
		return errors.New("component type not found in editWebconfig")
	}

	return saveData(data)
}

func EditOptions(componentType string, newOptions Options) error {
	data := GetWebconfig()
	switch componentType {
	case "header":
		data.Header.Options = newOptions
	case "section":
		data.Section.Options = newOptions
	case "footer":
		data.Footer.Options = newOptions
	default:
		return errors.New("invalid component type in editOption")
	}
	return saveData(data)
}

func saveData(data JSONData) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(DataDirectory+WebconfigPath, bytes, 0644)
	return err
}
