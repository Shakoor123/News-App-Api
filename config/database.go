
package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/newsApp")
	checkErr(err)

	fmt.Println("db connected")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}