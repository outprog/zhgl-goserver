package menu

import (
	"log"
	"net/http"

	"github.com/elgs/gosqljson"

	"zhgl-goserver/lib/condb"
	"zhgl-goserver/lib/httpjsondone"
)

// 查询菜单
func Query(w http.ResponseWriter, r *http.Request) {
	db := condb.Open()
	defer db.Close()

	res := map[string]string{
		"stat": "false",
		"info": "错误的输入格式",
	}
	template := map[string]string{
		"sys_id":      "",
		"sys_name":    "",
		"role_id":     "",
		"role_name":   "",
		"menu_id":     "",
		"menu_name":   "",
		"menu_url":    "",
		"menu_pla":    "",
		"menu_status": "",
		"mmenu_id":    "",
		"mmenu_name":  "",
	}

	body := httpjsondone.GetBody(r)
	sysid := body["sys_id"]
	sysname := body["sys_name"]
	roleid := body["role_id"]
	rolename := body["role_name"]
	menuid := body["menu_id"]
	menuname := body["menu_name"]
	menuurl := body["menu_url"]
	menupla := body["menu_pla"]
	menustatus := body["menu_status"]
	mmenuid := body["mmenu_id"]
	mmenuname := body["mmenu_name"]

	log.Println("query menu")

	sql := "select * from mis.menulist t " +
		"	where ('" + sysid + "' is null or '" + sysid + "' = '' or t.SYS_ID = '" + sysid + "') " +
		"	and ('" + sysname + "' is null or '" + sysname + "' = '' or t.SYS_ID in (select sys_id from mis.syslist where sys_name like '%" + sysname + "%')) " +
		"   and ('" + roleid + "' is null or '" + roleid + "' = '' or t.MENU_ID in (select menu_id from rel_sys_role_menu where role_id = '" + roleid + "'))" +
		"   and ('" + rolename + "' is null or '" + rolename + "' = '' or t.MENU_ID in (select menu_id from rel_sys_role_menu where role_id in (select role_id from rolelist where role_name like '%" + rolename + "%')))" +
		"	and ('" + menuid + "' is null or '" + menuid + "' = '' or t.MENU_ID = '" + menuid + "') " +
		"	and ('" + menuname + "' is null or '" + menuname + "' = '' or t.MENU_NAME like '%" + menuname + "%') " +
		"	and ('" + menuurl + "' is null or '" + menuurl + "' = '' or t.MENU_URL like '%" + menuurl + "%') " +
		"	and ('" + menupla + "' is null or '" + menupla + "' = '' or t.MENU_PLA like '%" + menupla + "%') " +
		"	and ('" + menustatus + "' is null or '" + menustatus + "' = '' or t.MENU_STATUS = '" + menustatus + "') " +
		"	and ('" + mmenuid + "' is null or '" + mmenuid + "' = '' or (t.SYS_ID in (select sys_id from mis.menulist where menu_id = '" + mmenuid + "') and substr(t.MENU_SEQ, 1, LENGTH(t.MENU_SEQ)-3) in (select menu_seq from mis.menulist where menu_id = '" + mmenuid + "'))) " +
		"	and ('" + mmenuname + "' is null or '" + mmenuname + "' = '' or (t.SYS_ID in (select sys_id from mis.menulist where menu_name like '%" + mmenuname + "%') and substr(t.MENU_SEQ, 1, LENGTH(t.MENU_SEQ)-3) in (select menu_seq from mis.menulist where menu_name like '%" + mmenuname + "%'))) " +
		"   order by t.sys_id, t.menu_seq"

	data, _ := gosqljson.QueryDbToMap(db, "upper", sql)

	res["stat"] = "true"
	res["info"] = "查询成功"
	httpjsondone.SendRes(w, data, res, template)
}
