package menus

import (
	"database/sql"

	"github.com/elgs/gosqljson"
	_ "github.com/go-sql-driver/mysql"
	"github.com/outprog/go-simplejson"
)

func GenSysMenu(db *sql.DB, userid string) *simplejson.Json {
	sys := GetSys(db, userid)
	jsArr, _ := simplejson.NewJson([]byte(`[]`))
	for _, v := range sys {
		js, _ := simplejson.NewJson([]byte(`{}`))
		js.Set("name", v["SYS_NAME"])
		if v["SYS_URL"] == "" {
			js.Set("href", "#")
			js.Set("sub",
				GenMenu(db, userid, v["SYS_ID"], "___").Interface())
		} else {
			js.Set("href", "#")
		}
		jsArr = simplejson.Append(jsArr, js.Interface())
	}
	return jsArr
}

func GenMenu(db *sql.DB, userid string, sysid string, menuseq string) *simplejson.Json {
	menu := GetMenu(db, userid, sysid, menuseq)
	jsArr, _ := simplejson.NewJson([]byte(`[]`))
	for _, v := range menu {
		js, _ := simplejson.NewJson([]byte(`{}`))
		js.Set("name", v["MENU_NAME"])
		js.Set("seq", v["MENU_SEQ"])
		js.Set("pla", v["MENU_PLA"])
		if v["MENU_URL"] == "" {
			js.Set("href", "#")
			js.Set("sub",
				GenMenu(db, userid, sysid, v["MENU_SEQ"]+"___").Interface())
		} else {
			js.Set("href", v["MENU_URL"])
		}
		jsArr = simplejson.Append(jsArr, js.Interface())
	}
	return jsArr
}

func GetSys(db *sql.DB, userid string) []map[string]string {
	sql := "select * from mis.syslist t" +
		"   where t.sys_id in (select sys_id from mis.rel_sys_role_menu " +
		"       where role_id in (select role_id from mis.rel_user_role where user_id = '" + userid + "')) " +
		"   order by sys_id"
	data, _ := gosqljson.QueryDbToMap(db, "upper", sql)
	return data
}

func GetMenu(db *sql.DB, userid string, sysid string, menuseq string) []map[string]string {
	sql := "select * from mis.menulist t " +
		"   where t.menu_id in  " +
		"       (select menu_id from mis.rel_sys_role_menu  " +
		"       where role_id in  " +
		"           (select role_id from mis.rel_user_role  " +
		"           where user_id = '" + userid + "'))  " +
		"   and t.sys_id = '" + sysid + "'  " +
		"   and menu_seq like '" + menuseq + "' " +
		"   order by t.menu_seq"
	data, _ := gosqljson.QueryDbToMap(db, "upper", sql)
	return data
}
