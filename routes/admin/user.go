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

	user.Init(db)
	// 新增用户
	subrouter.HandleFunc("/add", user.Add)
	// 删除用户
	subrouter.HandleFunc("/del", user.Del)
	// 查询用户
	subrouter.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
	})
	// 设置用户部门
	subrouter.HandleFunc("/setting-dept", user.SettingDept)
	// 验证密码并获取用户信息
	subrouter.HandleFunc("/confirm-passwd", user.ConfirmPasswd)

}
