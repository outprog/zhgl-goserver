package admin

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"zhgl-goserver/lib/menus"
	"zhgl-goserver/lib/httpjsondone"
)

func MenuSubrouter(r *mux.Router, db *sql.DB) {

	subrouter := r.PathPrefix("/menu").Subrouter()

	subrouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("app\n"))
	})

	// 获取菜单
	subrouter.HandleFunc("/query/{userid}", func(w http.ResponseWriter, r *http.Request) {

		userid := mux.Vars(r)["userid"]

    data := menus.GenSysMenu(db, userid)

		log.Println("user:", userid, "get menus")

		genres := httpjsondone.GenRes(data, nil, nil)
		w.Header().Set("Content-Type", "application/json")
		w.Write(genres)
	})

}