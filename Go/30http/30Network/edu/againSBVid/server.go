package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type service struct {
	store map[string]string
}

func (s *service) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		splittedContent := strings.Split(string(content), " ")
		s.store[splittedContent[0]] = string(content)

		w.WriteHeader
	}
	w.WriteHeader(http.StatusBadRequest)
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("I'm alive"))
	})
	fmt.Println("Server is listening")
	http.ListenAndServe("localhost:8080", mux)
}
