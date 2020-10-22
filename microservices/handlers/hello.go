package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello world!")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Oops, something went wrong.", http.StatusBadRequest)
		return
	}

	_, _ = fmt.Fprintf(w, "Hello %s", d)
}