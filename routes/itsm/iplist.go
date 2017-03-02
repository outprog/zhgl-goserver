package itsm

import (
	"net/http"

	"zhgl-goserver/routes/itsm/iplist"
)

func IpListSubrouter(path string) {

	subrouter := prouter.PathPrefix(path).Subrouter()

	subrouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ip list\n"))
	})

	// 新增ip
	subrouter.HandleFunc("/add", iplist.Add)
	// 删除ip
	subrouter.HandleFunc("/del", iplist.Del)
	// 更新ip
	subrouter.HandleFunc("/update", iplist.Update)
	// 查询系统
	subrouter.HandleFunc("/query", iplist.Query)

}
