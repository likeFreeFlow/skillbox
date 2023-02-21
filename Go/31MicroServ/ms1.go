package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

const addr1 string = "localhost:8080"

func main() {
	http.HandleFunc("/", handle)
	log.Fatalln(http.ListenAndServe(addr1, nil))
}

func handle(w http.ResponseWriter, r *http.Request) {

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Body.Close()

	text := string(bodyBytes)
	response := "1 instance = " + text
	if _, err := w.Write([]byte(response)); err != nil {
		log.Fatalln(err)
	}
}