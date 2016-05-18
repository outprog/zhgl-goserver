package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	//_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-oci8"

	"zhgl-goserver/routes"
)

func main() {
	// set log
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime)

	// connect db
	nlsLang := os.Getenv("NLS_LANG")
	if !strings.HasSuffix(nlsLang, "UTF8") {
		i := strings.LastIndex(nlsLang, ".")
		if i < 0 {
			os.Setenv("NLS_LANG", "AMERICAN_AMERICA.AL32UTF8")
		} else {
			nlsLang = nlsLang[:i+1] + "AL32UTF8"
			fmt.Fprintf(os.Stderr, "NLS_LANG error: should be %s, not %s!\n",
				nlsLang, os.Getenv("NLS_LANG"))
		}
	}

	db, err := sql.Open("oci8", getDSN())
	//db, err := sql.Open("mysql", "root@/mis")
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

func getDSN() string {
	var dsn string
	if len(os.Args) > 1 {
		dsn = os.Args[1]
		if dsn != "" {
			return dsn
		}
	}
	dsn = os.Getenv("GO_OCI8_CONNECT_STRING")
	if dsn != "" {
		return dsn
	}
	fmt.Fprintln(os.Stderr, `Please specifiy connection parameter in GO_OCI8_CONNECT_STRING environment variable,
or as the first argument! (The format is user/name@host:port/sid)`)
	//return "scott/tiger@XE"
	return "scott/tiger@XE"
}
