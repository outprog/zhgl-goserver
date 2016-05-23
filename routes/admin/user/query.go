package user

import (
	"log"
	"net/http"

	"github.com/elgs/gosqljson"

	"zhgl-goserver/lib/httpjsondone"
)

// 查询用户
func Query(w http.ResponseWriter, r *http.Request) {
	res := map[string]string{
		"stat": "false",
		"info": "错误的输入格式",
	}
	template := map[string]string{
		"user_id":     "",
		"user_status": "",
		"dept_id":     "",
		"role_id":     "",
		"sys_id":      "",
	}

	body := httpjsondone.GetBody(r)
	userid := body["user_id"]
	userstatus := body["user_status"]
	deptid := body["dept_id"]
	roleid := body["role_id"]
	sysid := body["sys_id"]

	log.Println("query user")

	sql := "SELECT t.user_id, t.user_name, t.user_status, t.tellerno, " +
		"          t.user_email, t.user_mobile, t.user_lastlogin, " +
		"          t2.dept_id, t3.dept_name " +
		"   FROM mis.userlist t LEFT JOIN mis.rel_user_dep t2 " +
		"   ON t.user_id = t2.user_id LEFT JOIN mis.department t3 " +
		"   ON t2.dept_id = t3.dept_id " +
		"    WHERE t.user_id LIKE '%" + userid + "%' " +
		"    AND t.user_status LIKE '%" + userstatus + "%' "
	if deptid != "" {
		sql = sql +
			"    AND t2.dept_id LIKE '%" + deptid + "%' "
	}
	if roleid != "" {
		sql = sql +
			"    AND t.user_id IN (SELECT user_id FROM mis.rel_user_role WHERE role_id LIKE '%" + roleid + "%') "
	}
	if sysid != "" {
		sql = sql +
			"    AND t.user_id IN (SELECT user_id FROM mis.sys_manager WHERE sys_id LIKE '%" + sysid + "%') "
	}
	data, _ := gosqljson.QueryDbToMap(db, "upper", sql)

	res["stat"] = "true"
	res["info"] = "查询成功"
	httpjsondone.SendRes(w, data, res, template)
}
