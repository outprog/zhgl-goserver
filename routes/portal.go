package routes

import (
	"net/http"
)

func PortalSubrouter(path string) {

	subrouter := prouter.PathPrefix(path).Subrouter()

	subrouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("portal\n"))
	})

}
