package user

import (
	"log"
	"net/http"

	"github.com/elgs/gosqljson"

	"zhgl-goserver/lib/httpjsondone"
	"zhgl-goserver/lib/md5passwd"
)

// 更改密码
func UpdatePasswd(w http.ResponseWriter, r *http.Request) {

	res := map[string]string{
		"stat": "false",
		"info": "错误的输入格式",
	}
	template := map[string]string{
		"user_id":     "",
		"user_passwd": "",
		"new_passwd":  "",
	}

	body := httpjsondone.GetBody(r)
	if (body["user_id"] == "") || (body["user_passwd"] == "") || (body["new_passwd"] == "") {
		httpjsondone.SendRes(w, nil, res, template)
		return
	}
	if len(body["new_passwd"]) < 8 {
		res["info"] = "密码不得少于8位"
		httpjsondone.SendRes(w, nil, res, template)
		return
	}

	userid := body["user_id"]
	userpasswd := md5passwd.Get(body["user_passwd"])
	newpasswd := md5passwd.Get(body["new_passwd"])

	log.Println("user:", userid, "change password")

	sql := "SELECT t.user_password " +
		" FROM mis.userlist t where t.user_id='" + userid + "'"
	data, _ := gosqljson.QueryDbToMap(db, "upper", sql)

	if len(data) == 1 {
		if userpasswd == data[0]["USER_PASSWORD"] {

			stmt, _ := db.Prepare("update mis.userlist set user_password = '" + newpasswd + "' where user_id = '" + userid + "'")
			defer stmt.Close()

			upres, err := stmt.Exec()
			if err != nil {
				res["stat"] = "false"
				res["info"] = err.Error()
				httpjsondone.SendRes(w, nil, res, template)
				return
			}

			rowCnt, err := upres.RowsAffected()
			if err != nil {
				res["stat"] = "false"
				res["info"] = err.Error()
				httpjsondone.SendRes(w, nil, res, template)
				return
			}
			if rowCnt == 1 {
				res["stat"] = "true"
				res["info"] = "更新成功"
			} else {
				res["stat"] = "false"
				res["info"] = "更新失败或没有改动"
			}

		} else {
			res["stat"] = "false"
			res["info"] = "密码错误"
		}
	} else {
		res["stat"] = "false"
		res["info"] = "没有该用户"
	}

	httpjsondone.SendRes(w, nil, res, template)
}
