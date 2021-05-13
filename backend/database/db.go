package database

import (
	"database/sql"

	"github.com/LKezHn/React-Go-Project/libs/ErrorManager"
	_ "github.com/go-sql-driver/mysql"
)

func Init() *sql.DB {
	db, err := sql.Open("mysql", "admin:admin@/Trill")
	ErrorManager.ErrorCheck(err)
	return db
}

func CloseConnection(db *sql.DB) {
	db.Close()
}
