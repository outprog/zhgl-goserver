package post

import (
	"log"
	"net/http"

	"github.com/elgs/gosqljson"

	"zhgl-goserver/lib/condb"
	"zhgl-goserver/lib/httpjsondone"
)

// 查询post详情
func QueryDetail(w http.ResponseWriter, r *http.Request) {
	db := condb.OpenApp()
	defer db.Close()

	res := map[string]string{
		"stat": "false",
		"info": "错误的输入格式",
	}
	template := map[string]string{
		"id": "",
	}

	body := httpjsondone.GetBody(r)
	id := body["id"]
	if body["id"] == "" {
		httpjsondone.SendRes(w, nil, res, template)
		return
	}
	log.Println("portal/admin query post detail")

	// 查询语句拼装
	sql := "select t.*, " +
		"(select v.name from app.portal_post_class v where v.id = t.class1) as class1_name, " +
		"(select v.name from app.portal_post_class v where v.id = t.class2) as class2_name, " +
		"(select v.USER_NAME from mis.userlist v where v.USER_ID = t.user_id) as user_name, " +
		"(select v.DEPT_NAME from mis.department v where v.DEPT_ID = t.dept_id) as dept_name " +
		"from portal_post t where t.id ='" + id + "'"

	data, _ := gosqljson.QueryDbToMap(db, "upper", sql)
	res["stat"] = "true"
	res["info"] = "查询成功"
	httpjsondone.SendRes(w, data, res, template)

}
