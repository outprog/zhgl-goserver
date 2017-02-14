package class

import (
	"log"
	"net/http"

	"github.com/satori/go.uuid"

	"zhgl-goserver/lib/condb"
	"zhgl-goserver/lib/httpjsondone"
)

// 新增类别
func Add(w http.ResponseWriter, r *http.Request) {
	db := condb.OpenApp()
	defer db.Close()

	res := map[string]string{
		"stat": "false",
		"info": "错误的输入格式",
	}
	template := map[string]string{
		"name":       "",
		"parent_seq": "",
	}

	body := httpjsondone.GetBody(r)
	if body["name"] == "" {
		httpjsondone.SendRes(w, nil, res, template)
		return
	}

	name := body["name"]
	parentseq := body["parent_seq"]

	log.Println("portal/admin add class named:", name)

	// 生成id，获取新的seq
	id := uuid.NewV4().String()
	var seq string
	rows, _ := db.Query("select concat(ifnull(max(t.seq)+1, concat('" + parentseq + "', '100')), '') from app.portal_post_class t where t.seq like '" + parentseq + "%' and length(t.seq) = length('" + parentseq + "') + 3")
	defer rows.Close()
	if rows.Next() {
		if err := rows.Scan(&seq); err != nil {
			log.Println(err)
		}
	}

	sql := "insert into app.portal_post_class values ('" +
		id + "', '" +
		name + "', '" +
		seq + "')"
	stmt, _ := db.Prepare(sql)
	_, err := stmt.Exec()
	defer stmt.Close()
	if err != nil {
		res["info"] = err.Error()
	} else {
		res["stat"] = "true"
		res["info"] = "添加成功"
	}

	data := []map[string]string{{
		"uuid": id,
		"id":   seq,
	}}

	httpjsondone.SendRes(w, data, res, template)
}
