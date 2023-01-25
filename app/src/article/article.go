package article

import (
	"database/sql"
	"fmt"
)

type Article struct {
	id    int
	title string
	body  string
}

func ReadAll(db *sql.DB) {
	var articles []Article
	rows, err := db.Query("select * from article;")
	if err != nil {
		fmt.Println("DBアクセス失敗")
		panic(err)
	}
	for rows.Next() {
		article := Article{}
		err = rows.Scan(&article.id, &article.title, &article.body)
		if err != nil {
			fmt.Println("実行結果のマッピング失敗")
			panic(err)
		}
		articles = append(articles, article)
	}
	rows.Close()

	fmt.Println(articles)
}
