package admin

import (
	"net/http"

	"zhgl-goserver/routes/admin/user"
)

func UserSubrouter(path string) {

	subrouter := prouter.PathPrefix(path).Subrouter()

	subrouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("user\n"))
	})

	// 新增用户
	subrouter.HandleFunc("/add", user.Add)
	// 删除用户
	subrouter.HandleFunc("/del", user.Del)
	// 修改用户
	subrouter.HandleFunc("/update", user.Update)
	// 查询用户
	subrouter.HandleFunc("/query", user.Query)
	subrouter.HandleFunc("/query/{userid}", user.QueryGet)
	// 设置用户部门
	subrouter.HandleFunc("/setting-dept", user.SettingDept)
	// 管理用户权限
	subrouter.HandleFunc("/add-role", user.AddRole)
	subrouter.HandleFunc("/del-role", user.DelRole)
	// 管理用户应用系统
	subrouter.HandleFunc("/add-app", user.AddApp)
	subrouter.HandleFunc("/del-app", user.DelApp)
	// 验证密码并获取用户信息
	subrouter.HandleFunc("/confirm-passwd", user.ConfirmPasswd)
	// 修改密码
	subrouter.HandleFunc("/update-passwd", user.UpdatePasswd)
	// 重置密码
	subrouter.HandleFunc("/reset-passwd", user.ResetPasswd)

}
