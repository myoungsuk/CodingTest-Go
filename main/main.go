package main

import (
	"database/sql"
	"embed"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

type Article struct {
	ID         int64
	CreatedAt  string
	CreatedBy  string
	ModifiedAt string
	ModifiedBy string
	Content    string
	Title      string
	UserID     string
}

var (
	tpl           *template.Template
	gormDB        *gorm.DB
	staticContent embed.FS
)

const (
	MaxPerPage = 5
)

//func init() {
//	tpl = template.Must(template.ParseFS(staticContent, "web/templates/*"))
//}

func (Article) TableName() string {
	return "article"
}

func main() {

	var host = "localhost"
	var database = "board"
	var user = "root"
	var password = "myoung1249!"
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True", user, password, host, database)
	mysqlDB, err := sql.Open("mysql", connectionString)
	defer mysqlDB.Close()

	gormDB, err = gorm.Open(mysql.New(mysql.Config{
		Conn: mysqlDB,
	}), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("success")
	}

	var article Article
	gormDB.First(&article, 1)
	fmt.Println(article)
	//	gormDB.AutoMigrate(&Board{})
	//
	//	http.HandleFunc("/", index)
	//	http.HandleFunc("/write", write)
	//	http.HandleFunc("/board/", board)
	//	http.HandleFunc("/post/", post)
	//	http.HandleFunc("/delete/", delete)
	//	http.HandleFunc("/edit/", edit)
	//	http.Handle("/web/", http.FileServer(http.FS(staticContent)))
	//
	//	fmt.Println("Listening ... !")
	//	http.ListenAndServe(":8080", nil)
	//}
	//
	//func index(w http.ResponseWriter, r *http.Request) {
	//	tpl.ExecuteTemplate(w, "index.gohtml", nil)
	//}
	//
	//func edit(w http.ResponseWriter, r *http.Request) {
	//
	//	id := strings.TrimPrefix(r.URL.Path, "/edit/")
	//	var b Board
	//
	//	gormDB.First(&b, id)
	//
}
