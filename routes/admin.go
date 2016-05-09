package routes

import (
	"github.com/gorilla/mux"
	"net/http"

	"zhgl-goserver/routes/admin"
)

func AdminSubrouter(r *mux.Router) {

	subrouter := r.PathPrefix("/admin").Subrouter()

	subrouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("admin\n"))
	})

	// subrouter list
	admin.UserSubrouter(subrouter)
	admin.DeptSubrouter(subrouter)
	admin.AppSubrouter(subrouter)

}
