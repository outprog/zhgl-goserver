package class

import (
	"log"
	"net/http"

	"github.com/elgs/gosqljson"

	"zhgl-goserver/lib/condb"
	"zhgl-goserver/lib/httpjsondone"
)

// 查询分类
func Query(w http.ResponseWriter, r *http.Request) {
	db := condb.OpenApp()
	defer db.Close()

	res := map[string]string{
		"stat": "false",
		"info": "错误的输入格式",
	}
	template := map[string]string{
		"id":   "",
		"name": "",
		"seq":  "",
	}

	body := httpjsondone.GetBody(r)
	id := body["id"]
	name := body["name"]
	seq := body["seq"]

	log.Println("query class")

	sql := "select * from app.portal_post_class t " +
		" where ('" + id + "' is null or '" + id + "' = '' or '" + id + "' = t.id) " +
		" and ('" + name + "' is null or '" + name + "' = '' or '" + name + "' = t.name) " +
		" and ('" + seq + "' is null or '" + seq + "' = '' or '" + seq + "' = t.seq) " +
		" order by t.seq"

	data, _ := gosqljson.QueryDbToMap(db, "upper", sql)

	res["stat"] = "true"
	res["info"] = "查询成功"
	httpjsondone.SendRes(w, data, res, template)
}
