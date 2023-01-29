package signup

import (
	"fmt"

	dbc "github.com/shinya-ac/1Q1A/dbconnection"
)

// usernameではなくemailにすべきかと（？）
func saveSessionToDB(sessionID, email string) {
	// データベースに接続
	db := dbc.ConnectDB()
	defer db.Close()
	//セッションを保存するためのSQLを作成
	sql := "INSERT INTO sessions (sessionID, email) VALUES (?, ?)" // DB内のカラムの名前とこのsaveSessionToDBメソッド内の突っ込みたいデータの名前が一致している必要がある
	// 今回だとsaveSessionToDBメソッドには「sessionID」と「email」という突っ込みたいデータが変数として存在していて、DBにも「sessionID」と「email」というカラムがあるのでうまくいっている
	stmt, err := db.Prepare(sql)
	if err != nil {
		fmt.Println("SQL準備中にエラー", err)
		return
	}
	defer stmt.Close()
	// SQLを実行し、セッションを保存
	_, err = stmt.Exec(sessionID, email)
	if err != nil {
		fmt.Println("SQL実行エラー", err)
		return
	}
}
