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
	StoreShoppingList(list model.ShoppingList) error
	DeleteShoppingList(id string) error
	StoreShoppingListEntry(listID string, entry model.Entry) error
	UpdateShoppingListEntry(entry model.Entry) error
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

func (s InMemStore) StoreShoppingList(list model.ShoppingList) error {
	s.lists = append(s.lists, list)
	return nil
}

func (s InMemStore) DeleteShoppingList(id string) error {
	return nil
}

func (s InMemStore) StoreShoppingListEntry(listID string, entry model.Entry) error {
	s.entries[listID] = append(s.entries[listID], entry)
	return nil
}

func (s InMemStore) UpdateShoppingListEntry(entry model.Entry) error {
	return nil
}
