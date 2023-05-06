package question

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	dbc "github.com/shinya-ac/1Q1A/dbconnection"
	s "github.com/shinya-ac/1Q1A/internal/structure"
)

// TODO: ロールバック処理を書く
func createQuestion(questions []s.Question, answers []s.Answer, userId int64, folderId int64) bool {
	// データベースに接続
	// db := dbc.ConnectDB()
	// defer db.Close()

	// // Start a transaction
	// tx, err := db.Begin()
	// if err != nil {
	// 	log.Fatalf("Error トランザクション処理開始時にエラー: %v", err)
	// }
	// defer tx.Rollback()

	// // バルクインサートのためのクエリの作成
	// var values []interface{}
	// var placeholders []string
	// var i int
	// for _, q := range questions {
	// 	i++
	// 	values = append(values, q.Content, userId, folderId)
	// 	placeholders = append(placeholders, "(?, ?, ?)")
	// }
	// query := "INSERT INTO questions (question_content, user_id, folder_id) VALUES " + strings.Join(placeholders, ", ")

	// stmt, err := tx.Prepare(query)
	// if err != nil {
	// 	log.Printf("SQL準備中にエラー: %v", err)
	// 	return false
	// }
	// defer stmt.Close()
	// // SQLを実行し、セッションを保存
	// res, err := stmt.Exec(values...)
	// if err != nil {
	// 	log.Printf("SQL実行中にエラー: %v", err)
	// 	return false
	// }
	db := dbc.ConnectDB()
	defer db.Close()

	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("Error トランザクション処理開始時にエラー: %v", err)
	}
	defer tx.Rollback()

	// クエリの作成
	query := "INSERT INTO questions (question_content, user_id, folder_id) VALUES (?, ?, ?)"

	stmt, err := tx.Prepare(query)
	if err != nil {
		log.Printf("SQL準備中にエラー: %v", err)
		return false
	}
	defer stmt.Close()

	// 解答のバルクインサートを行うクエリを定義
	var answerPlaceholders []string
	var answerValues []interface{}
	var questionIds []int64
	for _, q := range questions {
		// SQLを実行し、セッションを保存
		res, err := stmt.Exec(q.Content, userId, folderId)
		if err != nil {
			log.Printf("SQL実行中にエラー: %v", err)
			return false
		}
		lastInsertId, _ := res.LastInsertId()
		questionIds = append(questionIds, lastInsertId)
		log.Println("質問のインサート一件完了")
	}
	log.Printf("質問一覧: %v\n", questions)
	log.Printf("複数件登録した質問のID一覧: %v", questionIds)

	// questionIds, err := res.LastInsertId()
	// if err != nil {
	// 	log.Printf("直前にインサートした質問データのidの取得に失敗: %v", err)
	// 	return false
	// }

	// for _, _ = range answers {
	// 	answers = append(answers, a.Answer{Content: "これは解答です（バルクインサート）", QuestionId: questionIds})
	// }

	// 続き：以下のforぶんが多分おかしい。最後のインサートのIDをどこで取得して、どこで利用するのかとかをちゃんと考えたほうがいい
	// for _, _ = range questions {
	// 	lastInsertId, _ := res.LastInsertId()
	// 	questionIds = append(questionIds, lastInsertId)
	// 	for _, answer := range answers {
	// 		answerPlaceholders = append(answerPlaceholders, "(?, ?, ?, ?)")
	// 		answerValues = append(answerValues, answer.Content, userId, folderId, lastInsertId)
	// 	}
	// }

	for i, answer := range answers {
		answerPlaceholders = append(answerPlaceholders, "(?, ?, ?, ?)")
		answerValues = append(answerValues, answer.Content, userId, folderId, questionIds[i])
	}
	answerQuery := "INSERT INTO answers (answer_content, user_id, folder_id, question_id) VALUES " + strings.Join(answerPlaceholders, ", ")

	// 回答をバルクインサートする
	stmt, err = tx.Prepare(answerQuery)
	if err != nil {
		log.Printf("解答バルクインサートSQL準備中にエラー: %v", err)
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(answerValues...)
	if err != nil {
		log.Printf("解答バルクインサートSQL実行中にエラー: %v", err)
		return false
	}
	fmt.Println("質問、解答のインサート完了")

	// // Get the auto-generated ID of the inserted question
	// questionID, err := res.LastInsertId()
	// if err != nil {
	// 	log.Printf("直前にインサートした質問データのidの取得に失敗: %v", err)
	// 	return false
	// }

	// answer := "質問の答え"
	// _, err = tx.Exec("INSERT INTO answers (answer_content, question_id, user_id, folder_id) VALUES (?, ?, ?, ?)", answer, questionID, userId, folderId)
	// if err != nil {
	// 	log.Printf("解答データインサート中にエラー: %v", err)
	// 	return false
	// }

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
			a.id AS answer_id, a.answer_content AS answer_content, a.question_id AS question_id
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

func registerQAs(payload s.RequestPayload, userId int64, folderId int64) bool {
	db := dbc.ConnectDB()
	defer db.Close()

	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("Error トランザクション処理開始時にエラー: %v", err)
	}
	defer tx.Rollback()

	for _, qa := range payload.SelectedQAs {
		question := qa.Question
		answer := qa.Answer

		// 質問登録クエリ
		query := "INSERT INTO questions (question_content, user_id, folder_id) VALUES (?, ?, ?)"

		stmt, err := tx.Prepare(query)
		if err != nil {
			log.Printf("SQL準備中にエラー: %v", err)
			return false
		}
		defer stmt.Close()

		// 質問登録SQLを実行し、セッションを保存
		res, err := stmt.Exec(question.Content, userId, folderId)
		if err != nil {
			log.Printf("質問登録クエリ実行中にエラー: %v", err)
			return false
		}
		lastInsertId, _ := res.LastInsertId()
		log.Println("質問のインサート一件完了")

		// 解答登録クエリ
		answerQuery := "INSERT INTO answers (question_id, user_id, answer_content, folder_id) VALUES (?, ?, ?, ?)"
		stmt, err = tx.Prepare(answerQuery)
		if err != nil {
			log.Printf("解答インサートSQL準備中にエラー: %v", err)
			return false
		}

		_, err = stmt.Exec(lastInsertId, userId, answer.Content, folderId)
		if err != nil {
			log.Printf("解答インサートSQL実行中にエラー: %v", err)
			return false
		}
		fmt.Println("質問、解答のインサート完了")
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		log.Printf("トランザクションのコミットに失敗: %v", err)
		return false
	}

	return true

}
