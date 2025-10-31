package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB(filepath string)  {
 DB, err := sql.Open("sqlite3", "api.db")
 if err != nil {
	 panic("could not open the database")
 }
 DB.SetMaxOpenConns(10)
}