package main

import (
	"log"
	"net/http"

	user "main/user"

	chi "github.com/go-chi/chi/v5"
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

	/* r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello GolinuxCloud numbers!"))
	}) */
	err := http.ListenAndServe(":8080", r)
	if err != nil {

		log.Printf("Ошибка запуска сервера\n%v", err)
	}
}
