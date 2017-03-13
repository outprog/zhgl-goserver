package post

import (
	"log"
	"net/http"

	"github.com/elgs/gosqljson"

	"zhgl-goserver/lib/condb"
	"zhgl-goserver/lib/httpjsondone"
)

// 查询有封面的post
func QueryCover(w http.ResponseWriter, r *http.Request) {
	db := condb.OpenApp()
	defer db.Close()

	res := map[string]string{
		"stat": "false",
		"info": "错误的输入格式",
	}
	template := map[string]string{
		"class1": "",
	}

	body := httpjsondone.GetBody(r)
	class1 := body["class1"]
	if body["class1"] == "" {
		httpjsondone.SendRes(w, nil, res, template)
		return
	}
	log.Println("portal/admin query cover")

	// 查询语句拼装
	sql := "select t.id, t.title, t.cover from `portal_post` t " +
		"where t.class1 = '" + class1 + "' " +
		"and t.cover is not null and t.cover != '' order by added_date limit 3 "

	data, _ := gosqljson.QueryDbToMap(db, "upper", sql)
	res["stat"] = "true"
	res["info"] = "查询成功"
	httpjsondone.SendRes(w, data, res, template)

}
