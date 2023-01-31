package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func main() {

	message := map[string]interface{}{
		"name": "Viktor", "age": "24", "friends": ""}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("https:localhost:8080", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	log.Println(result)
	log.Println(result["data"])
}
