package link

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	content, err := ioutil.ReadFile("examples/ex1.html")
	if err != nil {
		panic(err)
	}
	parsed, err := Parse(strings.NewReader(string(content)))
	if err != nil {
		panic(err)
	}

	expectedResult := Link{"/other-page", "A link to another page"}

	if parsed[0] != expectedResult {
		t.Errorf("Parsed content was incorrect.")
	}
}
