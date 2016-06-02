package app

import (
	"log"
	"net/http"
	"strconv"

	"zhgl-goserver/lib/httpjsondone"
)

// 新增系统
func Add(w http.ResponseWriter, r *http.Request) {
	res := map[string]string{
		"stat": "false",
		"info": "错误的输入格式",
	}
	template := map[string]string{
		"sys_name":  "",
		"sys_url":   "",
		"sys_image": "",
		"sys_flag":  "",
		"sys_pict":  "",
	}

	body := httpjsondone.GetBody(r)
	if (body["sys_name"] == "") || (body["sys_flag"] == "") {
		httpjsondone.SendRes(w, nil, res, template)
		return
	}

	sysname := body["sys_name"]
	sysurl := body["sys_url"]
	sysimage := body["sys_image"]
	sysflag := body["sys_flag"]
	syspict := body["sys_pict"]

	log.Println("add app named:", sysname)

	var id int
	rows, _ := db.Query("select max(SYS_ID) + 1 from mis.syslist")
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			id = 100001
		}
	}

	sql := "insert into mis.syslist (sys_id, sys_name, sys_url, sys_image, sys_flag, sys_visit, sys_pict) values ('" +
		strconv.Itoa(id) + "', '" +
		sysname + "', '" +
		sysurl + "', '" +
		sysimage + "', '" +
		sysflag + "', 0, '" +
		syspict + "')"
	stmt, _ := db.Prepare(sql)
	_, err := stmt.Exec()
	defer stmt.Close()
	if err != nil {
		res["info"] = err.Error()
	} else {
		res["stat"] = "true"
		res["info"] = "系统添加成功"
	}

	httpjsondone.SendRes(w, nil, res, template)
}
