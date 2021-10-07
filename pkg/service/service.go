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

func (s Service) DeleteShoppingList(id string) error {
	return s.store.DeleteShoppingList(id)
}

func (s Service) CreateShoppingListEntry(listID string, entry model.Entry) (model.Entry, error) {
	entry.ID = uuid.New().String()
	entry.Buy = true
	if err := s.store.StoreShoppingListEntry(listID, entry); err != nil {
		return model.Entry{}, err
	}
	return entry, nil
}

func (s Service) UpdateShoppingListEntry(entry model.Entry) (model.Entry, error) {
	if err := s.store.UpdateShoppingListEntry(entry); err != nil {
		return model.Entry{}, err
	}
	return entry, nil
}
