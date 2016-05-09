package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"

	"zhgl-goserver/routes"
)

func main() {

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
	routes.AdminSubrouter(services)

	// Bind to a port and pass our router in
	fmt.Println("services start :8080")
	http.ListenAndServe(":8000", r)
}
