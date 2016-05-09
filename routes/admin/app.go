package admin

import (
	"github.com/gorilla/mux"
	"net/http"
)

func AppSubrouter(r *mux.Router) {

	subrouter := r.PathPrefix("/app").Subrouter()

	subrouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("app\n"))
	})

}
