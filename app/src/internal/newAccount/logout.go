package signup

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	dbc "github.com/shinya-ac/1Q1A/dbconnection"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	// クッキーからセッションIDを取得
	sessionID, err := r.Cookie("session_id")
	if err != nil {
		// セッションが存在しない場合は、ログイン画面にリダイレクト（一旦homeにリダイレクトする）
		// React側の処理もかけたらhomeにリダイレクトではなく失敗をAPIで伝えてReact側でログイン画面に切り替えてもらう
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		errMessage := err.Error()
		errResponse := []byte(`{"error": "` + errMessage + `"}`)
		w.Write(errResponse)
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		log.Printf("クッキーにセッションが存在しません： %s\n", err.Error())
		return
	}

	// DBに接続
	//データベースに接続
	db := dbc.ConnectDB()
	defer db.Close()

	// セッションIDを持つ（DB内の）セッションを削除
	_, err = db.Exec("DELETE FROM sessions WHERE sessionID = ?", sessionID.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		errMessage := err.Error()
		errResponse := []byte(`{"error": "` + errMessage + `"}`)
		w.Write(errResponse)
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		log.Printf("ログアウト失敗・DBのセッションを削除できませんでした： %s\n", err.Error())
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	// クッキーからセッションIDを削除（セッションの期限を強制的に現在時刻にする＝セッションを消すという挙動になる）
	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   "",
		Expires: time.Now(),
	}
	http.SetCookie(w, cookie)

	// ログアウト成功後はログイン画面にリダイレクト
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "ログアウト成功"})
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
