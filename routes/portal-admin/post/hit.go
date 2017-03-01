package post

import (
	"log"
	"net/http"

	"zhgl-goserver/lib/condb"
	"zhgl-goserver/lib/httpjsondone"
)

// 增加文章点击次数
func Hit(w http.ResponseWriter, r *http.Request) {
	db := condb.OpenApp()
	defer db.Close()

	res := map[string]string{
		"stat": "false",
		"info": "错误的输入格式",
	}
	template := map[string]string{
		"id": "",
	}

	body := httpjsondone.GetBody(r)
	if body["id"] == "" {
		httpjsondone.SendRes(w, nil, res, template)
		return
	}

	id := body["id"]

	log.Println("portal add post hits")

	sql := "update app.portal_post set " +
		"`hits` = `hits` + 1 " +
		"where `id` =  '" + id + "'"

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
