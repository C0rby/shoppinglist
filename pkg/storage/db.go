package storage

import (
	"database/sql"

	"github.com/c0rby/shoppinglist/pkg/model"
	_ "github.com/mattn/go-sqlite3"
)

func NewSqlite3DB(file string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewSqlStore(db *sql.DB) (Store, error) {
	return SqlStore{db: db}, nil
}

func CreateTables(db *sql.DB) error {
	shoppinglistsTable := `
	CREATE TABLE IF NOT EXISTS shoppinglists(
		Id TEXT NOT NULL PRIMARY KEY,
		Name TEXT
	);`

	listEntriesTable := `
	CREATE TABLE IF NOT EXISTS entries(
		Id TEXT NOT NULL PRIMARY KEY,
		Name TEXT,
		List_Id TEXT,
		FOREIGN KEY(List_Id) REFERENCES shoppinglists(Id)
	);`
	if _, err := db.Exec(shoppinglistsTable); err != nil {
		return err
	}
	_, err := db.Exec(listEntriesTable)
	return err
}

type SqlStore struct {
	db *sql.DB
}

func (s SqlStore) GetShoppingLists() ([]model.ShoppingList, error) {
	rows, err := s.db.Query("SELECT id, name FROM shoppinglists")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		id    string
		name  string
		lists []model.ShoppingList
	)
	for rows.Next() {
		err = rows.Scan(&id, &name)
		if err != nil {
			return nil, err
		}
		lists = append(lists, model.ShoppingList{ID: id, Name: name})
	}
	return lists, nil
}

func (s SqlStore) GetShoppingList(id string) (model.ShoppingList, error) {
	stmt, err := s.db.Prepare("SELECT id, name FROM shoppinglists where id=?")
	if err != nil {
		return model.ShoppingList{}, err
	}
	defer stmt.Close()

	var (
		listId string
		name   string
	)
	err = stmt.QueryRow(id).Scan(&listId, &name)
	if err != nil {
		return model.ShoppingList{}, err
	}

	return model.ShoppingList{ID: listId, Name: name}, nil
}

func (s SqlStore) GetShoppingListEntries(id string) ([]model.Entry, error) {
	stmt, err := s.db.Prepare("SELECT id, name FROM entries where list_id=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var (
		entryId string
		name    string
		entries []model.Entry
	)
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&entryId, &name)
		if err != nil {
			return nil, err
		}
		entries = append(entries, model.Entry{ID: entryId, Name: name})
	}
	return entries, nil
}

func (s SqlStore) StoreShoppingList(list model.ShoppingList) error {
	sqlAddShoppingList := `
	INSERT INTO shoppinglists(
		Id,
		Name
	) VALUES(?, ?)`

	stmt, err := s.db.Prepare(sqlAddShoppingList)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(list.ID, list.Name)
	return err
}
