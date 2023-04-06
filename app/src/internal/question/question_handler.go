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
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, FolderId")
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

		//openAIからbody内の本文を元に質問と解答を作成
		questionContents, answerContents := CallOpenAI(string(body))
		//questionBytes := []byte(questionJSON)
		fmt.Printf("OpenAIからのレス：%v\n", questionContents)

		// スライスをQuestion構造体のスライスに変換する
		var questions []Question
		for i, qContent := range questionContents {
			q := Question{Id: int64(i + 1), Content: qContent}
			questions = append(questions, q)
		}

		var answers []a.Answer
		for i, aContent := range answerContents {
			a := a.Answer{Id: int64(i + 1), Content: aContent}
			answers = append(answers, a)
		}

		//var question Question
		// err := json.Unmarshal([]byte(question_content), &question)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusBadRequest)
		// 	w.Header().Set("Content-Type", "application/json")
		// 	errMessage := err.Error()
		// 	errResponse := []byte(`{"error": "` + errMessage + `"}`)
		// 	w.Write(errResponse)
		// 	json.NewEncoder(w).Encode(map[string]string{"message": "Unmarshal Error"})
		// 	fmt.Printf("Error decoding JSON: %s\n", err.Error())
		// 	return
		// }
		fmt.Println("質問Unmarshal完了")
		jsonData, err := json.Marshal(questions)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "Marshal Error"})
			return
		}
		//fmt.Printf("jsonData：%s", jsonData)
		fmt.Println("質問Marshal完了")

		answerJsonData, err := json.Marshal(answers)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "Marshal Error"})
			return
		}
		fmt.Println("解答Marshal完了")

		// ユーザー情報を使って、データベースに新しいアカウントを作成
		userId := r.Context().Value("userId").(int64)
		// ヘッダーからFolderIdを取得する
		folderIdStr := r.Header.Get("FolderId")
		folderId, err := strconv.ParseInt(folderIdStr, 10, 64)
		if err != nil {
			panic(err)
		}
		fmt.Printf("folderIdの値：%d", folderId)
		if createQuestion(questions, answers, userId, folderId) {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{"message": "create question succcess", "userId": userId, "question": questions, "answers": answers})
			//http.Redirect(w, r, "/home", http.StatusFound)
		} else {
			// 質問作成失敗の場合、エラー画面にリダイレクト
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{"message": "Create question fail", "question": questions})
			w.Write(jsonData)
			w.Write(answerJsonData)
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
			type ResponseQuestion struct {
				Id      int64  `json:"Id"`
				Content string `json:"Content"`
			}
			type ResponseAnswer struct {
				Id         int64  `json:"Id"`
				Content    string `json:"Content"`
				QuestionId int64  `json:"QuestionId,"`
			}
			// 返却するquestion一覧のための配列領域
			questions := []ResponseQuestion{}
			answers := []ResponseAnswer{}
			for rows.Next() {
				var question ResponseQuestion
				var answer ResponseAnswer

				err := rows.Scan(&question.Id, &question.Content, &answer.Id, &answer.Content, &answer.QuestionId)
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

			// // 質問一覧スライスをJSONに変換
			// questionJsonBytes, err := json.Marshal(questions)
			// if err != nil {
			// 	log.Fatal(err)
			// }

			// // 解答一覧スライスをJSONに変換
			// answerJsonBytes, err := json.Marshal(answers)
			// if err != nil {
			// 	log.Fatal(err)
			// }

			// 質問と回答をまとめたレスポンスを作成←ここにJsonに変換するのではなく一括してJsonに変換する
			type Response struct {
				Questions []ResponseQuestion `json:"questions"`
				Answers   []ResponseAnswer   `json:"answers"`
			}
			responseData := Response{
				Questions: questions,
				Answers:   answers,
			}
			// レスポンスをJSONに変換して返す
			responseBytes, err := json.Marshal(responseData)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			fmt.Printf("ログインしているユーザーのidは：%dです。以下のようなQuestionを取得できました： %+v", userId, questions)
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Printf("\n返却質問確認用：%v\n", questions)
			fmt.Printf("返却解答確認用：%v\n", answers)
			// w.Write(questionJsonBytes)
			// w.Write(answerJsonBytes)
			w.Write(responseBytes)
		}
	}
}
