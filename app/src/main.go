package main

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/shinya-ac/GoChat/article"
	"github.com/shinya-ac/GoChat/domain"
	"github.com/shinya-ac/GoChat/handler"
	_ "github.com/shinya-ac/GoChat/handler"

	_ "github.com/go-sql-driver/mysql"
)

func open(path string, count uint) *sql.DB {
	db, err := sql.Open("mysql", path)
	if err != nil {
		log.Fatal("open error:", err)
	}

	if err = db.Ping(); err != nil {
		time.Sleep(time.Second * 2)
		count--
		fmt.Printf("retry... count:%v\n", count)
		return open(path, count)
	}

	fmt.Println("db connected!!")
	return db
}

func connectDB() *sql.DB {
	var path string = fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8&parseTime=true",
		os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DATABASE"))

	return open(path, 100)
}

var indexTmpl embed.FS

func index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(indexTmpl, "index.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, "")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("やぁ、処理を開始するよ")
	db := connectDB()
	defer db.Close()
	article.ReadAll(db)
	fmt.Println("処理終了")
	fmt.Println("webサーバー起動開始")
	//↓一つ目のハンドラ。これは「/」のパスに割り当てた静的ファイルを配信する部分
	// files := http.FileServer(http.Dir("public"))
	// http.Handle("/", files)
	hub := domain.NewHub()
	go hub.RunLoop()
	// 「http://localhost:8080/ws」と言うリクエストが来た際はhttp通信をsocketにアップグレードする
	http.HandleFunc("/ws", handler.NewWebsocketHandler(hub).Handle)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Panicln("Serve Error:", err)
	}
	fmt.Println("webサーバー起動終了")

}
