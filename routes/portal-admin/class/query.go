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
		"id":    "",
		"name":  "",
		"seq":   "",
		"p_seq": "",
	}

	body := httpjsondone.GetBody(r)
	id := body["id"]
	name := body["name"]
	seq := body["seq"]
	pseq := body["p_seq"]

	log.Println("portal/admin query class")

	sql := "select * from app.portal_post_class t " +
		" where ('" + id + "' is null or '" + id + "' = '' or '" + id + "' = t.id) " +
		" and ('" + name + "' is null or '" + name + "' = '' or '" + name + "' = t.name) " +
		" and ('" + seq + "' is null or '" + seq + "' = '' or '" + seq + "' = t.seq) " +
		" and ('" + pseq + "' is null or '" + pseq + "'= '' or '" + pseq + "' = '0' or (t.seq like '" + pseq + "%' and length(t.seq) = length('" + pseq + "') + 3)) " +
		" and ('" + pseq + "' != '0' or length(t.seq) = 0 + 3) " +
		" order by t.seq"

	data, _ := gosqljson.QueryDbToMap(db, "upper", sql)

	res["stat"] = "true"
	res["info"] = "查询成功"
	httpjsondone.SendRes(w, data, res, template)
}
