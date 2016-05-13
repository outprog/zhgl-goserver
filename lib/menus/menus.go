package menus

import (
	"database/sql"
	"strings"

	"github.com/bitly/go-simplejson"
	"github.com/elgs/gosqljson"
	_ "github.com/go-sql-driver/mysql"
)

func GenSysMenu(db *sql.DB, userid string) []byte {
	sys := GetSys(db, userid)
	js := GenJsonArray(len(sys))
	for i, v := range sys {
		js.GetIndex(i).Set("name", v["SYS_NAME"])
		if v["SYS_URL"] == "" {
			js.GetIndex(i).Set("href", "#")
			js.GetIndex(i).Set("sub",
				GenMenu(db, userid, v["SYS_ID"], "___"))
		} else {
			js.GetIndex(i).Set("href", "#")
		}
	}
	bt, _ := js.MarshalJSON()
	return bt
}

func GenMenu(db *sql.DB, userid string, sysid string, menuseq string) *simplejson.Json {
	menu := GetMenu(db, userid, sysid, menuseq)
	js := GenJsonArray(len(menu))
	for i, v := range menu {
		js.GetIndex(i).Set("name", v["MENU_NAME"])
		js.GetIndex(i).Set("seq", v["MENU_SEQ"])
		js.GetIndex(i).Set("pla", v["MENU_PLA"])
		if v["MENU_URL"] == "" {
			js.GetIndex(i).Set("href", "#")
			js.GetIndex(i).Set("sub",
				GenMenu(db, userid, sysid, v["MENU_SEQ"]+"___"))
		} else {
			js.GetIndex(i).Set("href", v["MENU_URL"])
		}
	}
	return js
}

func GetSys(db *sql.DB, userid string) []map[string]string {
	sql := "select * from mis.syslist t" +
		"   where t.sys_id in (select sys_id from mis.rel_sys_role_menu " +
		"       where role_id in (select role_id from mis.rel_user_role where user_id = ?)) " +
		"   order by sys_id"
	data, _ := gosqljson.QueryDbToMap(db, "upper", sql, userid)
	return data
}

func GetMenu(db *sql.DB, userid string, sysid string, menuseq string) []map[string]string {
	sql := "select * from mis.menulist t " +
		"   where t.menu_id in  " +
		"       (select menu_id from mis.rel_sys_role_menu  " +
		"       where role_id in  " +
		"           (select role_id from mis.rel_user_role  " +
		"           where user_id = ?))  " +
		"   and t.sys_id = ?  " +
		"   and menu_seq like ? " +
		"   order by t.menu_seq"
	data, _ := gosqljson.QueryDbToMap(db, "upper", sql, userid, sysid, menuseq)
	return data
}

func GenJsonArray(len int) *simplejson.Json {
	arr := []string{}
	for i := 0; i < len; i++ {
		arr = append(arr, "{}")
	}
	str := "[" + strings.Join(arr, ",") + "]"
	js, _ := simplejson.NewJson([]byte(str))
	return js
}
