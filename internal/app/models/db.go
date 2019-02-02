package models

import "database/sql"

type Datastore interface {
	Info() string
}

type DB struct {
	*sql.DB
}

func NewDB(connectString string) (*DB, error) {
	db, err := sql.Open("sqlserver", connectString)
	if err != nil {
		return nil, err
	}
	// if err = db.Ping(); err != nil {
	// 	return nil, err
	// }
	return &DB{db}, nil
}

func (db *DB) Info() string {
	return "MS SQL Server"
}
