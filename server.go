package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"zhgl-goserver/routes"
)

func main() {
	// set log
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime)

	// connect db
	db, err := sql.Open("mysql", "root@/mis")
	if err != nil {
		log.Println("database error")
	}
	defer db.Close()

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("zhgl-goservices\n"))
	})

	services := r.PathPrefix("/services").Subrouter()
	services.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("services\n"))
	})
	// services list
	routes.AdminSubrouter(services, db)

	// Bind to a port and pass our router in
	log.Println("services started at port:8000")
	http.ListenAndServe(":8000", r)
}
