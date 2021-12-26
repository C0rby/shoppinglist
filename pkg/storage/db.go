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
		Name TEXT NOT NULL
	);`

	listEntriesTable := `
	CREATE TABLE IF NOT EXISTS entries(
		Id TEXT NOT NULL PRIMARY KEY,
		Name TEXT NOT NULL,
		Amount TEXT,
		Buy BOOLEAN,
		List_Id TEXT,
		CONSTRAINT fk_shoppinglists
			FOREIGN KEY(List_Id)
			REFERENCES shoppinglists(Id)
			ON DELETE CASCADE
	);`

	usersTable := `
	CREATE TABLE IF NOT EXISTS users(
		Id TEXT NOT NULL PRIMARY KEY,
		Name TEXT NOT NULL,
		Password TEXT NOT NULL
	);`
	if _, err := db.Exec(shoppinglistsTable); err != nil {
		return err
	}
	if _, err := db.Exec(listEntriesTable); err != nil {
		return err
	}
	_, err := db.Exec(usersTable)
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
	stmt, err := s.db.Prepare("SELECT Id, Name FROM shoppinglists where Id=?")
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
	stmt, err := s.db.Prepare("SELECT Id, Name, Amount, Buy FROM entries where List_Id=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var (
		entryId string
		name    string
		amount  string
		buy     bool
		entries []model.Entry
	)
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&entryId, &name, &amount, &buy)
		if err != nil {
			return nil, err
		}
		entries = append(entries, model.Entry{ID: entryId, Name: name, Amount: amount, Buy: buy})
	}
	return entries, nil
}

func (s SqlStore) StoreShoppingList(l model.ShoppingList) error {
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

	_, err = stmt.Exec(l.ID, l.Name)
	return err
}

func (s SqlStore) DeleteShoppingList(id string) error {
	sqlDeleteShoppingList := "DELETE FROM shoppinglists WHERE Id = ?"

	_, err := s.db.Exec(sqlDeleteShoppingList, id)
	return err
}

func (s SqlStore) StoreShoppingListEntry(listID string, e model.Entry) error {
	sqlAddShoppingListEntry := `
	INSERT INTO entries(
		Id,
		Name,
		Amount,
		Buy,
		List_Id
	) VALUES(?, ?, ?, ?, ?)`

	stmt, err := s.db.Prepare(sqlAddShoppingListEntry)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID, e.Name, e.Amount, e.Buy, listID)
	return err
}

func (s SqlStore) UpdateShoppingListEntry(e model.Entry) error {
	sqlUpdateShoppingListEntry := `
	UPDATE entries 
	SET Name = ?, Amount = ?, Buy = ?
	WHERE Id = ?;`

	stmt, err := s.db.Prepare(sqlUpdateShoppingListEntry)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Amount, e.Buy, e.ID)
	return err
}

func (s SqlStore) GetUsers() ([]model.User, error) {
	rows, err := s.db.Query("SELECT Id, Name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		id    string
		name  string
		users []model.User
	)
	for rows.Next() {
		err = rows.Scan(&id, &name)
		if err != nil {
			return nil, err
		}
		users = append(users, model.User{ID: id, Name: name})
	}
	return users, nil
}

func (s SqlStore) StoreUser(u model.User) error {
	sqlAddUser := `
	INSERT INTO users(
		Id,
		Name,
		Password
	) VALUES(?, ?, ?)`

	stmt, err := s.db.Prepare(sqlAddUser)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.ID, u.Name, u.Password)
	return err
}

func (s SqlStore) FindUserByName(n string) (model.User, error) {
	stmt, err := s.db.Prepare("SELECT Id, Name, Password FROM users where Name=?")
	if err != nil {
		return model.User{}, nil
	}
	defer stmt.Close()
	var (
		id       string
		name     string
		password string
	)
	row := stmt.QueryRow(n)
	switch err := row.Scan(&id, &name, &password); err {
	case sql.ErrNoRows:
		return model.User{}, err
	case nil:
		return model.User{ID: id, Name: name, Password: password}, nil
	default:
		return model.User{}, err
	}
}
