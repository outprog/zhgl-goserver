package menu

import (
	"log"
	"net/http"

	"zhgl-goserver/lib/httpjsondone"
)

// 删除菜单
func Del(w http.ResponseWriter, r *http.Request) {

	res := map[string]string{
		"stat": "false",
		"info": "错误的输入格式",
	}
	template := map[string]string{
		"sys_id":   "",
		"menu_seq": "",
	}

	body := httpjsondone.GetBody(r)
	if (body["sys_id"] == "") || (body["menu_seq"] == "") {
		httpjsondone.SendRes(w, nil, res, template)
		return
	}

	sysid := body["sys_id"]
	menuseq := body["menu_seq"]
	log.Println("delete menu, sysid is:", sysid, "menuseq is:", menuseq)

	tx, _ := db.Begin()
	_, errMenuList := tx.Exec("delete from menulist where sys_id = '" + sysid + "' and menu_seq like '" + menuseq + "%'")
	_, errRelSysRoleMenu := tx.Exec("delete from rel_sys_role_menu where menu_id not in (select menu_id from menulist)")

	if (errMenuList != nil) ||
		(errRelSysRoleMenu != nil) {

		tx.Rollback()
		res["info"] = "执行失败"
	} else {
		tx.Commit()
		res["stat"] = "true"
		res["info"] = "删除成功"
	}

	httpjsondone.SendRes(w, nil, res, template)
}
