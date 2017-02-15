package routes

import (
	"net/http"
	"strings"
)

func JsonIp(path string) {

	subrouter := prouter.PathPrefix(path).Subrouter()

	subrouter.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		arr := strings.Split(r.RemoteAddr, ":")
		ip, port := arr[0], arr[1]
		r.ParseForm()
		callback := r.Form.Get("callback")
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.Write([]byte(callback + "({\"ip\":\"" + ip + "\", \"port\":\"" + port + "\"})"))
	})

}
