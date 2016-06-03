package user

import (
	"log"
	"net/http"

	"zhgl-goserver/lib/httpjsondone"
)

// 删除用户应用系统
func DelApp(w http.ResponseWriter, r *http.Request) {

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

	log.Println("delete user", userid, "sys", sysid)

	sql := "delete from mis.sys_manager " +
		"   where SYS_ID = '" + sysid + "'" +
		"   and USER_ID = '" + userid + "'"

	upres, err := db.Exec(sql)
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
		res["info"] = "删除成功"
	} else {
		res["stat"] = "false"
		res["info"] = "删除失败或没有变动"
	}

	httpjsondone.SendRes(w, nil, res, template)
}
