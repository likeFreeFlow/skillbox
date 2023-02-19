package User

import (
	"encoding/json"
	"fmt"
	"log"
	check "main/pkg/checks"
	"net/http"
	"strconv"

	"main/pkg/storage"

	"github.com/go-chi/chi/v5"
)

type Useru struct {
	ID      int
	Name    string `json:"name"`
	Age     int    `json:"age"`
	friends []int
}

var db = storage.ConnectDB()

func (u Useru) String() string {
	return fmt.Sprintf("<User> %v. %v (%v)", u.ID, u.Name, u.Age)
}

func (u Useru) Create(w http.ResponseWriter, r *http.Request) {
	content, ok := check.IsPOSTAndGetContent(&w, r, "Ошибка при создании пользователя")
	if !ok {
		return
	}
	if err := json.Unmarshal(content, &u); err != nil {
		check.ReturnIntServErr(&w, err)
		return
	}

	uNew := storage.User{
		Name: u.Name,
		Age:  u.Age,
	}
	id, err := uNew.Create(db)
	if err != nil {
		check.ReturnIntServErr(&w, err)
		return
	}
	u.ID = id

	log.Printf("Создан пользователь: %v\n", u.String())
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(strconv.Itoa(u.ID)))
}

type newFriends struct {
	Id1 int `json:"source_id"`
	Id2 int `json:"target_id"`
}

func (u Useru) AddFriend(w http.ResponseWriter, r *http.Request) {
	content, ok := check.IsPOSTAndGetContent(&w, r, "Ошбика при создании пользователя")
	if !ok {
		return
	}
	nFriends := newFriends{}
	if err := json.Unmarshal(content, &nFriends); err != nil {
		check.ReturnIntServErr(&w, err)
		return
	}

	u1 := storage.User{ID: nFriends.Id1}
	if err := u1.Get(db); err != nil {
		check.ReturnIntServErr(&w, err)
		return
	}
	u2 := storage.User{ID: nFriends.Id2}
	if err := u2.Get(db); err != nil {
		check.ReturnIntServErr(&w, err)
		return
	}
	if err := u1.AddFriend(&u2, db); err != nil {
		check.ReturnIntServErr(&w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%v и %v теперь друзья", u1.Name, u2.Name)))
	log.Printf("Добавление в друзья: %v <-> %v\n", u1.String(), u2.String())
}

type userDelete struct {
	Id int `json:"target_id"`
}

func (u Useru) Delete(w http.ResponseWriter, r *http.Request) {
	content, ok := check.ReadContent(&w, r.Body)
	if !ok {
		check.ReturnIntServErr(&w, check.Err{MSG: "Ошибка при удалении пользователя"})
		return
	}
	var uDelete = userDelete{}
	if err := json.Unmarshal(content, &uDelete); err != nil {
		check.ReturnIntServErr(&w, err)
		return
	}

	u1 := storage.User{ID: uDelete.Id}
	if err := u1.Get(db); err != nil {
		check.ReturnIntServErr(&w, err)
		return
	}
	
	if err := u1.Delete(db); err != nil {
		check.ReturnIntServErr(&w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(u1.Name))
	log.Printf("Удален пользователь: %v\n", u1.String())
}

func (u Useru) GetFriend(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		check.ReturnIntServErr(&w, err)
		return
	}
	uu := storage.User{ID: id}
	friends, err := uu.GetFriends(db)
	if err != nil {
		check.ReturnIntServErr(&w, err)
		return
	}
	b, err := json.Marshal(friends)
	if err != nil {
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
	log.Println("Запрошен список друзей для пользователя с id =", id)
}

type userAge struct {
	Age int `json:"new_age"`
}

func (u Useru) SetAge(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		check.ReturnIntServErr(&w, err)
		return
	}
	content, ok := check.ReadContent(&w, r.Body)
	if !ok {
		check.ReturnIntServErr(&w, check.Err{MSG: "Ошибка при обновлении возраста пользователя"})
		log.Println("Ошбика при обновлении возраста пользователя")
		return
	}
	uAge := userAge{}
	if err = json.Unmarshal(content, &uAge); err != nil {
		check.ReturnIntServErr(&w, err)
		return
	}
	uu := storage.User{ID: id}
	if err = uu.SetAge(uAge.Age, db); err != nil {
		check.ReturnIntServErr(&w, err)
		return
	}
	uu.Get(db)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprint("Возраст пользователя обновлен")))
	log.Printf("Возраст пользователя %v обновлен", uu)
}
