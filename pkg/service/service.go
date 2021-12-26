package service

import (
	"github.com/c0rby/shoppinglist/pkg/model"
	"github.com/c0rby/shoppinglist/pkg/storage"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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
	list.ID = uuid.NewString()
	if err := s.store.StoreShoppingList(list); err != nil {
		return model.ShoppingList{}, err
	}
	return list, nil
}

func (s Service) DeleteShoppingList(id string) error {
	return s.store.DeleteShoppingList(id)
}

func (s Service) CreateShoppingListEntry(listID string, entry model.Entry) (model.Entry, error) {
	entry.ID = uuid.NewString()
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

func (s Service) CreateUser(u model.User) (model.User, error) {
	u.ID = uuid.NewString()

	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return u, err
	}
	u.Password = string(hashed)
	return u, s.store.StoreUser(u)
}

func (s Service) GetUsers() ([]model.User, error) {
	return s.store.GetUsers()
}

func (s Service) AuthenticateUser(name, password string) (model.Session, error) {
	user, err := s.store.FindUserByName(name)
	if err != nil {
		return model.Session{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return model.Session{}, err
	}

	return model.Session{
		ID:   uuid.NewString(),
		User: user,
	}, nil
}

func (s Service) AuthenticateSession(sid string) (model.Session, error) {
	return model.Session{}, nil
}
