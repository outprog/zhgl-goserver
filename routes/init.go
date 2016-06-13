package routes

import (
	"github.com/gorilla/mux"
)

var prouter *mux.Router

func Init(r *mux.Router) {
	prouter = r

	return
}
