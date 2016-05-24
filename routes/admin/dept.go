package admin

import (
	"net/http"

	"zhgl-goserver/routes/admin/dept"
)

func DeptSubrouter(path string) {

	subrouter := prouter.PathPrefix(path).Subrouter()

	subrouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("dept\n"))
	})

	dept.Init(db)
	// 新增部门
	subrouter.HandleFunc("/add", dept.Add)
	// 删除部门
	subrouter.HandleFunc("/del", dept.Del)
	// 更新部门
	subrouter.HandleFunc("/update", dept.Update)
}
