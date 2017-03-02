package iplist

import (
	"log"
	"net/http"

	"zhgl-goserver/lib/condb"
	"zhgl-goserver/lib/httpjsondone"
)

// 更新ip信息
func Update(w http.ResponseWriter, r *http.Request) {
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

	log.Println("itsm iplist ip info update")

	sql := "update app.itsm_ip_list t  " +
		"	set t.mac = '" + mac + "'," +
		"	    t.user_id = '" + userid + "'," +
		"	    t.seq = '" + seq + "'," +
		"	    t.remark = '" + remark + "'" +
		"	where t.ip = '" + ip + "'"
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
