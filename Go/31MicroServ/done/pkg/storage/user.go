package storage

import (
	"database/sql"
	"fmt"

	"log"
)

type User struct {
	ID      int
	Name    string
	Age     int
	Friends []int
}

func (u *User) String() string {
	return fmt.Sprintf("<User> %v. %v (%v)", u.ID, u.Name, u.Age)
}

func (u *User) Create(db *sql.DB) (int, error) {
	result, err := db.Exec("INSERT INTO users (name, age) VALUES (?, ?);", u.Name, u.Age)

	if err != nil {
		log.Println("Storage Users Create Exec", err)
		return 0, err
	}
	id, err := result.LastInsertId()
	return int(id), err
}

func (u *User) Delete(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM users WHERE id = ?", u.ID)
	return err
}

func (u *User) Get(db *sql.DB) error {
	row := db.QueryRow("SELECT * FROM users WHERE id=?", u.ID)
	err := row.Scan(&u.ID, &u.Name, &u.Age)
	if err != nil {
		log.Println(err)
		return err
	}
	err = u.getFriendsID(db)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *User) AddFriend(newFriend *User, db *sql.DB) error {
	_, err := db.Exec("insert into friends (user1id, user2id) values ($1, $2);", u.ID, newFriend.ID)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = db.Exec("insert into friends (user1id, user2id) values ($1, $2);", newFriend.ID, u.ID)
	if err != nil {
		log.Println(err)
		return err
	}
	_ = u.getFriendsID(db)
	_ = newFriend.getFriendsID(db)
	return nil
}

func (u *User) getFriendsID(db *sql.DB) error {
	rows, err := db.Query("select user2id from friends where user1id=$1", u.ID)
	if err != nil {
		log.Println(err)
		return err
	}
	defer rows.Close()
	counter := 0
	for rows.Next() {
		counter++
		log.Println(counter)
		friendID := 0
		err = rows.Scan(&friendID)
		if err != nil {
			log.Println(err)
			continue
		}
		u.Friends = append(u.Friends, friendID)
	}
	return nil
}

func (u *User) GetFriends(db *sql.DB) ([]User, error) {

	rows, err := db.Query("select users.id, users.name, users.age from users join friends on friends.user2id=? where friends.user1id=users.id  ", u.ID)

	if err != nil {
		log.Println(err)
		return []User{}, err
	}
	defer rows.Close()

	friends := []User{}
	for rows.Next() {
		uu := User{}
		err = rows.Scan(&uu.ID, &uu.Name, &uu.Age)
		if err != nil {
			log.Println(err)
			continue
		}
		friends = append(friends, uu)
	}
	return friends, nil
}

func (u *User) SetAge(age int, db *sql.DB) error {
	_, err := db.Exec("update users set age=$1 where id=$2", age, u.ID)
	return err
}
