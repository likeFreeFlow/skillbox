package main

import (
	"fmt"
	"net/http"
)

type msg string

func (m msg) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprint(resp, m)
}

func main() {
	msgHandler := msg("Hello from Web Server in Go")
	msgHandler2 := msg("ssssssssss 2 строка")
	fmt.Println("Server is listening...")
	http.ListenAndServe("localhost:8080", msgHandler)
	http.ListenAndServe("localhost:8080", msgHandler2)
}
