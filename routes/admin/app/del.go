package app

import (
	"log"
	"net/http"

	"zhgl-goserver/lib/httpjsondone"
)

// 删除系统
func Del(w http.ResponseWriter, r *http.Request) {

	res := map[string]string{
		"stat": "false",
		"info": "错误的输入格式",
	}
	template := map[string]string{
		"sys_id": "",
	}

	body := httpjsondone.GetBody(r)
	if body["sys_id"] == "" {
		httpjsondone.SendRes(w, nil, res, template)
		return
	}

	sysid := body["sys_id"]
	log.Println("delete sys:", sysid)

	tx, _ := db.Begin()
	_, errSysList := tx.Exec("delete from mis.syslist where sys_id = '" + sysid + "'")
	_, errSysManager := tx.Exec("delete from mis.sys_manager where sys_id = '" + sysid + "'")
	_, errMenuList := tx.Exec("delete from mis.menulist where sys_id = '" + sysid + "'")
	_, errRelSysRoleMenu := tx.Exec("delete from rel_sys_role_menu where sys_id = '" + sysid + "'")
	_, errRelSysSys := tx.Exec("delete from rel_sys_sys where sys_id = '" + sysid + "'")
	_, errRelUserSysDept := tx.Exec("delete from rel_user_sys_dept where sys_id = '" + sysid + "'")

	if (errSysList != nil) ||
		(errSysManager != nil) ||
		(errMenuList != nil) ||
		(errRelSysRoleMenu != nil) ||
		(errRelSysSys != nil) ||
		(errRelUserSysDept != nil) {

		tx.Rollback()
		res["info"] = "执行失败"
	} else {
		tx.Commit()
		res["stat"] = "true"
		res["info"] = "删除成功"
	}

	httpjsondone.SendRes(w, nil, res, template)
}
