package user

import (
	"log"
	"net/http"

	"zhgl-goserver/lib/condb"
	"zhgl-goserver/lib/httpjsondone"
)

// 设置用户部门
func SettingDept(w http.ResponseWriter, r *http.Request) {
	db := condb.Open()
	defer db.Close()

	res := map[string]string{
		"stat": "false",
		"info": "错误的输入格式",
	}
	template := map[string]string{
		"user_id":   "",
		"user_dept": "",
	}

	body := httpjsondone.GetBody(r)
	if (body["user_id"] == "") || (body["user_dept"] == "") {
		httpjsondone.SendRes(w, nil, res, template)
		return
	}

	userid := body["user_id"]
	userdept := body["user_dept"]

	log.Println("user", userid, "set dept", userdept)

	tx, _ := db.Begin()
	_, errDelRel := tx.Exec("delete from rel_user_dep where user_id = '" + userid + "'")
	_, errInsertRel := tx.Exec("insert into rel_user_dep values('" + userid + "', '" + userdept + "')")

	if (errDelRel != nil) ||
		(errInsertRel != nil) {

		tx.Rollback()
		res["info"] = "执行失败"
		httpjsondone.SendRes(w, nil, res, template)
		return
	}
	tx.Commit()

	res["stat"] = "true"
	res["info"] = "设置成功"
	httpjsondone.SendRes(w, nil, res, template)
}
