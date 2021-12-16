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
