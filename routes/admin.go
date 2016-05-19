package routes

import (
	"net/http"

	"zhgl-goserver/routes/admin"
)

func AdminSubrouter(path string) {

	subrouter := prouter.PathPrefix(path).Subrouter()

	subrouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("admin\n"))
	})

	admin.Init(db, subrouter)
	// subrouter list
	admin.UserSubrouter("/user")
	admin.DeptSubrouter("/dept")
	admin.AppSubrouter("/app")
	admin.MenuSubrouter("/menu")

}
