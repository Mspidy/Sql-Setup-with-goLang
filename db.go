package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id   int
	Name string
	Text string
}

func main() {
	db, e := sql.Open("mysql", "root:Prabhat@2022@/student")
	ErrorCheck(e)

	//close database after all work is done
	defer db.Close()
	PingDB(db)
}
func PingDB(db *sql.DB) {
	err := db.Ping()
	ErrorCheck(err)
}

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}
