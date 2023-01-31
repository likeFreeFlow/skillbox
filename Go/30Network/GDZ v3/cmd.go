package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	MakeRequest()
}

func MakeRequest() {
	data := []byte(`{"Corado","24"}`)
	r := bytes.NewReader(data)
	resp, err := http.Post("http://localhost:8080/create", "application/json", r)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)

	/* body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body)) */
}
