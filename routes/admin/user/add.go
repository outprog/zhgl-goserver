package user

import (
	"log"
	"net/http"

	"zhgl-goserver/lib/condb"
	"zhgl-goserver/lib/httpjsondone"
	"zhgl-goserver/lib/md5passwd"
)

// 添加用户
func Add(w http.ResponseWriter, r *http.Request) {
	db := condb.Open()
	defer db.Close()

	res := map[string]string{
		"stat": "false",
		"info": "错误的输入格式",
	}
	template := map[string]string{
		"user_id":     "",
		"user_name":   "",
		"user_status": "",
		"user_passwd": "",
		"tellerno":    "",
	}

	body := httpjsondone.GetBody(r)
	if (body["user_id"] == "") || (body["user_name"] == "") || (body["user_passwd"] == "") || (body["user_status"] == "") {
		httpjsondone.SendRes(w, nil, res, template)
		return
	}

	userid := body["user_id"]
	username := body["user_name"]
	userstatus := body["user_status"]
	userpasswd := md5passwd.Get(body["user_passwd"])
	tellerno := body["tellerno"]

	log.Println("add user:", userid)

	stmt, _ := db.Prepare("insert into userlist (user_id, user_name, user_status, user_password, tellerno) values ('" + userid + "', '" + username + "', '" + userstatus + "', '" + userpasswd + "', '" + tellerno + "')")
	_, err := stmt.Exec()
	stmt.Close()

	if err != nil {
		res["stat"] = "false"
		res["info"] = err.Error()
		httpjsondone.SendRes(w, nil, res, template)
		return
	}

	res["stat"] = "true"
	res["info"] = "用户添加成功"
	httpjsondone.SendRes(w, nil, res, template)
}
