package condb

import (
	"database/sql"
	//	"fmt"
	"log"
	//	"os"
	//	"strings"

	_ "github.com/go-sql-driver/mysql"
	//	_ "github.com/mattn/go-oci8"
)

func Open() *sql.DB {
	// connect db
	//	nlsLang := os.Getenv("NLS_LANG")
	//	if !strings.HasSuffix(nlsLang, "UTF8") {
	//		i := strings.LastIndex(nlsLang, ".")
	//		if i < 0 {
	//			os.Setenv("NLS_LANG", "AMERICAN_AMERICA.AL32UTF8")
	//		} else {
	//			nlsLang = nlsLang[:i+1] + "AL32UTF8"
	//			fmt.Fprintf(os.Stderr, "NLS_LANG error: should be %s, not %s!\n",
	//				nlsLang, os.Getenv("NLS_LANG"))
	//		}
	//	}
	//mydb, err := sql.Open("oci8", getDSN())
	mydb, err := sql.Open("mysql", "root@/mis")
	if err != nil {
		log.Println(err.Error())
	}

	return mydb
}

//func getDSN() string {
//	var dsn string
//	if len(os.Args) > 1 {
//		dsn = os.Args[1]
//		if dsn != "" {
//			return dsn
//		}
//	}
//	dsn = os.Getenv("GO_OCI8_CONNECT_STRING")
//	if dsn != "" {
//		return dsn
//	}
//	fmt.Fprintln(os.Stderr, `Please specifiy connection parameter in GO_OCI8_CONNECT_STRING environment variable,
//or as the first argument! (The format is user/name@host:port/sid)`)
//	return "scott/tiger@XE"
//}
