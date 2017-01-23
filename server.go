package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"zhgl-goserver/routes"
)

func main() {
	// set log
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime)

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("zhgl-goservices\n"))
	})

	services := r.PathPrefix("/services").Subrouter()
	services.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("services\n"))
	})

	routes.Init(services)
	// services list
	routes.AdminSubrouter("/admin")
	//routes.PortalSubrouter("/portal")         // 分行门户

	// Bind to a port and pass our router in
	log.Println("services started at port:8000")
	http.ListenAndServe(":8000", r)
}
