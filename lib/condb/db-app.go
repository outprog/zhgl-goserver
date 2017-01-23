package condb

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func OpenApp() *sql.DB {
	// connect db
	mydb, err := sql.Open("mysql", "root@/app")
	if err != nil {
		log.Println(err.Error())
	}

	return mydb
}
