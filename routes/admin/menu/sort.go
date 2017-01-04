package menu

import (
	"log"
	"net/http"

	"zhgl-goserver/lib/condb"
	"zhgl-goserver/lib/httpjsondone"

	"github.com/outprog/go-simplejson"
)

// 更新菜单排序
func Sort(w http.ResponseWriter, r *http.Request) {
	db := condb.Open()
	defer db.Close()

	res := map[string]string{
		"stat": "false",
		"info": "错误的输入格式",
	}
	template := map[string]string{
		"data": "[]",
	}

	body := httpjsondone.GetBody(r)
	if body["data"] == "" {
		httpjsondone.SendRes(w, nil, res, template)
		return
	}

	data := body["data"]

	log.Println("sort menu:")

	js, errJS := simplejson.NewJson([]byte(data))
	if errJS != nil {
		res["stat"] = "false"
		res["info"] = errJS.Error()
		httpjsondone.SendRes(w, nil, res, template)
		return
	}

	tx, _ := db.Begin()
	for i, _ := range js.MustArray() {
		menu_id, _ := js.GetIndex(i).Get("menu_id").String()
		menu_seq, _ := js.GetIndex(i).Get("menu_seq").String()
		tx.Exec("update menulist set menu_seq ='" + menu_seq + "' where menu_id = '" + menu_id + "'")
	}
	tx.Commit()
	res["stat"] = "true"
	res["info"] = "执行成功"

	httpjsondone.SendRes(w, nil, res, template)
}
