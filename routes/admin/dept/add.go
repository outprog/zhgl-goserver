package dept

import (
	"log"
	"net/http"

	"zhgl-goserver/lib/condb"
	"zhgl-goserver/lib/httpjsondone"
)

// 添加部门
func Add(w http.ResponseWriter, r *http.Request) {
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

	log.Println("add dept:", deptid)

	stmt, _ := db.Prepare("insert into department (dept_id, dept_name, mdept_id, dept_class) values ('" + deptid + "', '" + deptname + "', '" + mdeptid + "', '" + deptclass + "')")
	_, err := stmt.Exec()
	stmt.Close()

	if err != nil {
		res["stat"] = "false"
		res["info"] = err.Error()
	} else {
		res["stat"] = "true"
		res["info"] = "部门添加成功"
	}

	httpjsondone.SendRes(w, nil, res, template)
}
