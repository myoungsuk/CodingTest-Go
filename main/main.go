package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:myoung1249!@tcp(127.0.0.1:3306)/board")
	defer db.Close()
	if err != nil {
		panic(err)
	}
	var title string
	err = db.QueryRow("SELECT title, created_by, DATE_FORMAT(created_at, '%Y-%m-%d %H:%i:%s') as created_at FROM article WHERE id > 0").Scan(&title)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(title)
	print("success")
}
