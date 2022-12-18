// package dbs

// import (
//     "fmt"
//     "database/sql"
//     _ "github.com/go-sql-driver/mysql"
//     "log"
//     // "context"
//     // "time"
// )
// // var db *sql.DB
// // func Database()*sql.DB {
// //     db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")

// //     if err != nil {
// //         log.Printf("Error %s when opening DB\n", err)
// //         panic(err.Error())
// //     }else{
// //     fmt.Println(" MySQL connected",db)
// //     }
// //     return db
// // }
// // const (
// // 	username = "root"
// // 	password = ""
// // 	hostname = "127.0.0.1:3306"
// // 	dbname   = "test"
// // )

// // var db *sql.DB

// // func init() {
// // 	db = prepareDb(dbname)
// // 	defer db.Close()

// // }

// // func dsn(dbName string) string {
// // 	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
// // }

// // func connectToDb(db *sql.DB) {
// // 	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
// // 	defer cancelFunc()
// // 	res, err := db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbname)
// // 	if err != nil {
// // 		log.Printf("error % s when creating db\n", err)
// // 		return
// // 	}
// // 	no, err := res.RowsAffected()

// // 	if err != nil {
// // 		log.Printf("error %s when fetching db\n", err)
// // 	}
// // 	log.Printf("rows affected %d\n", no)
// // }

// // func prepareDb(dbname string) *sql.DB {
// // 	db, err := sql.Open("mysql", dsn(dbname))
// // 	if err != nil {
// // 		log.Printf("error %s during the open db\n", err)
// // 	}
// // 	connectToDb(db)
// // 	return db
// // }



// var db *sql.DB

// func init() {

// 	db = prepareDb()
// 	defer db.Close()

// }
// func prepareDb() *sql.DB {
//     db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")

//     if err != nil {
//         log.Printf("Error %s when opening DB\n", err)
//         panic(err.Error())
//     }else{
//     fmt.Println(" MySQL connected",db)
//     }
//     return db
// }

package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	checkErr(err)

	//fmt.Println("You connected to your database.")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}