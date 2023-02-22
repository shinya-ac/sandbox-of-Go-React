package question

import (
	"database/sql"
	"log"

	dbc "github.com/shinya-ac/1Q1A/dbconnection"
)

// TODO: ロールバック処理を書く
func createQuestion(question_content string, userId int64, folderId int64) bool {
	// データベースに接続
	db := dbc.ConnectDB()
	defer db.Close()

	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("Error トランザクション処理開始時にエラー: %v", err)
	}
	defer tx.Rollback()

	sql := "INSERT INTO questions (question_content, user_id, folder_id)          VALUES (?, ?, ?)"

	stmt, err := tx.Prepare(sql)
	if err != nil {
		log.Println("SQL準備中にエラー")
		return false
	}
	defer stmt.Close()
	// SQLを実行し、セッションを保存
	res, err := stmt.Exec(question_content, userId, folderId)
	if err != nil {
		log.Println("SQL実行中にエラー")
		return false
	}

	// Get the auto-generated ID of the inserted question
	questionID, err := res.LastInsertId()
	if err != nil {
		log.Printf("直前にインサートした質問データのidの取得に失敗: %v", err)
		return false
	}

	answer := "質問の答え"
	_, err = tx.Exec("INSERT INTO answers (answer_content, question_id, user_id, folder_id) VALUES (?, ?, ?, ?)", answer, questionID, userId, folderId)
	if err != nil {
		log.Printf("解答データインサート中にエラー: %v", err)
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

func ReadQuestions(userId int64, folderId int64) (*sql.Rows, error) {
	// データベースに接続
	db := dbc.ConnectDB()
	defer db.Close()

	query := `
		SELECT q.id AS question_id, q.question_content AS question_content,
			a.id AS answer_id, a.answer_content AS answer_content
		FROM questions q
		LEFT JOIN answers a ON q.id = a.question_id
		WHERE q.folder_id = ? AND q.user_id = ? AND a.user_id = ?
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Printf("SQL準備中にエラー: %v", err)
		return nil, err
	}
	defer stmt.Close()

	// SQLを実行し、セッションを保存
	rows, err := stmt.Query(folderId, userId, userId)
	if err != nil {
		log.Println("SQL実行中にエラー")
		return nil, err
	}

	return rows, nil
}
