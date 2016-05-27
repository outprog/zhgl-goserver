package admin

import (
	"net/http"

	"zhgl-goserver/routes/admin/role"
)

func RoleSubrouter(path string) {

	subrouter := prouter.PathPrefix(path).Subrouter()

	subrouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("role\n"))
	})

	role.Init(db)
	// 新增权限
	subrouter.HandleFunc("/add", role.Add)
	// 查询权限
	subrouter.HandleFunc("/query", role.Query)
	// 修改权限
	subrouter.HandleFunc("/update", role.Update)
	// 删除权限
	subrouter.HandleFunc("/del", role.Del)

}
