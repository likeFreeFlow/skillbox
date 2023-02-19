package main

import (
	"fmt"
	"log"
	user "main/pkg/user"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	log.Println("31.5 Users\n------------")

	r := chi.NewRouter()

	r.Post("/", user.Useru{}.Create)
	r.Delete("/", user.Useru{}.Delete)
	r.Put("/{id}", user.Useru{}.SetAge)

	err := http.ListenAndServe(":8081", r)
	fmt.Println("Server is lstening")
	if err != nil {
		log.Printf("Ошибка запуска сервера\n%v", err)
	}
}
