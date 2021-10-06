package storage

import (
	"github.com/c0rby/shoppinglist/pkg/model"
)

func New() Store {
	return Store{lists: []model.List{
		{ID: "123", Name: "Corby's list"},
		{ID: "456", Name: "Cri-kee's list"},
	}}
}

type Store struct {
	lists []model.List
}

func (s Store) GetLists() ([]model.List, error) {
	return s.lists, nil
}
