package main

import (
	"fmt"
	"log"
	user "main/pkg/user"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	log.Println("30.5\n------------")

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
	/* 	err :=http.ListenAndServe(":8080", r)
	   	if err != nil {
	   		log.Printf("Ошибка запуска сервера\n%v", err)
	   	} else {
	   		fmt.Println("Server is running")
	   	}
	*/
}
