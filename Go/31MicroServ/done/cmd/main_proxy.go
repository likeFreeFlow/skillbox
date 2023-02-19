package main

import (
	"io"
	"log"
	"net/http"
	"strings"
)

const (
	proxyHost  = ":8080"
	userHost   = "http://localhost:8081"
	friendHost = "http://localhost:8082"
)

func main() {
	log.Println("31.5 Proxy\n-----------")
	log.Println("Ищу практику в интересных проектах на го. barakovvm@ya.ru ")

	mux := http.NewServeMux()

	mux.HandleFunc("/user/", userHandler)
	mux.HandleFunc("/friend/", friendHandler)

	err := http.ListenAndServe(proxyHost, mux)
	if err != nil {
		log.Printf("Ошибка запуска сервера\n%v", err)
	}
}

func getURL(domain, uri string) string {
	buf := strings.Split(uri, "/")
	buf[1] = domain
	url := strings.Join(buf, "/")
	return strings.TrimPrefix(url, "/")
}

func getRequest(domain string, r *http.Request) ([]byte, error) {
	content := []byte("")

	target := getURL(domain, r.RequestURI)
	newReq, err := http.NewRequest(r.Method, target, r.Body)
	if err != nil {
		return content, err
	}
	newReq.Header.Add("Content-Type", r.Header.Get("Content-Type"))

	client := &http.Client{}
	newResp, err := client.Do(newReq)
	if err != nil {
		return content, err
	}
	content, err = io.ReadAll(newResp.Body)
	if err != nil {
		return content, err
	}

	return content, nil
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	content, err := getRequest(userHost, r)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(content)
}

func friendHandler(w http.ResponseWriter, r *http.Request) {
	content, err := getRequest(friendHost, r)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(content)
	w.WriteHeader(http.StatusOK)
}
