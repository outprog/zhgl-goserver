package portalAdmin

import (
	"net/http"

	"zhgl-goserver/routes/portal-admin/visitor"
)

func VisitorSubrouter(path string) {

	subrouter := prouter.PathPrefix(path).Subrouter()

	subrouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("post\n"))
	})

	// 新增访客
	subrouter.HandleFunc("/add", visitor.Add)
	// 查询访客
	subrouter.HandleFunc("/query", visitor.Query)

}
