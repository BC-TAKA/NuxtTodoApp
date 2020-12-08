package common

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DbConn() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/todo")
	if err != nil {
		panic(err.Error())
	}
	return db
}
