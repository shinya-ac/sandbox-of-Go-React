package signup

import (
	"golang.org/x/crypto/bcrypt"

	dbc "github.com/shinya-ac/1Q1A/dbconnection"
)

func createAccount(username, email, password string) bool {
	// パスワードをハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return false
	}
	// データベースに接続
	db := dbc.ConnectDB()
	defer db.Close()
	// アカウントを作成するためのSQLを作成
	sql := "INSERT INTO users (username, password, email) VALUES (?, ?, ?)" //まだテーブルが作られていないので動かないと思われる
	stmt, err := db.Prepare(sql)
	if err != nil {
		return false
	}
	defer stmt.Close()
	// SQLを実行し、アカウントを作成
	_, err = stmt.Exec(username, hashedPassword, email)
	if err != nil {
		return false
	}
	return true
}
