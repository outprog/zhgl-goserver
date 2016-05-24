package dept

import (
	"database/sql"
)

var db *sql.DB

func Init(mydb *sql.DB) {
	db = mydb
	return
}
