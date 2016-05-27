package role

import (
	"log"
	"net/http"

	"github.com/elgs/gosqljson"

	"zhgl-goserver/lib/httpjsondone"
)

// 查询权限
func Query(w http.ResponseWriter, r *http.Request) {
	res := map[string]string{
		"stat": "false",
		"info": "错误的输入格式",
	}
	template := map[string]string{
		"role_id":     "",
		"role_name":   "",
		"user_id":     "",
		"user_name":   "",
		"operator_id": "",
		"sys_id":      "",
		"sys_name":    "",
		"menu_id":     "",
		"menu_name":   "",
	}

	body := httpjsondone.GetBody(r)
	roleid := body["role_id"]
	rolename := body["role_name"]
	userid := body["user_id"]
	username := body["user_name"]
	operatorid := body["operator_id"]
	sysid := body["sys_id"]
	sysname := body["sys_name"]
	menuid := body["menu_id"]
	menuname := body["menu_name"]

	log.Println("query role")

	sql := "select * from mis.rolelist t " +
		"	where ('" + roleid + "' = '' or t.ROLE_ID = '" + roleid + "') " +
		"	and ('" + rolename + "' = '' or t.ROLE_NAME like '%" + rolename + "%') " +
		"	and ('" + userid + "' = '' or t.ROLE_ID in (select role_id from mis.rel_user_role where user_id like '%" + userid + "%')) " +
		"	and ('" + username + "' = '' or t.ROLE_ID in (select role_id from mis.rel_user_role where user_id in (select user_id from mis.userlist where user_name like '%" + username + "%'))) " +
		"	and ('" + operatorid + "' = '' or t.ROLE_ID in (select role_id from mis.rel_sys_role_menu where sys_id in (select sys_id from mis.sys_manager where user_id like '%" + operatorid + "%'))) " +
		"	and ('" + sysid + "' = '' or t.ROLE_ID in (select role_id from mis.rel_sys_role_menu where sys_id like '%" + sysid + "%')) " +
		"	and ('" + sysname + "' = '' or t.ROLE_ID in (select role_id from mis.rel_sys_role_menu where sys_id in (select sys_id from mis.syslist where sys_name like '%" + sysname + "%'))) " +
		"	and ('" + menuid + "' = '' or t.ROLE_ID in (select role_id from mis.rel_sys_role_menu where menu_id like '%" + menuid + "%')) " +
		"	and ('" + menuname + "' = '' or t.ROLE_ID in (select role_id from mis.rel_sys_role_menu where menu_id in (select menu_id from mis.menulist where menu_name like '%" + menuname + "%')))"

	data, _ := gosqljson.QueryDbToMap(db, "upper", sql)

	res["stat"] = "true"
	res["info"] = "查询成功"
	httpjsondone.SendRes(w, data, res, template)
}
