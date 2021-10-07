package service

import (
	"github.com/c0rby/shoppinglist/pkg/model"
	"github.com/c0rby/shoppinglist/pkg/storage"
	"github.com/google/uuid"
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

func (s Service) CreateShoppingList(list model.ShoppingList) (model.ShoppingList, error) {
	list.ID = uuid.New().String()
	if err := s.store.StoreShoppingList(list); err != nil {
		return model.ShoppingList{}, err
	}
	return list, nil
}
