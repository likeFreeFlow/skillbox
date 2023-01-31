package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u *User) toString() string {
	return fmt.Sprintf("name is %s and age is %d\n", u.Name, u.Age)
}

type service struct {
	store map[string]*User
}

func main() {
	mux := http.NewServeMux()
	srv := service{make(map[string]*User)}
	mux.HandleFunc("/create", srv.Create)
	mux.HandleFunc("/get", srv.GetAll)

	http.ListenAndServe("localhost:8080", mux)
}

// Create Создание юзера
func (s *service) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		content, err := ioutil.ReadAll((r.Body))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		var u User
		if err := json.Unmarshal(content, &u); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		// splittedContent := strings.Split(string(content), " ")
		s.store[u.Name] = &u

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("User was created " + u.Name))

		return
	}

	// а если не ПОСТ, то верну ошибку
	w.WriteHeader(http.StatusBadRequest)
}

// GetAll Получение всех пользователей
func (s *service) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		response := ""
		for _, user := range s.store {
			response += user.toString()
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))

		return
	}

	w.WriteHeader(http.StatusBadRequest)
}
