package model

import (
	"database/sql"
	"log"
	"time"
)

// UserStore implements all user model access functions
type TestTableStore struct {
	db *sql.DB
}

func initUserStore(db *sql.DB) *TestTableStore {
	return &TestTableStore{db}
}

// User represents a user in the system
type TestTable struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	RandomString string `json:"random_string"`
	DateCreate   int    `json:"date_create"`
}

// Create inserts a new user with the given name / password
func (store *TestTableStore) Create(name, randomString string) *TestTable {
	date := int64(time.Now().UnixNano() / 1000)

	result, err := store.db.Exec("INSERT OR REPLACE INTO test_table (id, name, random_string, date_create) VALUES (null, :name, :random_string, :date_create)",
		sql.Named("name", name),
		sql.Named("random_string", randomString),
		sql.Named("date_create", date),
	)

	if err != nil {
		log.Printf("error creating user: %v", err)
		return nil
	}

	id, err := result.LastInsertId()
	return &TestTable{ID: int(id), Name: name, RandomString: randomString, DateCreate: int(date)}
}

// GetByName returns a user with the given name
func (store *TestTableStore) GetByName(name string) *TestTable {
	var user = TestTable{}
	row := store.db.QueryRow("select id,name,random_string,date_create from test_table where name = :name", sql.Named("name", name))
	switch err := row.Scan(&user.ID, &user.Name, &user.RandomString, &user.DateCreate); err {
	case nil:
		return &user
	case sql.ErrNoRows:
		return nil
	default:
		log.Printf("error running db query: %v", err)
		return nil
	}
}

// GetByID returns a user with the given ID
func (store *TestTableStore) GetByID(id int) *TestTable {
	var user = TestTable{}
	row := store.db.QueryRow("select id,name,random_string,date_create from test_table where id = :id", sql.Named("id", id))
	switch err := row.Scan(&user.ID, &user.Name, &user.RandomString, &user.DateCreate); err {
	case nil:
		return &user
	case sql.ErrNoRows:
		return nil
	default:
		log.Printf("error running db query: %v", err)
		return nil
	}
}

// Check verifies the strings match
func (testTable *TestTable) CheckStrings(randomString string) bool {
	return testTable.RandomString == randomString
}
