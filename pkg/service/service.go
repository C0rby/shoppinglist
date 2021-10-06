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

func (s Service) GetLists() ([]model.List, error) {
	return s.store.GetLists()
}
