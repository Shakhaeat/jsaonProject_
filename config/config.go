package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var Con *sql.DB

func Connect() (*sql.DB, error) {

	//db, err := sql.Open("mysql", "root:@/resultview")
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/result")
	if err != nil {
		//fmt.Println("paynay")
		panic(err)

	}

	Con = db
	//defer db.Close()
	return db, nil
}
