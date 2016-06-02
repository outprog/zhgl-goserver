package menu

import (
	"log"
	"net/http"

	"zhgl-goserver/lib/httpjsondone"
)

// 更新菜单
func Update(w http.ResponseWriter, r *http.Request) {

	res := map[string]string{
		"stat": "false",
		"info": "错误的输入格式",
	}
	template := map[string]string{
		"menu_id":     "",
		"menu_name":   "",
		"menu_url":    "",
		"menu_pla":    "",
		"menu_status": "",
	}

	body := httpjsondone.GetBody(r)
	if (body["menu_id"] == "") || (body["menu_name"] == "") || (body["menu_pla"] == "") {
		httpjsondone.SendRes(w, nil, res, template)
		return
	}

	menuid := body["menu_id"]
	menuname := body["menu_name"]
	menuurl := body["menu_url"]
	menupla := body["menu_pla"]
	menustatus := body["menu_status"]

	log.Println("update menu:", menuid)

	sql := "update mis.menulist t  " +
		"	set t.menu_NAME = '" + menuname + "'," +
		"       t.menu_URL = '" + menuurl + "'," +
		"       t.menu_pla = '" + menupla + "'," +
		"       t.menu_status = '" + menustatus + "'" +
		"	where t.menu_ID = '" + menuid + "'"
	stmt, _ := db.Prepare(sql)
	defer stmt.Close()
	upres, _ := stmt.Exec()
	rowCnt, err := upres.RowsAffected()

	if err != nil {
		res["stat"] = "false"
		res["info"] = err.Error()
		httpjsondone.SendRes(w, nil, res, template)
		return
	}
	if rowCnt == 1 {
		res["stat"] = "true"
		res["info"] = "更新成功"
	} else {
		res["stat"] = "false"
		res["info"] = "更新失败或没有改动"
	}

	httpjsondone.SendRes(w, nil, res, template)
}
