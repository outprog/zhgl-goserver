package admin

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"zhgl-goserver/lib/httpjsondone"
	"zhgl-goserver/lib/menus"
)

func MenuSubrouter(path string) {

	subrouter := prouter.PathPrefix(path).Subrouter()

	subrouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("app\n"))
	})

	// 获取菜单
	subrouter.HandleFunc("/query/{userid}", func(w http.ResponseWriter, r *http.Request) {

		userid := mux.Vars(r)["userid"]

		data := menus.GenSysMenu(db, userid)

		log.Println("user:", userid, "get menus")

		httpjsondone.SendRes(w, data, nil, nil)
	})

}
