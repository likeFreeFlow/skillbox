package handler

import (
	"log"
	"net/http"
)

func handler(w http.Response, r *http.Request) {
	text := "message from testhandler"
	if _, err := w.Write([]byte(text)); err != nil {
		log.Fatalln(err)
	}
}
