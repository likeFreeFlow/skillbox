package main

import (
	user "../../../31.5/pkg/user"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

func main() {
	log.Println("30.5 Friends\n------------")

	r := chi.NewRouter()

	r.Post("/", user.User{}.AddFriend)
	r.Get("/{id}", user.User{}.GetFriend)

	err := http.ListenAndServe(":8082", r)
	if err != nil {
		log.Printf("Ошибка запуска сервера\n%v", err)
	}
}
