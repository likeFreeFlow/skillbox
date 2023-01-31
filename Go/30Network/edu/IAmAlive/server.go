package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hey bruh"))
	})
	http.ListenAndServe("localhost:8080", mux)
}
