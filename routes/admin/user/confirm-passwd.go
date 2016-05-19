package user

import (
	"log"
	"net/http"

	"github.com/elgs/gosqljson"

	"zhgl-goserver/lib/httpjsondone"
	"zhgl-goserver/lib/md5passwd"
)

// 验证密码
func ConfirmPasswd(w http.ResponseWriter, r *http.Request) {

	res := map[string]string{
		"stat": "false",
		"info": "错误的输入格式",
	}
	template := map[string]string{
		"user_id":     "",
		"user_passwd": "",
	}

	body := httpjsondone.GetBody(r)
	if (body["user_id"] == "") || (body["user_passwd"] == "") {
		httpjsondone.SendRes(w, nil, res, template)
		return
	}

	userid := body["user_id"]
	passwd := md5passwd.Get(body["user_passwd"])

	log.Println("user:", userid, "confirm password")

	sql := "SELECT t.user_id, " +
		"   t.user_name, " +
		"   t.user_password, " +
		"   t.user_status, " +
		"   (SELECT dept_id FROM mis.rel_user_dep where user_id = t.user_id) as dept_id " +
		" FROM userlist t where t.user_id='" + userid + "'"
	data, _ := gosqljson.QueryDbToMap(db, "upper", sql)

	if len(data) == 1 {
		if passwd == data[0]["USER_PASSWORD"] {
			res["stat"] = "true"
			res["info"] = "密码正确,返回信息"
			delete(data[0], "USER_PASSWORD")
		} else {
			res["stat"] = "false"
			res["info"] = "密码错误"
			data = data[:0]
		}
	} else {
		res["stat"] = "false"
		res["info"] = "没有该用户"
	}

	httpjsondone.SendRes(w, data, res, template)
}
