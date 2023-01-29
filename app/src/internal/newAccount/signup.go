package signup

import "net/http"

func Signup(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// ユーザー名、パスワード、その他の必要な情報を取得
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	email := r.Form.Get("email")
	// ユーザー情報を使って、データベースに新しいアカウントを作成
	if createAccount(username, password, email) {
		// アカウント作成成功の場合、ログイン処理を行う
		createSession(w, r, email)
		http.Redirect(w, r, "/home", http.StatusFound)
	} else {
		// アカウント作成失敗の場合、エラー画面にリダイレクト
		http.Redirect(w, r, "/error", http.StatusFound)
	}
}
