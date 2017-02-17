package post

import (
	"log"
	"net/http"

	"github.com/elgs/gosqljson"

	"zhgl-goserver/lib/condb"
	"zhgl-goserver/lib/httpjsondone"
)

// 查询post详情
func QueryDetail(w http.ResponseWriter, r *http.Request) {
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
	id := body["id"]
	if body["id"] == "" {
		httpjsondone.SendRes(w, nil, res, template)
		return
	}
	log.Println("portal/admin query post detail")

	// 查询语句拼装
	sql := "select * from portal_post t where t.id ='" + id + "'"

	data, _ := gosqljson.QueryDbToMap(db, "upper", sql)
	res["stat"] = "true"
	res["info"] = "查询成功"
	httpjsondone.SendRes(w, data, res, template)

}
