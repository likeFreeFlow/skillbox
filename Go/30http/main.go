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

/*
30.5 Практическая работа
Цель практической работы
Научиться:

работать с запросами POST, GET, PUT, DELETE;
применять принципы написания обработчиков HTTP-запросов.


Что нужно сделать
Напишите HTTP-сервис, который принимает входящие соединения с JSON-данными и обрабатывает их следующим образом:

    1. Сделайте обработчик создания пользователя. У пользователя должны быть следующие поля: имя, возраст и массив друзей. Пользователя необходимо сохранять в мапу. Пример запроса:

POST /create HTTP/1.1
Content-Type: application/json; charset=utf-8
Host: localhost:8080
{"name":"some name","age":"24","friends":[]}
Данный запрос должен возвращать ID пользователя и статус 201.



    2. Сделайте обработчик, который делает друзей из двух пользователей. Например, если мы создали двух пользователей и нам вернулись их ID, то в запросе мы можем указать ID пользователя, который инициировал запрос на дружбу, и ID пользователя, который примет инициатора в друзья. Пример запроса:

    POST /make_friends HTTP/1.1
Content-Type: application/json; charset=utf-8
Host: localhost:8080
{"source_id":"1","target_id":"2"}
Данный запрос должен возвращать статус 200 и сообщение «username_1 и username_2 теперь друзья».



    3. Сделайте обработчик, который удаляет пользователя. Данный обработчик принимает ID пользователя и удаляет его из хранилища, а также стирает его из массива friends у всех его друзей. Пример запроса:

    DELETE /user HTTP/1.1
Content-Type: application/json; charset=utf-8
Host: localhost:8080
{"target_id":"1"}
Данный запрос должен возвращать 200 и имя удалённого пользователя.



    4. Сделайте обработчик, который возвращает всех друзей пользователя. Пример запроса:

    GET /friends/user_id HTTP/1.1
Host: localhost:8080
Connection: close
После /friends/ указывается id пользователя, друзей которого мы хотим увидеть.



    5. Сделайте обработчик, который обновляет возраст пользователя. Пример запроса:

PUT /user_id HTTP/1.1
Content-Type: application/json; charset=utf-8
Host: localhost:8080
{"new age":"28"}
Запрос должен возвращать 200 и сообщение «возраст пользователя успешно обновлён».



Советы и рекомендации
Воспользуйтесь библиотекой Chi.*/
