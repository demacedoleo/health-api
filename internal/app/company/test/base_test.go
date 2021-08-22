package test

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"log"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Println("an error '%s' was not expected when opening a stub database connection: " + err.Error())
	}

	return db, mock
}
