package signup

import (
	"encoding/json"
	"net/http"

	dbc "github.com/shinya-ac/1Q1A/dbconnection"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// ユーザーが入力したメアドとパスワードを取得
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	// ユーザー名を使って、データベースからハッシュ化されたパスワードを取得
	hashedPassword, err := getHashedPassword(email)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "ハッシュされたパスワードの取得に失敗"})
		http.Redirect(w, r, "/home", http.StatusFound)
	}

	// パスワードを比較
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err == nil {
		// パスワードが一致した場合、セッションを作成
		createSession(w, r, email)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "ログイン成功"})
		http.Redirect(w, r, "/home", http.StatusFound)
	} else {
		//w.WriteHeader(http.StatusUnauthorized) // 401 Unauthorized
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		// パスワードが一致しなかった場合、ログイン画面にリダイレクト
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "ログイン失敗"})
		http.Redirect(w, r, "/home", http.StatusFound)
	}
}

func getHashedPassword(email string) ([]byte, error) {
	//データベースに接続
	db := dbc.ConnectDB()
	defer db.Close()

	//データベースからハッシュ化されたパスワードを取得
	var hashedPassword []byte
	err := db.QueryRow("SELECT password FROM users WHERE email = ?", email).Scan(&hashedPassword)
	if err != nil {
		return nil, err
	}
	return hashedPassword, nil
}
