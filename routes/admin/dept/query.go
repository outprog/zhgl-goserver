package dept

import (
	"log"
	"net/http"

	"github.com/elgs/gosqljson"

	"zhgl-goserver/lib/condb"
	"zhgl-goserver/lib/httpjsondone"
)

// 查询部门
func Query(w http.ResponseWriter, r *http.Request) {
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
		"mdept_name": "",
		"dept_class": "",
	}

	body := httpjsondone.GetBody(r)
	deptid := body["dept_id"]
	deptname := body["dept_name"]
	mdeptid := body["mdept_id"]
	mdeptname := body["mdept_name"]
	deptclass := body["dept_class"]

	log.Println("query department")

	sql := "select t.DEPT_ID, t.DEPT_NAME, t.MDEPT_ID, t1.DEPT_NAME as MDEPT_NAME, t.DEPT_CLASS  " +
		"	from mis.department t left join mis.department t1 on t.MDEPT_ID = t1.DEPT_ID " +
		"	where t.DEPT_ID like '%" + deptid + "%' " +
		"	and t.DEPT_NAME like '%" + deptname + "%' " +
		"	and t.MDEPT_ID like '%" + mdeptid + "%' " +
		"	and t1.DEPT_NAME like '%" + mdeptname + "%' " +
		"	and t.DEPT_CLASS like '%" + deptclass + "%' " +
		"	order by t.DEPT_ID"
	data, _ := gosqljson.QueryDbToMap(db, "upper", sql)

	res["stat"] = "true"
	res["info"] = "查询成功"
	httpjsondone.SendRes(w, data, res, template)
}
