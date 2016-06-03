package role

import (
	"log"
	"net/http"

	"zhgl-goserver/lib/httpjsondone"
)

// 删除权限菜单
func DelMenu(w http.ResponseWriter, r *http.Request) {

	res := map[string]string{
		"stat": "false",
		"info": "错误的输入格式",
	}
	template := map[string]string{
		"sys_id":  "",
		"role_id": "",
		"menu_id": "",
	}

	body := httpjsondone.GetBody(r)
	if (body["sys_id"] == "") || (body["role_id"] == "") || (body["menu_id"] == "") {
		httpjsondone.SendRes(w, nil, res, template)
		return
	}

	sysid := body["sys_id"]
	roleid := body["role_id"]
	menuid := body["menu_id"]

	log.Println("role", roleid, "del menu", menuid)

	sql := "delete from mis.rel_sys_role_menu " +
		"   where SYS_ID = '" + sysid + "'" +
		"   and ROLE_ID = '" + roleid + "'" +
		"   and MENU_ID = '" + menuid + "'"

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
