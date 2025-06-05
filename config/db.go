package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	dsn := "root:Liftbringer67@#@tcp(127.0.0.1:3306)/mydb?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	DB = db
	fmt.Println("Connected to database")
}
