package model

type ShoppingList struct {
	ID   string
	Name string
}

type Entry struct {
	ID     string
	Name   string
	Amount string
	Buy    bool
}

type User struct {
	ID       string
	Name     string
	Password string
}

type Session struct {
	ID   string
	User User
}
