package signup

import (
	"database/sql"
	"log"

	dbc "github.com/shinya-ac/1Q1A/dbconnection"
)

func getHashedPassword(email string) ([]byte, error) {
	//データベースに接続
	db := dbc.ConnectDB()
	defer db.Close()

	//データベースからハッシュ化されたパスワードを取得
	var hashedPassword []byte
	err := db.QueryRow("SELECT password FROM users WHERE email = ?", email).Scan(&hashedPassword)
	if err != nil {
		log.Println("DB上のハッシュ化されたパスワードの取得に失敗")
		return nil, err
	}
	log.Println("DB上のハッシュ化されたパスワードの取得に成功")
	return hashedPassword, nil
}

func getUserIdByEmail(tx *sql.Tx, email string) (int64, error) {
	//データベースからユーザーidを取得（取得失敗の場合は0とエラーが返る）
	var userId int64
	err := tx.QueryRow("SELECT id FROM users WHERE email = ?", email).Scan(&userId)
	if err != nil {
		log.Println("メアドからユーザーIDの取得に失敗")
		return 0, err
	}
	log.Println("メアドからユーザーIDの取得に成功：%w", userId)
	return userId, nil
}
