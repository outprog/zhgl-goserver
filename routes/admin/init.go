package admin

import (
	"database/sql"

	"github.com/gorilla/mux"
)

var db *sql.DB
var prouter *mux.Router

func Init(mydb *sql.DB, r *mux.Router) {
	db = mydb
	prouter = r

	return
}
