package storage

import (
	"github.com/c0rby/shoppinglist/pkg/model"
)

func NewInMemStore() Store {
	return InMemStore{
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

type Store interface {
	GetShoppingLists() ([]model.ShoppingList, error)
	GetShoppingList(id string) (model.ShoppingList, error)
	GetShoppingListEntries(id string) ([]model.Entry, error)
	StoreShoppingList(l model.ShoppingList) error
	DeleteShoppingList(id string) error
	StoreShoppingListEntry(listID string, entry model.Entry) error
	UpdateShoppingListEntry(e model.Entry) error
	GetUsers() ([]model.User, error)
	StoreUser(u model.User) error
	FindUserByName(name string) (model.User, error)
}

type InMemStore struct {
	lists   []model.ShoppingList
	entries map[string][]model.Entry
}

func (s InMemStore) GetShoppingLists() ([]model.ShoppingList, error) {
	return s.lists, nil
}

func (s InMemStore) GetShoppingList(id string) (model.ShoppingList, error) {
	for _, l := range s.lists {
		if l.ID == id {
			return l, nil
		}
	}
	return model.ShoppingList{}, nil
}

func (s InMemStore) GetShoppingListEntries(id string) ([]model.Entry, error) {
	return s.entries[id], nil
}

func (s InMemStore) StoreShoppingList(l model.ShoppingList) error {
	s.lists = append(s.lists, l)
	return nil
}

func (s InMemStore) DeleteShoppingList(id string) error {
	return nil
}

func (s InMemStore) StoreShoppingListEntry(listID string, e model.Entry) error {
	s.entries[listID] = append(s.entries[listID], e)
	return nil
}

func (s InMemStore) UpdateShoppingListEntry(e model.Entry) error {
	return nil
}

func (s InMemStore) GetUsers() ([]model.User, error) {
	return nil, nil
}

func (s InMemStore) StoreUser(u model.User) error {
	return nil
}

func (s InMemStore) FindUserByName(n string) (model.User, error) {
	return model.User{}, nil
}
