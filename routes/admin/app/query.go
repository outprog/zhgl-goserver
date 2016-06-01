package app

import (
	"log"
	"net/http"

	"github.com/elgs/gosqljson"

	"zhgl-goserver/lib/httpjsondone"
)

// 查询系统
func Query(w http.ResponseWriter, r *http.Request) {
	res := map[string]string{
		"stat": "false",
		"info": "错误的输入格式",
	}
	template := map[string]string{
		"sys_id":       "",
		"sys_name":     "",
		"sys_url":      "",
		"sys_flag":     "",
		"manager_id":   "",
		"manager_name": "",
	}

	body := httpjsondone.GetBody(r)
	sysid := body["sys_id"]
	sysname := body["sys_name"]
	sysurl := body["sys_url"]
	sysflag := body["sys_flag"]
	managerid := body["manager_id"]
	managername := body["manager_name"]

	log.Println("query app")

	sql := "select * from mis.syslist t " +
		"	where ('" + sysid + "' is null or '" + sysid + "' = '' or t.SYS_ID = '" + sysid + "') " +
		"	and ('" + sysname + "' is null or '" + sysname + "' = '' or t.SYS_NAME like '%" + sysname + "%') " +
		"	and ('" + sysurl + "' is null or '" + sysurl + "' = '' or t.SYS_URL like '%" + sysurl + "%') " +
		"	and ('" + sysflag + "' is null or '" + sysflag + "' = '' or t.SYS_FLAG like '%" + sysflag + "%') " +
		"	and ('" + managerid + "' is null or '" + managerid + "' = '' or t.SYS_ID in (select sys_id from mis.sys_manager where user_id like '%" + managerid + "%')) " +
		"	and ('" + managername + "' is null or '" + managername + "' = '' or t.SYS_ID in (select sys_id from mis.sys_manager where user_id in (select user_id from mis.userlist where user_name like '%" + managername + "%'))) "

	data, _ := gosqljson.QueryDbToMap(db, "upper", sql)

	res["stat"] = "true"
	res["info"] = "查询成功"
	httpjsondone.SendRes(w, data, res, template)
}
