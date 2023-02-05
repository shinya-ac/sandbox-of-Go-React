package signup

import (
	"database/sql"
	"log"
)

// usernameではなくemailにすべきかと（？）
func saveSessionToDB(tx *sql.Tx, sessionID string, user_id int64) bool {
	if user_id == 0 {
		log.Println("セッション作成エラー：user_idが不正です")
		return false
	}
	// セッションを保存するためのSQLを作成
	sql := "INSERT INTO sessions (sessionID, user_id) VALUES (?, ?)"
	stmt, err := tx.Prepare(sql)
	if err != nil {
		log.Println("SQL準備中にエラー")
		return false
	}
	defer stmt.Close()
	// SQLを実行し、セッションを保存
	_, err = stmt.Exec(sessionID, user_id)
	if err != nil {
		log.Println("SQL実行エラー")
		return false
	}
	return true
}
