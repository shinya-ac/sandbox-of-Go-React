package signup

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Signup(w http.ResponseWriter, r *http.Request) {

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
		// ユーザー名、パスワード、その他の必要な情報を取得
		var user User
		err := json.Unmarshal(body, &user)
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
		jsonData, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "Marshal Error"})
			return
		}
		fmt.Println("Marshal完了")
		// ユーザー情報を使って、データベースに新しいアカウントを作成
		if createAccount(user.Username, user.Email, user.Password) {
			// アカウント作成成功の場合、ログイン処理を行う
			createSession(w, r, user.Email)
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "Signin succcess", "username": user.Username})
			//http.Redirect(w, r, "/home", http.StatusFound)
		} else {
			// アカウント作成失敗の場合、エラー画面にリダイレクト
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "Signin fail", "username": user.Username, "email": user.Email})
			w.Write(jsonData)
			//http.Redirect(w, r, "/home", http.StatusFound)
		}

	}

}
