package portalAdmin

import (
	"net/http"

	"zhgl-goserver/routes/portal-admin/class"
)

func ClassSubrouter(path string) {

	subrouter := prouter.PathPrefix(path).Subrouter()

	subrouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("class\n"))
	})

	// 新增分类
	subrouter.HandleFunc("/add", class.Add)
	// 删除分类
	subrouter.HandleFunc("/del", class.Del)
	// 修改分类
	subrouter.HandleFunc("/update", class.Update)
	// 查询分类
	subrouter.HandleFunc("/query", class.Query)

}
