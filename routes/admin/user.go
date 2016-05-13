package admin

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/elgs/gosqljson"
	"github.com/gorilla/mux"

	"zhgl-goserver/lib/httpjsondone"
	"zhgl-goserver/lib/md5passwd"
)

func UserSubrouter(r *mux.Router, db *sql.DB) {

	subrouter := r.PathPrefix("/user").Subrouter()

	subrouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("user\n"))
	})

	// 验证密码并获取用户信息
	subrouter.HandleFunc("/confirm-passwd", func(w http.ResponseWriter, r *http.Request) {

		stat := "false"
		info := "error input"
		template := []map[string]string{}
		template = append(template, map[string]string{
			"user_id":     "",
			"user_passwd": "",
		})

		body := httpjsondone.GetBody(r)
		if (body["user_id"] == "") || (body["user_passwd"] == "") {
			res := httpjsondone.GenRes(nil, stat, info, template)
			w.Header().Set("Content-Type", "application/json")
			w.Write(res)
			return
		}

		userid := body["user_id"]
		passwd := md5passwd.Get(body["user_passwd"])

		log.Println("user:", userid, "confirm password")

		data, _ := gosqljson.QueryDbToMap(db, "upper",
			"SELECT t.user_id, "+
				"   t.user_name, "+
				"   t.user_password, "+
				"   t.user_status, "+
				"   (SELECT dept_id FROM mis.rel_user_dep where user_id = t.user_id) as dept_id "+
				" FROM userlist t where t.user_id=?",
			userid)

		if len(data) == 1 {
			if passwd == data[0]["USER_PASSWORD"] {
				stat = "true"
				info = ""
				delete(data[0], "USER_PASSWORD")
			} else {
				stat = "false"
				info = "wrong password"
				data = data[:0]
			}
		} else {
			stat = "false"
			info = "no user"
		}

		res := httpjsondone.GenRes(data, stat, info, template)
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	})

}
