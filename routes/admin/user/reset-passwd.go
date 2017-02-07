package user

import (
	"log"
	"net/http"

	"zhgl-goserver/lib/condb"
	"zhgl-goserver/lib/httpjsondone"
)

// 重置密码
func ResetPasswd(w http.ResponseWriter, r *http.Request) {
	db := condb.Open()
	defer db.Close()

	res := map[string]string{
		"stat": "false",
		"info": "错误的输入格式",
	}
	template := map[string]string{
		"user_id": "",
	}

	body := httpjsondone.GetBody(r)
	if body["user_id"] == "" {
		httpjsondone.SendRes(w, nil, res, template)
		return
	}

	userid := body["user_id"]
	resetpasswd := "E10ADC3949BA59ABBE56E057F20F883E"

	log.Println("user:", userid, "reset password")

	stmt, _ := db.Prepare("update mis.userlist set user_password = '" + resetpasswd + "' where user_id = '" + userid + "'")
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

	httpjsondone.SendRes(w, nil, res, template)
}
