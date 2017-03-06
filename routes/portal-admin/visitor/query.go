package visitor

import (
	"log"
	"net/http"

	"github.com/elgs/gosqljson"

	"zhgl-goserver/lib/condb"
	"zhgl-goserver/lib/httpjsondone"
)

// 查询分类
func Query(w http.ResponseWriter, r *http.Request) {
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

	log.Println("portal/visitor query visitor")

	sql := "select min(added_date) as added_date, " +
		"(select v.user_name from mis.userlist v where v.user_id = tt.user_id) as user_name, " +
		"(select v.dept_name from mis.department v where v.dept_id = (select vv.dept_id from mis.rel_user_dep vv where vv.user_id = tt.user_id)) as dept_name, " +
		"if(ifnull(tt.seq, '') = '', '9999', tt.seq) as seq " +
		"from " +
		"	(select (select v.user_id from app.itsm_ip_list v where v.ip = t.ip) as user_id,  " +
		"	min(added_date) as added_date, " +
		"	(select v.seq from app.itsm_ip_list v where v.ip = t.ip) as seq " +
		"	from portal_post_visitor t " +
		"	where id = '" + id + "' " +
		"	group by ip " +
		"	union  " +
		"	select user_id, min(added_date) as added_date, " +
		"	(select v.seq from app.itsm_ip_list v where v.user_id = t.user_id) as seq " +
		"	from portal_post_visitor t " +
		"	where id = '" + id + "' and user_id is not null and user_id != ''  " +
		"	group by user_id) tt " +
		"where user_id is not null and user_id != '' " +
		"group by user_id " +
		"order by seq "

	data, _ := gosqljson.QueryDbToMap(db, "upper", sql)

	res["stat"] = "true"
	res["info"] = "查询成功"
	httpjsondone.SendRes(w, data, res, template)
}
