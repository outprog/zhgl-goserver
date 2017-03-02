package routes

import (
	"net/http"

	"zhgl-goserver/routes/itsm"
)

func ItsmSubrouter(path string) {

	subrouter := prouter.PathPrefix(path).Subrouter()

	subrouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("itsm\n"))
	})

	// 初始化子路由全局变量
	itsm.Init(subrouter)
	// 子路由配置
	itsm.IpListSubrouter("/iplist")

}
