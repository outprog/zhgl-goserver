package admin

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"zhgl-goserver/lib/condb"
	"zhgl-goserver/lib/httpjsondone"
	"zhgl-goserver/lib/menus"
	"zhgl-goserver/routes/admin/menu"
)

func MenuSubrouter(path string) {

	subrouter := prouter.PathPrefix(path).Subrouter()

	subrouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("app\n"))
	})

	// 获取菜单
	subrouter.HandleFunc("/query/{userid}", func(w http.ResponseWriter, r *http.Request) {
		db := condb.Open()
		defer db.Close()

		userid := mux.Vars(r)["userid"]

		data := menus.GenSysMenu(db, userid)

		log.Println("user:", userid, "get menus")

		httpjsondone.SendRes(w, data, nil, nil)
	})

	// 新增菜单
	subrouter.HandleFunc("/add", menu.Add)
	// 删除菜单
	subrouter.HandleFunc("/del", menu.Del)
	// 更新菜单
	subrouter.HandleFunc("/update", menu.Update)
	// 查询菜单
	subrouter.HandleFunc("/query", menu.Query)
	// 菜单排序
	subrouter.HandleFunc("/sort", menu.Sort)

}
