package class

import (
	"log"
	"net/http"

	"zhgl-goserver/lib/condb"
	"zhgl-goserver/lib/httpjsondone"
)

// 更新分类名称
func Update(w http.ResponseWriter, r *http.Request) {
	db := condb.OpenApp()
	defer db.Close()

	res := map[string]string{
		"stat": "false",
		"info": "错误的输入格式",
	}
	template := map[string]string{
		"uuid": "",
		"name": "",
	}

	body := httpjsondone.GetBody(r)
	if body["uuid"] == "" || body["name"] == "" {
		httpjsondone.SendRes(w, nil, res, template)
		return
	}

	uuid := body["uuid"]
	name := body["name"]

	log.Println("portal/admin update class name:", uuid)

	sql := "update app.portal_post_class t  " +
		"	set t.name = '" + name + "'" +
		"	where t.id = '" + uuid + "'"
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
