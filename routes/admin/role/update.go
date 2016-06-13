package role

import (
	"log"
	"net/http"

	"zhgl-goserver/lib/condb"
	"zhgl-goserver/lib/httpjsondone"
)

// 更新权限
func Update(w http.ResponseWriter, r *http.Request) {
	db := condb.Open()
	defer db.Close()

	res := map[string]string{
		"stat": "false",
		"info": "错误的输入格式",
	}
	template := map[string]string{
		"role_id":   "",
		"role_name": "",
	}

	body := httpjsondone.GetBody(r)
	if body["role_id"] == "" {
		httpjsondone.SendRes(w, nil, res, template)
		return
	}

	roleid := body["role_id"]
	rolename := body["role_name"]

	log.Println("update role:", roleid)

	sql := "update mis.rolelist t  " +
		"	set t.ROLE_NAME = '" + rolename + "'" +
		"	where t.ROLE_ID = '" + roleid + "'"
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
