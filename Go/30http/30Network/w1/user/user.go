package User

import (
	"encoding/json"
	"fmt"
	"log"
	check "main/checks"
	"net/http"
	"strconv"
	"sync"

	"github.com/go-chi/chi/v5"
)

type User struct {
	ID      int
	Name    string `json:"name"`
	Age     int    `json:"age"`
	friends []*User
}

var UserStorage = make(Storage)

func (u User) String() string {
	return fmt.Sprintf("<User> %v. %v (%v)", u.ID, u.Name, u.Age)
}

func (u User) Create(w http.ResponseWriter, r *http.Request) {
	content, ok := check.IsPOSTAndGetContent(&w, r, "Ошбика при создании пользователя")
	if !ok {
		return
	}
	if err := json.Unmarshal(content, &u); err != nil {
		check.ReturnIntServErr(&w, err)
		return
	}

	u.ID = len(UserStorage) + 1
	UserStorage[u.ID] = &u

	log.Printf("Создан пользователь: %v\n", u.String())

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(strconv.Itoa(u.ID)))
}

type newFriends struct {
	Id1 int `json:"source_id"`
	Id2 int `json:"target_id"`
}

func (u User) AddFriend(w http.ResponseWriter, r *http.Request) {
	content, ok := check.IsPOSTAndGetContent(&w, r, "Ошбика при создании пользователя")
	if !ok {
		return
	}
	nFriends := newFriends{}
	if err := json.Unmarshal(content, &nFriends); err != nil {
		check.ReturnIntServErr(&w, err)
		return
	}
	u1, ok := UserStorage.Search(nFriends.Id1)
	if !ok {
		check.ReturnIntServErr(&w, check.Err{MSG: "Не найден пользователь отправляющий запрос"})
		return
	}
	u2, ok := UserStorage.Search(nFriends.Id2)
	if !ok {
		check.ReturnIntServErr(&w, check.Err{MSG: "Не найден пользователь принимающий запрос"})
		return
	}
	u1.friends = append(u1.friends, u2)
	u2.friends = append(u2.friends, u1)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%v и %v теперь друзья", u1.Name, u2.Name)))
	log.Printf("Добавление в друзья: %v <-> %v\n", u1.String(), u2.String())
}

func (u *User) DeleteFriend(uTarget *User, wg *sync.WaitGroup) {
	defer wg.Done()
	i := 0
	for iEl, el := range u.friends {
		if el == uTarget {
			i = iEl
			break
		}
	}
	u.friends = append(u.friends[:i], u.friends[i+1:]...)
}

type userDelete struct {
	Id int `json:"target_id"`
}

func (u User) Delete(w http.ResponseWriter, r *http.Request) {
	content, ok := check.ReadContent(&w, r.Body)
	if !ok {
		check.ReturnIntServErr(&w, check.Err{MSG: "Ошибка при удалении пользователя"})
		log.Println("Ошбика при удалении пользователя")
		return
	}
	uDelete := userDelete{}
	if err := json.Unmarshal(content, &uDelete); err != nil {
		check.ReturnIntServErr(&w, err)
		return
	}
	uu := UserStorage[uDelete.Id]
	wg := sync.WaitGroup{}
	wg.Add(len(uu.friends))
	for _, el := range uu.friends {
		go el.DeleteFriend(uu, &wg)
	}
	wg.Wait()
	delete(UserStorage, uu.ID)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(uu.Name))
	log.Printf("Удален пользователь: %v\n", uu.String())
}

func (u User) GetFriend(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		check.ReturnIntServErr(&w, err)
		return
	}
	u = *UserStorage[id]
	b, err := json.Marshal(u.friends)
	if err != nil {
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
	log.Printf("Запрошен список друзей для %v\n", u)
}

type userAge struct {
	Age int `json:"new_age"`
}

func (u User) SetAge(w http.ResponseWriter, r *http.Request) {
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
	UserStorage[id].Age = uAge.Age
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Возраст пользователя обновлен"))
	log.Printf("Возраст пользователя %v обновлен", UserStorage[id])
}
