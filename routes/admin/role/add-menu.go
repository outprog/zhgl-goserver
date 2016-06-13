package role

import (
	"log"
	"net/http"

	"zhgl-goserver/lib/condb"
	"zhgl-goserver/lib/httpjsondone"
)

// 添加权限菜单
func AddMenu(w http.ResponseWriter, r *http.Request) {
	db := condb.Open()
	defer db.Close()

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

	log.Println("role", roleid, "add menu", menuid)

	sql := "insert into rel_sys_role_menu (sys_id, role_id, menu_id) values ('" +
		sysid + "', '" +
		roleid + "', '" +
		menuid + "')"
	stmt, _ := db.Prepare(sql)
	_, err := stmt.Exec()
	defer stmt.Close()
	if err != nil {
		res["stat"] = "false"
		res["info"] = err.Error()
	} else {
		res["stat"] = "true"
		res["info"] = "添加成功"
	}

	httpjsondone.SendRes(w, nil, res, template)
}
