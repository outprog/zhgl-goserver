package role

import (
	"log"
	"net/http"
	"strconv"

	"zhgl-goserver/lib/condb"
	"zhgl-goserver/lib/httpjsondone"
)

// 新增权限
func Add(w http.ResponseWriter, r *http.Request) {
	db := condb.Open()
	defer db.Close()

	res := map[string]string{
		"stat": "false",
		"info": "错误的输入格式",
	}
	template := map[string]string{
		"role_name": "",
	}

	body := httpjsondone.GetBody(r)
	if body["role_name"] == "" {
		httpjsondone.SendRes(w, nil, res, template)
		return
	}

	rolename := body["role_name"]

	log.Println("add role named:", rolename)

	var id int
	rows, _ := db.Query("select max(ROLE_ID) + 1 from mis.rolelist")
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			id = 101
		}
	}

	stmt, _ := db.Prepare("insert into mis.rolelist (role_id, role_name) values ('" + strconv.Itoa(id) + "', '" + rolename + "')")
	_, err := stmt.Exec()
	defer stmt.Close()
	if err != nil {
		res["info"] = err.Error()
	} else {
		res["stat"] = "true"
		res["info"] = "权限添加成功"
	}

	httpjsondone.SendRes(w, nil, res, template)
}
