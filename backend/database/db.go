package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Init() *sql.DB {
	db, err := sql.Open("mysql", "admin:admin@/Trill")
	if err != nil {
		fmt.Println(err.Error())
	}
	return db
}

func CloseConnection(db *sql.DB) {
	db.Close()
}
