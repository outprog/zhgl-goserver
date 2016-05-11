package admin

import (
	"database/sql"
	"github.com/gorilla/mux"
	"net/http"
)

func AppSubrouter(r *mux.Router, db *sql.DB) {

	subrouter := r.PathPrefix("/app").Subrouter()

	subrouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("app\n"))
	})

}
