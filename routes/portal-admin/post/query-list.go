package post

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/elgs/gosqljson"

	"zhgl-goserver/lib/condb"
	"zhgl-goserver/lib/httpjsondone"
)

// 查询post清单
func QueryList(w http.ResponseWriter, r *http.Request) {
	db := condb.OpenApp()
	defer db.Close()

	res := map[string]string{
		"stat": "false",
		"info": "错误的输入格式",
	}
	template := map[string]string{
		"id":      "",
		"title":   "",
		"content": "",
		"class1":  "",
		"class2":  "",
		"user_id": "",
		"dept_id": "",
		"num":     "",
	}

	body := httpjsondone.GetBody(r)
	id := body["id"]
	title := body["title"]
	content := body["content"]
	class1 := body["class1"]
	class2 := body["class2"]
	userid := body["user_id"]
	deptid := body["dept_id"]
	stat := body["stat"]
	num := "1"
	if body["num"] != "" {
		num = body["num"]
	}
	log.Println("portal/admin query list")

	// 条件语句
	where := " where ('" + id + "' is null or '" + id + "' = '' or '" + id + "' = t.id) " +
		"and ('" + title + "' is null or '" + title + "' = '' or t.title like '%" + title + "%') " +
		"and ('" + content + "' is null or '" + content + "' = '' or t.content like '%" + content + "%') " +
		"and ('" + class1 + "' is null or '" + class1 + "' = '' or '" + class1 + "' = t.class1) " +
		"and ('" + class2 + "' is null or '" + class2 + "' = '' or '" + class2 + "' = t.class2) " +
		"and ('" + userid + "' is null or '" + userid + "' = '' or '" + userid + "' = t.user_id) " +
		"and ('" + deptid + "' is null or '" + deptid + "' = '' or '" + deptid + "' = t.dept_id) " +
		"and ('" + stat + "' is null or '" + stat + "' = '' or '" + stat + "' = t.stat)  "

	// 获取结果数和总页数
	sql := "select count(*) as pages from portal_post t " + where
	countsbase, _ := gosqljson.QueryDbToMap(db, "upper", sql)
	counts, _ := strconv.ParseFloat(countsbase[0]["PAGES"], 64)
	pages := int(math.Ceil(counts / 10))

	// 获取类别列表
	sql = "select * from " +
		"(select distinct(t.class1) as class1, (select v.name from app.portal_post_class v where v.id = t.class1) as class1name " +
		"from portal_post t " +
		where + ") tt  " +
		" where class1name is not null"
	class, _ := gosqljson.QueryDbToMap(db, "upper", sql)
	classByte, _ := json.Marshal(class)

	// 查询列表
	// 设置页码
	inum, numErr := strconv.Atoi(num)
	if numErr != nil {
		res["stat"] = "false"
		res["info"] = "页码设置错误"
		httpjsondone.SendRes(w, nil, res, template)
		return
	}
	page := (inum - 1) * 10
	// 查询语句拼装
	sql = "select t.id,  " +
		"t.title,  " +
		"fnStripTags(t.content) as content,  " +
		"(select v.name from portal_post_class v where v.id = t.class1) as class1_name,  " +
		"(select v.name from portal_post_class v where v.id = t.class2) as class2_name, " +
		"(select v.user_name from mis.userlist v where v.user_id = t.user_id) as user_name,  " +
		"(select v.dept_name from mis.department v where v.dept_id =t.dept_id) as dept_name,  " +
		"t.added_date,  " +
		"t.hits " +
		"from portal_post t " +
		where +
		"order by t.added_date DESC " +
		"limit " + strconv.Itoa(page) + ", 10"

	data, _ := gosqljson.QueryDbToMap(db, "upper", sql)
	res["stat"] = "true"
	res["info"] = "{\"num\": \"" + num + "\", \"pages\": \"" + strconv.Itoa(pages) + "\", \"counts\": \"" + strconv.FormatFloat(counts, 'f', 0, 64) + "\", \"class\": " + string(classByte) + "}"
	httpjsondone.SendRes(w, data, res, template)

}
