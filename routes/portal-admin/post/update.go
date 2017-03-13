package post

import (
	"log"
	"net/http"
	"strings"

	"zhgl-goserver/lib/condb"
	"zhgl-goserver/lib/httpjsondone"
)

// 更新文章
func Update(w http.ResponseWriter, r *http.Request) {
	db := condb.OpenApp()
	defer db.Close()

	res := map[string]string{
		"stat": "false",
		"info": "错误的输入格式",
	}
	template := map[string]string{
		"id":                "",
		"title":             "",
		"title_hide":        "",
		"content":           "",
		"content_hide":      "",
		"class1":            "",
		"class2":            "",
		"images":            "",
		"images_hide":       "",
		"documents":         "",
		"documents_hide":    "",
		"attachment":        "",
		"attachment_hide":   "",
		"post_dept":         "",
		"post_dept_hide":    "",
		"cover":             "",
		"last_edit_ip":      "",
		"last_edit_user_id": "",
	}

	body := httpjsondone.GetBody(r)
	if body["id"] == "" || body["title"] == "" || body["last_edit_user_id"] == "" {
		httpjsondone.SendRes(w, nil, res, template)
		return
	}

	id := body["id"]
	title := body["title"]
	titlehide := body["title_hide"]
	content := body["content"]
	contenthide := body["content_hide"]
	class1 := body["class1"]
	class2 := body["class2"]
	images := body["images"]
	imageshide := body["images_hide"]
	documents := body["documents"]
	documentshide := body["documents_hide"]
	attachment := body["attachment"]
	attachmenthide := body["attachment_hide"]
	postdept := body["post_dept"]
	postdepthide := body["post_dept_hide"]
	cover := body["cover"]
	lsip := body["last_edit_ip"]
	lsuserid := body["last_edit_user_id"]

	log.Println("portal/admin update post title:", title)

	sql := "update app.portal_post set " +
		"`id` =                 '" + id + "', " +
		"`title` =              '" + title + "', " +
		"`title_hide` =         " + titlehide + ", " +
		"`content` =            '" + strings.Replace(content, "'", "\\'", -1) + "', " +
		"`content_hide` =       " + contenthide + ", " +
		"`class1` =             '" + class1 + "', " +
		"`class2` =             '" + class2 + "', " +
		"`images` =             '" + images + "', " +
		"`images_hide` =        " + imageshide + ", " +
		"`documents` =          '" + documents + "', " +
		"`documents_hide` =     " + documentshide + ", " +
		"`attachment` =         '" + attachment + "', " +
		"`attachment_hide` =    " + attachmenthide + ", " +
		"`post_dept` =          '" + postdept + "', " +
		"`post_dept_hide` =     " + postdepthide + ", " +
		"`cover` =              '" + cover + "', " +
		"`last_edit_ip` =       '" + lsip + "', " +
		"`last_edit_user_id` =  '" + lsuserid + "' " +
		"where `id` =  '" + id + "'"

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
