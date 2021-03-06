package visitor

import (
	"log"
	"net/http"

	"zhgl-goserver/lib/condb"
	"zhgl-goserver/lib/httpjsondone"
)

// 新增访客
func Add(w http.ResponseWriter, r *http.Request) {
	db := condb.OpenApp()
	defer db.Close()
	res := map[string]string{
		"stat": "false",
		"info": "错误的输入格式",
	}
	template := map[string]string{
		"id":      "",
		"ip":      "",
		"user_id": "",
	}

	body := httpjsondone.GetBody(r)
	if body["id"] == "" || body["ip"] == "" {
		httpjsondone.SendRes(w, nil, res, template)
		return
	}

	id := body["id"]
	ip := body["ip"]
	userid := body["user_id"]

	log.Println("portal add visitor")

	sql := "insert into app.portal_post_visitor values ('" +
		id + "', '" +
		ip + "', '" +
		userid + "', now())"
	stmt, _ := db.Prepare(sql)
	_, err := stmt.Exec()
	defer stmt.Close()
	if err != nil {
		res["info"] = err.Error()
	} else {
		res["stat"] = "true"
		res["info"] = "添加成功"
	}

	httpjsondone.SendRes(w, nil, res, template)

}
