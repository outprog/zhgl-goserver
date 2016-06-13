package role

import (
	"log"
	"net/http"

	"zhgl-goserver/lib/condb"
	"zhgl-goserver/lib/httpjsondone"
)

// 删除权限
func Del(w http.ResponseWriter, r *http.Request) {
	db := condb.Open()
	defer db.Close()

	res := map[string]string{
		"stat": "false",
		"info": "错误的输入格式",
	}
	template := map[string]string{
		"role_id": "",
	}

	body := httpjsondone.GetBody(r)
	if body["role_id"] == "" {
		httpjsondone.SendRes(w, nil, res, template)
		return
	}

	roleid := body["role_id"]
	log.Println("delete role:", roleid)

	tx, _ := db.Begin()
	_, errRoleList := tx.Exec("delete from rolelist where role_id = '" + roleid + "'")
	_, errRelSysRoleMenu := tx.Exec("delete from rel_sys_role_menu where role_id = '" + roleid + "'")
	_, errRelUserRole := tx.Exec("delete from rel_user_role where role_id = '" + roleid + "'")

	if (errRoleList != nil) ||
		(errRelSysRoleMenu != nil) ||
		(errRelUserRole != nil) {

		tx.Rollback()
		res["info"] = "执行失败"
	} else {
		tx.Commit()
		res["stat"] = "true"
		res["info"] = "删除成功"
	}

	httpjsondone.SendRes(w, nil, res, template)
}
