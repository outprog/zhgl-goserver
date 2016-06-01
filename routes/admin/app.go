package admin

import (
	"net/http"

	"zhgl-goserver/routes/admin/app"
)

func AppSubrouter(path string) {

	subrouter := prouter.PathPrefix(path).Subrouter()

	subrouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("app\n"))
	})

	app.Init(db)
	// 新增系统
	subrouter.HandleFunc("/add", app.Add)
	// 删除系统
	subrouter.HandleFunc("/del", app.Del)
	// 更新系统
	subrouter.HandleFunc("/update", app.Update)
	// 查询系统
	subrouter.HandleFunc("/query", app.Query)

}
