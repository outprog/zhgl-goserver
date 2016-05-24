package user

import (
	"log"
	"net/http"

	"zhgl-goserver/lib/httpjsondone"
)

// 删除用户
func Del(w http.ResponseWriter, r *http.Request) {

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
	log.Println("delete user:", userid)

	tx, _ := db.Begin()
	_, errUserList := tx.Exec("delete from userlist where user_id = '" + userid + "'")
	_, errRelUserDep := tx.Exec("delete from rel_user_dep where user_id = '" + userid + "'")
	_, errRelUserSysDept := tx.Exec("delete from rel_user_sys_dept where user_id = '" + userid + "'")
	_, errSysManager := tx.Exec("delete from sys_manager where user_id = '" + userid + "'")
	_, errRelUserRole := tx.Exec("delete from rel_user_role where user_id = '" + userid + "'")

	if (errUserList != nil) ||
		(errRelUserDep != nil) ||
		(errRelUserSysDept != nil) ||
		(errSysManager != nil) ||
		(errRelUserRole != nil) {

		tx.Rollback()
		res["info"] = "执行失败"
		httpjsondone.SendRes(w, nil, res, template)
		return
	}
	tx.Commit()

	res["stat"] = "true"
	res["info"] = "删除成功"
	httpjsondone.SendRes(w, nil, res, template)
}
