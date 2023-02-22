package question

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"

	a "github.com/shinya-ac/1Q1A/internal/answer"
)

func QuestionHandler(w http.ResponseWriter, r *http.Request) {
	//TODO: リファクタリングをする

	// リクエストヘッダにAccess-Control-Allow-Originを含める
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	// リクエストメソッドがOPTIONSの場合
	if r.Method == http.MethodOptions {
		// Access-Control-Allow-MethodsとAccess-Control-Allow-Headersを含める
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	}

	// リクエストメソッドがOPTIONS以外の場合
	if r.Method != http.MethodOptions {
		// ここにAPIの処理を記述する
		r.ParseForm()
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Println(r.Form)
		fmt.Println(string(body))

		//REST API的にURLから対応するフォルダーを特定し質問を作成する場合はこちら
		// pathPattern := regexp.MustCompile(`/folders/(\d+)/question`)
		// matches := pathPattern.FindStringSubmatch(r.URL.Path)
		// if matches == nil {
		// 	http.NotFound(w, r)
		// 	json.NewEncoder(w).Encode(map[string]string{"message": "No Matches folder id in URL"})
		// 	return
		// }
		// folderId, err := strconv.ParseInt(matches[1], 10, 64)
		// if err != nil {
		// 	http.Error(w, "Invalid folder ID", http.StatusBadRequest)
		// 	return
		// }

		var question Question
		err := json.Unmarshal(body, &question)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			errMessage := err.Error()
			errResponse := []byte(`{"error": "` + errMessage + `"}`)
			w.Write(errResponse)
			json.NewEncoder(w).Encode(map[string]string{"message": "Unmarshal Error"})
			fmt.Printf("Error decoding JSON: %s\n", err.Error())
			return
		}
		fmt.Println("Unmarshal完了")
		fmt.Println(question)
		jsonData, err := json.Marshal(question)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "Marshal Error"})
			return
		}
		fmt.Println("Marshal完了")
		// ユーザー情報を使って、データベースに新しいアカウントを作成
		userId := r.Context().Value("userId").(int64)
		if createQuestion(question.Content, userId, question.FolderId) {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{"message": "create question succcess", "userId": userId, "question": question.Content})
			//http.Redirect(w, r, "/home", http.StatusFound)
		} else {
			// 質問作成失敗の場合、エラー画面にリダイレクト
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "Create question fail", "question": question.Content})
			w.Write(jsonData)
			//http.Redirect(w, r, "/home", http.StatusFound)
		}
	}
}

// フォルダーに属する質問とそれに紐づく解答一覧を閲覧するエンドポイント
func FolderReadHandler(w http.ResponseWriter, r *http.Request) {
	// リクエストヘッダにAccess-Control-Allow-Originを含める
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	// リクエストメソッドがOPTIONSの場合
	if r.Method == http.MethodOptions {
		// Access-Control-Allow-MethodsとAccess-Control-Allow-Headersを含める
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	}

	// リクエストメソッドがOPTIONS以外の場合
	if r.Method != http.MethodOptions {
		// ここにAPIの処理を記述する

		pathPattern := regexp.MustCompile(`/folders/(\d+)`)
		matches := pathPattern.FindStringSubmatch(r.URL.Path)
		if matches == nil {
			http.NotFound(w, r)
			json.NewEncoder(w).Encode(map[string]string{"message": "No Matches folder id in URL"})
			return
		}
		folderId, err := strconv.ParseInt(matches[1], 10, 64)
		if err != nil {
			http.Error(w, "Invalid folder ID", http.StatusBadRequest)
			return
		}

		// ユーザー情報を使って、データベースに新しいアカウントを作成
		userId := r.Context().Value("userId").(int64)
		rows, err := ReadQuestions(userId, folderId)
		if err != nil {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			w.Header().Set("Content-Type", "application/json")
			errMessage := err.Error()
			errResponse := []byte(`{"error": "` + errMessage + `"}`)
			w.Write(errResponse)
			log.Println("QA取得エラー：QAの取得に失敗しました")
			return
		} else {
			// 返却するquestion一覧のための配列領域
			questions := []Question{}
			answers := []a.Answer{}
			for rows.Next() {
				var question Question
				var answer a.Answer

				type Question struct {
					Content  string `json:"Content"`
					FolderId int64  `json:"FolderId,omitempty"`
				}

				err := rows.Scan(&question.Id, &question.Content, &answer.Id, &answer.Content)
				if err != nil {
					panic(err.Error())
				}

				questions = append(questions, question)
				answers = append(answers, answer)
			}

			if err := rows.Err(); err != nil {
				w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
				w.Header().Set("Content-Type", "application/json")
				errMessage := err.Error()
				errResponse := []byte(`{"error": "` + errMessage + `"}`)
				w.Write(errResponse)
				log.Println("rowsエラー：QAの取得に失敗しました")
			}

			// スライスをJSONに変換
			jsonBytes, err := json.Marshal(questions)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("ログインしているユーザーのidは：%dです。以下のようなQuestionを取得できました： %+v", userId, questions)
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(jsonBytes)
		}
	}
}
