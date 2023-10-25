package main

import (
	"database/sql"
	"embed"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"text/template"
	"time"
)

type Board struct {
	ID        uint `gorm: "primarykey"`
	CreatedAt time.Time
	UpdateAt  time.Time
	Title     string
	Author    string
	Content   string
}

type PassData struct {
	PostData []Board
	Target   string
	Value    string
	PageList []string
	Page     string
}

var (
	tpl           *template.Template
	gormDB        *gorm.DB
	staticContent embed.FS
)

const (
	MaxPerPage = 5
)

func init() {
	tpl = template.Must(template.ParseFS(staticContent, "web/templates/*"))
}

func main() {

	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True", user, password, host, database)
	mysqlDB, err := sql.Open("mysql", connectionString)
	defer mysqlDB.Close()

	gormDB, err = gorm.Open(mysql.New(mysql.Config{
		Conn: mysqlDB,
	}), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	gormDB.AutoMigrate(&Board{})

	http.HandleFunc("/", index)
	http.HandleFunc("/write", write)
	http.HandleFunc("/board/", board)
	http.HandleFunc("/post/", post)
	http.HandleFunc("/delete/", delete)
	http.HandleFunc("/edit/", edit)
	http.Handle("/web/", http.FileServer(http.FS(staticContent)))

	fmt.Println("Listening ... !")
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func edit(w http.ResponseWriter, r *http.Request) {

	id := strings.TrimPrefix(r.URL.Path, "/edit/")
	var b Board

	gormDB.First(&b, id)

}
