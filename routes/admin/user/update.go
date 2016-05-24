package user

import (
	"log"
	"net/http"

	"zhgl-goserver/lib/httpjsondone"
)

// 更新用户信息
func Update(w http.ResponseWriter, r *http.Request) {

	res := map[string]string{
		"stat": "false",
		"info": "错误的输入格式",
	}
	template := map[string]string{
		"user_id":     "",
		"user_name":   "",
		"tellerno":    "",
		"user_email":  "",
		"user_mobile": "",
		"user_status": "",
	}

	body := httpjsondone.GetBody(r)
	if body["user_id"] == "" {
		httpjsondone.SendRes(w, nil, res, template)
		return
	}

	userid := body["user_id"]
	username := body["user_name"]
	tellerno := body["tellerno"]
	useremail := body["user_email"]
	usermobile := body["user_mobile"]
	userstatus := body["user_status"]

	log.Println("update user:", userid)

	sql := "update mis.userlist t  " +
		"	set t.USER_NAME = '" + username + "', " +
		"		t.TELLERNO = '" + tellerno + "', " +
		"		t.USER_EMAIL = '" + useremail + "', " +
		"		t.USER_MOBILE = '" + usermobile + "', " +
		"		t.USER_STATUS = '" + userstatus + "' " +
		"	where t.USER_ID = '" + userid + "'"
	stmt, _ := db.Prepare(sql)
	defer stmt.Close()
	upres, _ := stmt.Exec()
	rowCnt, err := upres.RowsAffected()

	if err != nil {
		res["stat"] = "false"
		res["info"] = err.Error()
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
