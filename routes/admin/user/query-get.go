package user

import (
	"log"
	"net/http"

	"github.com/elgs/gosqljson"
	"github.com/gorilla/mux"

	"zhgl-goserver/lib/httpjsondone"
)

// 查询用户 GET 方法
func QueryGet(w http.ResponseWriter, r *http.Request) {

	userid := mux.Vars(r)["userid"]

	log.Println("query get user")

	sql := "SELECT t.user_id, t.user_name, t.user_status, t.tellerno, " +
		"          t.user_email, t.user_mobile, t.user_lastlogin, " +
		"          t2.dept_id, t3.dept_name " +
		"   FROM mis.userlist t LEFT JOIN mis.rel_user_dep t2 " +
		"   ON t.user_id = t2.user_id LEFT JOIN mis.department t3 " +
		"   ON t2.dept_id = t3.dept_id " +
		"    WHERE t.user_id = '" + userid + "' "
	data, _ := gosqljson.QueryDbToMap(db, "upper", sql)

	httpjsondone.SendRes(w, data, nil, nil)
}
