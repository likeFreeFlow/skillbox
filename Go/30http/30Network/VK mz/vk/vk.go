package vk

type User struct {
	ID      int
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Friends []*User
}
