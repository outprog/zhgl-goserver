package admin

import (
	"github.com/gorilla/mux"
	"net/http"
)

func UserSubrouter(r *mux.Router) {

	subrouter := r.PathPrefix("/user").Subrouter()

	subrouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("user\n"))
	})

}
