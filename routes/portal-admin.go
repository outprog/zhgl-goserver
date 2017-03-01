package routes

import (
	"net/http"

	"zhgl-goserver/routes/portal-admin"
)

func PortalAdminSubrouter(path string) {

	subrouter := prouter.PathPrefix(path).Subrouter()

	subrouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("portal-admin\n"))
	})

	// 初始化子路由全局变量
	portalAdmin.Init(subrouter)
	// 子路由配置
	portalAdmin.PostSubrouter("/post")
	portalAdmin.ClassSubrouter("/class")
	portalAdmin.VisitorSubrouter("/visitor")

}
