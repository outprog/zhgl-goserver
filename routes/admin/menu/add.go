package menu

import (
	"log"
	"net/http"
	"strconv"

	"zhgl-goserver/lib/condb"
	"zhgl-goserver/lib/httpjsondone"
)

// 新增菜单
func Add(w http.ResponseWriter, r *http.Request) {
	db := condb.Open()
	defer db.Close()

	res := map[string]string{
		"stat": "false",
		"info": "错误的输入格式",
	}
	template := map[string]string{
		"sys_id":    "",
		"mmenu_seq": "",
		"menu_name": "",
		"menu_url":  "",
		"menu_pla":  "",
	}

	body := httpjsondone.GetBody(r)
	if (body["sys_id"] == "") || (body["menu_name"] == "") || (body["menu_pla"] == "") {
		httpjsondone.SendRes(w, nil, res, template)
		return
	}

	sysid := body["sys_id"]
	mmenu_seq := body["mmenu_seq"]
	menuname := body["menu_name"]
	menuurl := body["menu_url"]
	menupla := body["menu_pla"]

	log.Println("add menu named:", menuname)

	var id, seq int
	var menuid, menuseq string
	idrows, _ := db.Query("select max(t.MENU_ID) + 1 from menulist t")
	defer idrows.Close()
	if idrows.Next() {
		err := idrows.Scan(&id)
		if err != nil {
			id = 10000001
		}
	}
	menuid = strconv.Itoa(id)

	seqrows, _ := db.Query("select max(t.MENU_SEQ) + 1 from menulist t where t.SYS_ID = '" + sysid + "' and t.MENU_SEQ like '" + mmenu_seq + "___'")
	defer seqrows.Close()
	if seqrows.Next() {
		err := seqrows.Scan(&seq)
		if err != nil {
			menuseq = mmenu_seq + "101"
		} else {
			menuseq = strconv.Itoa(seq)
		}
	}

	sql := "insert into mis.menulist (menu_id, menu_name, sys_id, menu_url, menu_pla, menu_seq) values ('" +
		menuid + "', '" +
		menuname + "', '" +
		sysid + "', '" +
		menuurl + "', '" +
		menupla + "', '" +
		menuseq + "')"
	stmt, _ := db.Prepare(sql)
	_, err := stmt.Exec()
	defer stmt.Close()
	if err != nil {
		res["info"] = err.Error()
	} else {
		res["stat"] = "true"
		res["info"] = "菜单添加成功"
	}

	httpjsondone.SendRes(w, nil, res, template)
}
