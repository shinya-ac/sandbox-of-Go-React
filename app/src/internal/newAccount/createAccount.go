package signup

import (
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func createAccount(tx *sql.Tx, username, email, password string) bool {
	// パスワードをハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("パスワードのハッシュ化に失敗")
		return false
	}

	// アカウントを作成するためのSQLを作成
	sql := "INSERT INTO users (username, password, email) VALUES (?, ?, ?)"
	stmt, err := tx.Prepare(sql)
	if err != nil {
		log.Println("SQLクエリ準備中にエラー")
		return false
	}
	defer stmt.Close()
	// SQLを実行し、アカウントを作成
	_, err = stmt.Exec(username, hashedPassword, email)
	if err != nil {
		log.Println("SQL実行時にエラー：アカウントを作成できませんでした")
		return false
	}
	return true
}
