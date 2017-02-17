package portalAdmin

import (
	"net/http"

	"zhgl-goserver/routes/portal-admin/post"
)

func PostSubrouter(path string) {

	subrouter := prouter.PathPrefix(path).Subrouter()

	subrouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("post\n"))
	})

	// 新增文章
	subrouter.HandleFunc("/add", post.Add)
	// 删除文章
	subrouter.HandleFunc("/del", post.Del)
	// 修改文章
	//subrouter.HandleFunc("/update", post.Update)
	// 查询文章列表
	subrouter.HandleFunc("/querylist", post.QueryList)
	// 查询文章详情
	subrouter.HandleFunc("/querydetail", post.QueryDetail)

}
