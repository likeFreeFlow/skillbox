package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hey bruh"))
	})
	http.ListenAndServe("localhost:8080", mux)
}
func (s *service) Create(whttp.ResponseWriter,r *http.Request){
	if r.Method=="POST"{
		content,err:=ioutil.ReadAll(r.Body)
		if err !=nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))

		}
		splittedContent:=strings.Split(string(content),"")
		s.store[splittedContent[0]]=string(content)

		w.WriteHeader(http.StatusCreated)
		w.Write("user was created "+splittedContent[0])
		return
	}
	w.WriteHeader(http.StatusRequest)
}
func(s*service) GetAll(w http ResponseWriter,r*http.Request){
	if r.Method=="GET"{
		response:=""
		for _,user:=range s.store{
			response+=user+"\n"

		}
		w.WriteHeader(http.StatusOk)
		w.Write([]byte(response))
		returh
	}
	w.WreteHeader(http.StatusBadRequest) 
}