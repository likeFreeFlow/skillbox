package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
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
	result, err := db.Exec("insert into users (name, age) values ($1, $2);", u.Name, u.Age)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	id, err := result.LastInsertId()
	return int(id), err
}

func (u *User) Delete(db *sql.DB) error {
	_, err := db.Exec("delete from users where id = $1", u.ID)
	return err
}

func (u *User) Get(db *sql.DB) error {
	row := db.QueryRow("select * from users where id=$1", u.ID)
	err := row.Scan(u.ID, u.Name, u.Age)
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
	_ = u.getFriendsID(db)
	_ = newFriend.getFriendsID(db)
	return nil
}

func (u *User) getFriendsID(db *sql.DB) error {
	rows, err := db.Query("select * from friends where user1id=$1 or user2id=$2", u.ID, u.ID)
	if err != nil {
		log.Println(err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		friendID1, friendID2 := 0, 0
		err = rows.Scan(&friendID1, friendID2)
		if err != nil {
			log.Println(err)
			continue
		}
		friendID := 0
		if friendID1 == u.ID {
			friendID = friendID2
		} else {
			friendID = friendID1
		}
		u.Friends = append(u.Friends, friendID)
	}
	return nil
}

func (u *User) GetFriends(db *sql.DB) ([]User, error) {
	rows, err := db.Query("select * from users join friends on friends.user1id=$1 or friends.user2id=$2", u.ID, u.ID)
	if err != nil {
		log.Println(err)
		return []User{}, err
	}
	defer rows.Close()
	friends := []User{}
	for rows.Next() {
		uu := User{}
		err = rows.Scan(&uu.ID, uu.Name, uu.Age)
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
