package iplist

import (
	"log"
	"net/http"

	"github.com/elgs/gosqljson"

	"zhgl-goserver/lib/condb"
	"zhgl-goserver/lib/httpjsondone"
)

// 查询ip info
func Query(w http.ResponseWriter, r *http.Request) {
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
	ip := body["ip"]
	mac := body["mac"]
	userid := body["user_id"]
	seq := body["seq"]
	remark := body["remark"]

	log.Println("itsm iplist query")

	sql := "select * from app.itsm_ip_list t " +
		" where ('" + ip + "' is null or '" + ip + "' = '' or '" + ip + "' = t.ip) " +
		" and ('" + mac + "' is null or '" + mac + "' = '' or '" + mac + "' = t.mac) " +
		" and ('" + userid + "' is null or '" + userid + "' = '' or '" + userid + "' = t.user_id) " +
		" and ('" + seq + "' is null or '" + seq + "' = '' or '" + seq + "' = t.seq) " +
		" and ('" + remark + "' is null or '" + remark + "' = '' or '" + remark + "' = t.remark) " +
		" order by t.seq"

	data, _ := gosqljson.QueryDbToMap(db, "upper", sql)

	res["stat"] = "true"
	res["info"] = "查询成功"
	httpjsondone.SendRes(w, data, res, template)
}
