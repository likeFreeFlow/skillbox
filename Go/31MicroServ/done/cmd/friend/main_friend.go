package main

import (
	"log"
	"net/http"

	user "main/pkg/user"

	"github.com/go-chi/chi/v5"
)

func main() {
	log.Println("30.5 Friends\n------------")

	r := chi.NewRouter()

	r.Post("/", user.Useru{}.AddFriend)
	r.Get("/{id}", user.Useru{}.GetFriend)

	err := http.ListenAndServe(":8082", r)
	if err != nil {
		log.Printf("Ошибка запуска сервера\n%v", err)
	}
}
