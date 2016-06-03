package user

import (
	"log"
	"net/http"

	"zhgl-goserver/lib/httpjsondone"
)

// 添加用户权限
func AddRole(w http.ResponseWriter, r *http.Request) {

	res := map[string]string{
		"stat": "false",
		"info": "错误的输入格式",
	}
	template := map[string]string{
		"user_id": "",
		"role_id": "",
	}

	body := httpjsondone.GetBody(r)
	if (body["user_id"] == "") || (body["role_id"] == "") {
		httpjsondone.SendRes(w, nil, res, template)
		return
	}

	userid := body["user_id"]
	roleid := body["role_id"]

	log.Println("user", userid, "add role", roleid)

	sql := "insert into rel_user_role (user_id, role_id) values ('" +
		userid + "', '" +
		roleid + "')"
	stmt, _ := db.Prepare(sql)
	_, err := stmt.Exec()
	defer stmt.Close()
	if err != nil {
		res["stat"] = "false"
		res["info"] = err.Error()
	} else {
		res["stat"] = "true"
		res["info"] = "添加成功"
	}

	httpjsondone.SendRes(w, nil, res, template)
}
