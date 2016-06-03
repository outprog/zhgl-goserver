package user

import (
	"log"
	"net/http"

	"zhgl-goserver/lib/httpjsondone"
)

// 添加用户应用系统
func AddApp(w http.ResponseWriter, r *http.Request) {

	res := map[string]string{
		"stat": "false",
		"info": "错误的输入格式",
	}
	template := map[string]string{
		"user_id": "",
		"sys_id":  "",
	}

	body := httpjsondone.GetBody(r)
	if (body["user_id"] == "") || (body["sys_id"] == "") {
		httpjsondone.SendRes(w, nil, res, template)
		return
	}

	userid := body["user_id"]
	sysid := body["sys_id"]

	log.Println("user", userid, "add sys", sysid)

	sql := "insert into mis.sys_manager (sys_id, user_id) values ('" +
		sysid + "', '" +
		userid + "')"
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
