package service

import (
	"github.com/c0rby/shoppinglist/pkg/model"
	"github.com/c0rby/shoppinglist/pkg/storage"
)

func New(store storage.Store) Service {
	return Service{store: store}
}

type Service struct {
	store storage.Store
}

func (s Service) GetShoppingLists() ([]model.ShoppingList, error) {
	return s.store.GetShoppingLists()
}

func (s Service) GetShoppingList(id string) (model.ShoppingList, error) {
	return s.store.GetShoppingList(id)
}

func (s Service) GetShoppingListEntries(id string) ([]model.Entry, error) {
	return s.store.GetShoppingListEntries(id)
}
