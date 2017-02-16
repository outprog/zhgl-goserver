package post

import (
	"log"
	"net/http"

	"github.com/satori/go.uuid"

	"zhgl-goserver/lib/condb"
	"zhgl-goserver/lib/httpjsondone"
)

// 新增文章
func Add(w http.ResponseWriter, r *http.Request) {
	db := condb.OpenApp()
	defer db.Close()

	res := map[string]string{
		"stat": "false",
		"info": "错误的输入格式",
	}
	template := map[string]string{
		"title":           "",
		"title_hide":      "",
		"content":         "",
		"content_hide":    "",
		"class1":          "",
		"class2":          "",
		"images":          "",
		"images_hide":     "",
		"documents":       "",
		"documents_hide":  "",
		"attachment":      "",
		"attachment_hide": "",
		"post_dept":       "",
		"post_dept_hide":  "",
		"cover":           "",
		"user_id":         "",
		"dept_id":         "",
		"added_ip":        "",
	}

	body := httpjsondone.GetBody(r)
	if body["title"] == "" || body["user_id"] == "" {
		httpjsondone.SendRes(w, nil, res, template)
		return
	}

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
	userid := body["user_id"]
	deptid := body["dept_id"]
	addedip := body["added_ip"]

	log.Println("portal/admin add post title:", title)

	// 生成id
	id := uuid.NewV4().String()

	sql := "insert into app.portal_post (" +
		" `id`," +
		" `title`," +
		" `title_hide`," +
		" `content`," +
		" `content_hide`," +
		" `class1`," +
		" `class2`," +
		" `images`," +
		" `images_hide`," +
		" `documents`," +
		" `documents_hide`," +
		" `attachment`," +
		" `attachment_hide`," +
		" `post_dept`," +
		" `post_dept_hide`," +
		" `cover`," +
		" `user_id`," +
		" `dept_id`," +
		" `added_date`," +
		" `hits`, " +
		" `stat`, " +
		" `added_ip`)" +
		" values(" +
		" '" + id + "'," +
		" '" + title + "'," +
		" " + titlehide + "," +
		" '" + content + "'," +
		" " + contenthide + "," +
		" '" + class1 + "'," +
		" '" + class2 + "'," +
		" '" + images + "'," +
		" " + imageshide + "," +
		" '" + documents + "'," +
		" " + documentshide + "," +
		" '" + attachment + "'," +
		" " + attachmenthide + "," +
		" '" + postdept + "'," +
		" " + postdepthide + "," +
		" '" + cover + "'," +
		" '" + userid + "'," +
		" '" + deptid + "'," +
		" now()," +
		" 0, " +
		" '0', " +
		" '" + addedip + "')"

	stmt, _ := db.Prepare(sql)
	_, err := stmt.Exec()
	defer stmt.Close()
	if err != nil {
		res["info"] = err.Error()
	} else {
		res["stat"] = "true"
		res["info"] = "添加成功"
	}

	httpjsondone.SendRes(w, nil, res, template)
}
