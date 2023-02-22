package folder

import (
	"log"

	dbc "github.com/shinya-ac/1Q1A/dbconnection"
)

// TODO: ロールバック処理を書く
func createFolder(folder_title string, userId int64) bool {
	// データベースに接続
	db := dbc.ConnectDB()
	defer db.Close()

	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("Error トランザクション処理開始時にエラー: %v", err)
	}
	defer tx.Rollback()

	sql := "INSERT INTO folders (title, user_id)          VALUES (?, ?)"

	stmt, err := tx.Prepare(sql)
	if err != nil {
		log.Fatalf("SQL準備中にエラー %v", err)
		return false
	}
	defer stmt.Close()
	// SQLを実行し、セッションを保存
	_, err = stmt.Exec(folder_title, userId)
	if err != nil {
		log.Fatalf("SQL実行中にエラー %v", err)
		return false
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		log.Printf("トランザクションのコミットに失敗: %v", err)
		return false
	}

	return true
}
