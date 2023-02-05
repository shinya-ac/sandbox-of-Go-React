package signup

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	dbc "github.com/shinya-ac/1Q1A/dbconnection"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		errMessage := err.Error()
		errResponse := []byte(`{"error": "` + errMessage + `"}`)
		w.Write(errResponse)
		return
	}
	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		errMessage := err.Error()
		errResponse := []byte(`{"error": "` + errMessage + `"}`)
		w.Write(errResponse)
		log.Printf("JSONのUnmarshalに失敗： %s\n", err.Error())
		return
	}
	log.Println("Unmarshal成功")

	// メアドを使って、データベースからハッシュ化されたパスワードを取得
	hashedPassword, err := getHashedPassword(user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		errMessage := err.Error()
		errResponse := []byte(`{"error": "` + errMessage + `"}`)
		w.Write(errResponse)
		log.Println("ハッシュ化されたパスワードの取得に失敗")
		return
	}

	// データベースに接続
	db := dbc.ConnectDB()
	defer db.Close()

	//トランザクション開始
	tx, err := db.Begin()

	//メアドを使ってそのユーザーのIDを取得
	userId, err := getUserIdByEmail(tx, user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		errMessage := err.Error()
		errResponse := []byte(`{"error": "` + errMessage + `"}`)
		w.Write(errResponse)
		log.Println("メアドからユーザーidの取得に失敗")
		return
	}

	// パスワードを比較
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(user.Password))
	if err == nil {
		// パスワードが一致した場合、セッションを作成
		if createSession(tx, w, r, userId) {
			err = tx.Commit()
			if err != nil {
				log.Printf("トランザクションのコミットに失敗: %v", err)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "ログイン成功"})
			http.Redirect(w, r, "/home", http.StatusFound)
			return
		} else {
			tx.Rollback()
			http.Error(w, err.Error(), http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "ログイン失敗：セッションの作成に失敗しました"})
			log.Printf("クッキーにセッションが存在しません： %s\n", err.Error())
			return
		}

	} else {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		w.Header().Set("Content-Type", "application/json")
		errMessage := err.Error()
		errResponse := []byte(`{"error": "` + errMessage + `"}`)
		w.Write(errResponse)
		log.Println("メアド認証失敗")
		return
	}
}
