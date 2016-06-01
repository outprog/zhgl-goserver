package app

import (
	"log"
	"net/http"

	"zhgl-goserver/lib/httpjsondone"
)

// 更新系统
func Update(w http.ResponseWriter, r *http.Request) {

	res := map[string]string{
		"stat": "false",
		"info": "错误的输入格式",
	}
	template := map[string]string{
		"sys_id":    "",
		"sys_name":  "",
		"sys_url":   "",
		"sys_image": "",
		"sys_flag":  "",
		"sys_pict":  "",
	}

	body := httpjsondone.GetBody(r)
	if body["sys_id"] == "" {
		httpjsondone.SendRes(w, nil, res, template)
		return
	}

	sysid := body["sys_id"]
	sysname := body["sys_name"]
	sysurl := body["sys_url"]
	sysimage := body["sys_image"]
	sysflag := body["sys_flag"]
	syspict := body["sys_pict"]

	log.Println("update app:", sysid)

	sql := "update mis.syslist t  " +
		"	set t.SYS_NAME = '" + sysname + "'," +
		"       t.SYS_URL = '" + sysurl + "'," +
		"       t.SYS_IMAGE = '" + sysimage + "'," +
		"       t.SYS_FLAG = '" + sysflag + "'," +
		"       t.SYS_PICT = '" + syspict + "'" +
		"	where t.SYS_ID = '" + sysid + "'"
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
