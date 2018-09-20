package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// MySQLDB Conn
var MySQLDB *sql.DB

func init() {
	db, err := sql.Open("mysql", "root:calvin.1.@tcp(39.108.104.253:3306)/5iweb?charset=utf8")
	MySQLDB = db
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
