package class

import (
	"log"
	"net/http"

	"zhgl-goserver/lib/condb"
	"zhgl-goserver/lib/httpjsondone"
)

// 删除类别
func Del(w http.ResponseWriter, r *http.Request) {
	db := condb.OpenApp()
	defer db.Close()

	res := map[string]string{
		"stat": "false",
		"info": "错误的输入格式",
	}
	template := map[string]string{
		"seq": "",
	}

	body := httpjsondone.GetBody(r)
	if body["seq"] == "" {
		httpjsondone.SendRes(w, nil, res, template)
		return
	}

	seq := body["seq"]

	log.Println("portal/admin delete class", seq)

	sql := "delete from app.portal_post_class " +
		"   where seq like '" + seq + "%'"

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
