package admin

import (
	"database/sql"
	"github.com/elgs/gosqljson"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"zhgl-goserver/lib/md5passwd"
	"zhgl-goserver/lib/stdres"
)

func UserSubrouter(r *mux.Router, db *sql.DB) {

	subrouter := r.PathPrefix("/user").Subrouter()

	subrouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("user\n"))
	})

	// get user info and confirm passwd
	subrouter.HandleFunc("/confirm-passwd/{userid}/{passwd}", func(w http.ResponseWriter, r *http.Request) {

		userid := mux.Vars(r)["userid"]
		passwd := md5passwd.Get(mux.Vars(r)["passwd"])
		var stat, info string

		log.Println("user:", userid, "confirm passwd")

		data, _ := gosqljson.QueryDbToMap(db, "upper", "SELECT * FROM userlist where user_id=?", userid)

		if len(data) == 1 {
			if passwd == data[0]["USER_PASSWORD"] {
				stat = "true"
				info = ""
				delete(data[0], "USER_PASSWORD")
			} else {
				data = data[:0]
				stat = "false"
				info = "wrong password"
			}
		} else {
			stat = "false"
			info = "no user"
		}

		res := stdres.Get(data, stat, info)
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	})

}
