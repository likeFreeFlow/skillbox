package main

import (
	user "../../../31.5/pkg/user"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

func main() {
	log.Println("31.5 Users\n------------")

	r := chi.NewRouter()

	r.Post("/", user.User{}.Create)
	r.Delete("/", user.User{}.Delete)
	r.Put("/{id}", user.User{}.SetAge)

	err := http.ListenAndServe(":8081", r)
	if err != nil {
		log.Printf("Ошибка запуска сервера\n%v", err)
	}
}
