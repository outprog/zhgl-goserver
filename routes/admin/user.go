package admin

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func UserSubrouter(r *mux.Router, db *sql.DB) {

	subrouter := r.PathPrefix("/user").Subrouter()

	subrouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("user\n"))
	})

	// get passwd
	subrouter.HandleFunc("/confirm-passwd/{userid}", func(w http.ResponseWriter, r *http.Request) {
		userid := mux.Vars(r)["userid"]
		rows, err := db.Query("SELECT user_name FROM userlist where user_id=?", userid)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		var USER_NAME string
		for rows.Next() {
			if err := rows.Scan(&USER_NAME); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s\n", USER_NAME)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte(USER_NAME))
	})

}
