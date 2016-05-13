package routes

import (
	"database/sql"
	"github.com/gorilla/mux"
	"net/http"

	"zhgl-goserver/routes/admin"
)

func AdminSubrouter(r *mux.Router, db *sql.DB) {

	subrouter := r.PathPrefix("/admin").Subrouter()

	subrouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("admin\n"))
	})

	// subrouter list
	admin.UserSubrouter(subrouter, db)
	admin.DeptSubrouter(subrouter, db)
	admin.AppSubrouter(subrouter, db)
	admin.MenuSubrouter(subrouter, db)

}
