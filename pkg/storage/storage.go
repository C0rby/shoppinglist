package storage

import (
	"github.com/c0rby/shoppinglist/pkg/model"
)

func New() Store {
	return Store{
		lists: []model.ShoppingList{
			{ID: "123", Name: "Corby's list"},
			{ID: "456", Name: "Cri-kee's list"},
		},
		entries: map[string][]model.Entry{
			"123": {
				{
					ID:   "789",
					Name: "Tomatos",
				},
				{
					ID:   "101",
					Name: "Pizza",
				},
			},
			"456": {
				{
					ID:   "121",
					Name: "Pho",
				},
				{
					ID:   "141",
					Name: "Pepsi Max",
				},
			},
		},
	}
}

type Store struct {
	lists   []model.ShoppingList
	entries map[string][]model.Entry
}

func (s Store) GetShoppingLists() ([]model.ShoppingList, error) {
	return s.lists, nil
}

func (s Store) GetShoppingList(id string) (model.ShoppingList, error) {
	for _, l := range s.lists {
		if l.ID == id {
			return l, nil
		}
	}
	return model.ShoppingList{}, nil
}

func (s Store) GetShoppingListEntries(id string) ([]model.Entry, error) {
	return s.entries[id], nil
}
