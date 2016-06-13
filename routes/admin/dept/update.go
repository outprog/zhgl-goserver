package dept

import (
	"log"
	"net/http"

	"zhgl-goserver/lib/condb"
	"zhgl-goserver/lib/httpjsondone"
)

// 更新部门信息
func Update(w http.ResponseWriter, r *http.Request) {
	db := condb.Open()
	defer db.Close()

	res := map[string]string{
		"stat": "false",
		"info": "错误的输入格式",
	}
	template := map[string]string{
		"dept_id":    "",
		"dept_name":  "",
		"mdept_id":   "",
		"dept_class": "",
	}

	body := httpjsondone.GetBody(r)
	if (body["dept_id"] == "") || (body["dept_name"] == "") || (body["mdept_id"] == "") || (body["dept_class"] == "") {
		httpjsondone.SendRes(w, nil, res, template)
		return
	}

	deptid := body["dept_id"]
	deptname := body["dept_name"]
	mdeptid := body["mdept_id"]
	deptclass := body["dept_class"]

	log.Println("update dept:", deptid)

	sql := "update mis.department t  " +
		"	set t.DEPT_NAME = '" + deptname + "', " +
		"		t.MDEPT_ID = '" + mdeptid + "', " +
		"		t.DEPT_CLASS = '" + deptclass + "' " +
		"	where t.DEPT_ID = '" + deptid + "'"
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
