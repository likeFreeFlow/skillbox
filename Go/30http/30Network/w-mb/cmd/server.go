package main

import (
	"fmt"
	user "main/pkg/user"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	fmt.Println("30 блок, работа с http")
	fmt.Println("ищу практику на го, barakovvm@ya.ru")

	r := chi.NewRouter()

	r.Route("/user", func(r chi.Router) {
		r.Post("/create", user.User{}.Create)
		r.Post("/add_friend", user.User{}.AddFriend)
		r.Delete("/delete", user.User{}.Delete)
		r.Get("/friends/{id}", user.User{}.GetFriend)
		r.Put("/set_age/{id}", user.User{}.SetAge)
	})

	http.ListenAndServe(":8080", r)
	fmt.Println("Server is listening")
}
