package iplist

import (
	"log"
	"net/http"

	"zhgl-goserver/lib/condb"
	"zhgl-goserver/lib/httpjsondone"
)

// 新增ip
func Add(w http.ResponseWriter, r *http.Request) {
	db := condb.OpenApp()
	defer db.Close()

	res := map[string]string{
		"stat": "false",
		"info": "错误的输入格式",
	}
	template := map[string]string{
		"ip":      "",
		"mac":     "",
		"user_id": "",
		"seq":     "",
		"remark":  "",
	}

	body := httpjsondone.GetBody(r)
	if body["ip"] == "" {
		httpjsondone.SendRes(w, nil, res, template)
		return
	}

	ip := body["ip"]
	mac := body["mac"]
	userid := body["user_id"]
	seq := body["seq"]
	remark := body["remark"]

	log.Println("itsm iplist add ip")

	sql := "insert into app.itsm_ip_list values ('" +
		ip + "', '" +
		mac + "', '" +
		userid + "', '" +
		seq + "', '" +
		remark + "')"
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
