package admin

import (
	"database/sql"
	"github.com/gorilla/mux"
	"net/http"
)

func DeptSubrouter(r *mux.Router, db *sql.DB) {

	subrouter := r.PathPrefix("/dept").Subrouter()

	subrouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("dept\n"))
	})

}
