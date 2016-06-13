package dept

import (
	"log"
	"net/http"

	"zhgl-goserver/lib/condb"
	"zhgl-goserver/lib/httpjsondone"
)

// 删除部门
func Del(w http.ResponseWriter, r *http.Request) {
	db := condb.Open()
	defer db.Close()

	res := map[string]string{
		"stat": "false",
		"info": "错误的输入格式",
	}
	template := map[string]string{
		"dept_id": "",
	}

	body := httpjsondone.GetBody(r)
	if body["dept_id"] == "" {
		httpjsondone.SendRes(w, nil, res, template)
		return
	}

	deptid := body["dept_id"]

	log.Println("delete dept:", deptid)

	tx, _ := db.Begin()
	_, errDeptList := tx.Exec("delete from department where dept_id = '" + deptid + "'")
	_, errRelUserDep := tx.Exec("delete from rel_user_dep where dept_id = '" + deptid + "'")

	if (errDeptList != nil) ||
		(errRelUserDep != nil) {

		tx.Rollback()
		res["info"] = "执行失败"
	} else {
		tx.Commit()
		res["stat"] = "true"
		res["info"] = "删除成功"
	}

	httpjsondone.SendRes(w, nil, res, template)
}
