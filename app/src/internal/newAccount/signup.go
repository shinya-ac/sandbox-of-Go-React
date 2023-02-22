package signup

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	dbc "github.com/shinya-ac/1Q1A/dbconnection"
)

func Signup(w http.ResponseWriter, r *http.Request) {

	// リクエストヘッダにAccess-Control-Allow-Originを含める
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	// リクエストメソッドがOPTIONSの場合
	if r.Method == http.MethodOptions {
		// Access-Control-Allow-MethodsとAccess-Control-Allow-Headersを含める
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	}

	// リクエストメソッドがOPTIONS以外の場合
	if r.Method != http.MethodOptions {
		// ここにAPIの処理を記述する
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			errMessage := err.Error()
			errResponse := []byte(`{"error": "` + errMessage + `"}`)
			w.Write(errResponse)
			log.Println("サインアップエラー：リクエストのボディの読み込みに失敗しました")
			return
		}
		// ユーザー名、パスワード、その他の必要な情報を取得
		var user User
		err = json.Unmarshal(body, &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			errMessage := err.Error()
			errResponse := []byte(`{"error": "` + errMessage + `"}`)
			w.Write(errResponse)
			log.Println("サインアップエラー：JSONのunmarshalに失敗しました")
			return
		}
		log.Println("unmarshalに成功しました")
		jsonData, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			errMessage := err.Error()
			errResponse := []byte(`{"error": "` + errMessage + `"}`)
			w.Write(errResponse)
			log.Println("サインアップエラー：返却用JSONへのmarshalに失敗しました")
			return
		}
		log.Println("返却用JSONへのmarshalに成功しました")

		// データベースに接続
		db := dbc.ConnectDB()
		defer db.Close()

		//トランザクション開始
		tx, err := db.Begin()
		if err != nil {
			log.Println("トランザクションの開始に失敗しました。エラー：", err)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "サインアップに失敗しました"})
			return
		}

		// ユーザー情報を使って、データベースに新しいアカウントを作成
		if createAccount(tx, user.Username, user.Email, user.Password) {
			// アカウント作成成功の場合、ログイン処理を行う
			//emailをもとにユーザーidを特定
			userId, err := getUserIdByEmail(tx, user.Email)
			if err != nil {
				tx.Rollback()
				w.Header().Set("Content-Type", "application/json")
				errMessage := err.Error()
				errResponse := []byte(`{"error": "` + errMessage + `"}`)
				w.Write(errResponse)
				log.Println("サインアップエラー：メアドからユーザーidの取得に失敗しました")
				return
			}
			log.Println("ユーザーID(emailから特定)%w", userId)
			// TODO: セッション作成失敗時のサインアップのロールバック処理を書く
			if createSession(tx, w, r, userId) {
				err = tx.Commit()
				if err != nil {
					log.Printf("トランザクションのコミットに失敗: %v", err)
					return
				}
				w.Header().Set("Access-Control-Allow-Credentials", "true")
				w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(map[string]string{"message": "サインアップに成功しました", "username": user.Username})
				return
			} else {
				tx.Rollback()
				w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
				json.NewEncoder(w).Encode(map[string]string{"message": "セッションの作成に失敗しました"})
				return
			}
		} else {
			tx.Rollback()
			// アカウント作成失敗の場合、エラー画面にリダイレクト
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "Signin fail", "username": user.Username, "email": user.Email})
			w.Write(jsonData)
			return
		}

	}

}
