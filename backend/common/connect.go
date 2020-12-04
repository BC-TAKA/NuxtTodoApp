package common

import "database/sql"

func DbConn() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/todo")
	if err != nil {
		panic(err.Error())
	}
	return db
}
