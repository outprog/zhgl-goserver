package user

import (
	"log"
	"net/http"

	"zhgl-goserver/lib/httpjsondone"
)

// 删除用户权限
func DelRole(w http.ResponseWriter, r *http.Request) {

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

	log.Println("delete user", userid, "role", roleid)

	sql := "delete from mis.rel_user_role " +
		"   where ROLE_ID = '" + roleid + "'" +
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
