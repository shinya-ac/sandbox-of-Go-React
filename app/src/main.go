package main

import (
	"context"
	"database/sql"
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/shinya-ac/1Q1A/article"
	dbc "github.com/shinya-ac/1Q1A/dbconnection"
	"github.com/shinya-ac/1Q1A/domain"
	"github.com/shinya-ac/1Q1A/handler"
	_ "github.com/shinya-ac/1Q1A/handler"
	na "github.com/shinya-ac/1Q1A/internal/newAccount"

	_ "github.com/go-sql-driver/mysql"
)

var indexTmpl embed.FS

func index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(indexTmpl, "index.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, "")
	if err != nil {
		log.Fatal(err)
	}
}

//----------以下1Q1Aアプリの実装----------

func home(w http.ResponseWriter, r *http.Request) {
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
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "HOMEまで辿り着いています"})
		fmt.Println("空ページ")
		fmt.Fprintf(w, "<h1>Hello</h1>")
	}
}

// 認可機能の本体(ミドルウェア)↓
var sessions = map[string]string{} // セッションIDをキーにして、ログインしているユーザーのemailを保存している
func auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionID, err := r.Cookie("session_id")
		if err != nil {
			// セッションが存在しない場合は、ログイン画面にリダイレクト
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "認可失敗・セッションが存在しません"})
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return
		}
		var email string
		// データベースに接続
		db := dbc.ConnectDB()
		defer db.Close()
		//クライアントから受け取ったsessionID情報と一致する値があるかどうかをsessionsテーブルの中で探している。一致したデータがあればそのセッションと対応関係にあるメアドを取得するというコード
		err = db.QueryRow("SELECT email FROM sessions WHERE sessionID = ?", sessionID.Value).Scan(&email)
		if err == sql.ErrNoRows {
			// DBに認可リクエストに対して整合するデータが存在しない場合は、ログイン画面にリダイレクト
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "認可失敗・セッションリクエストと一致するセッションデータがありません"})
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return
		} else if err != nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "認可失敗・セッションSQLを実行した後にエラーが出ました"})
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		ctx := context.WithValue(r.Context(), "email", email)
		next.ServeHTTP(w, r.WithContext(ctx))

	}
}

// 認可機能確認のためのエンドポイント
func isAuthorized(w http.ResponseWriter, r *http.Request) {
	// requestからcontextを取得（認可機能で認可されたユーザーのemailを取得）
	email := r.Context().Value("email").(string)
	//w.Write([]byte("Hello, " + email))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "認可成功"})
	fmt.Println("認可実験・これが表示されていれば認可機能が機能している")
	fmt.Fprintf(w, "<h1>Hello %s さん</h1>", email)
}

func main() {
	fmt.Println("やぁ、処理を開始するよ")
	db := dbc.ConnectDB()
	defer db.Close()
	article.ReadAll(db)
	fmt.Println("処理終了")
	fmt.Println("webサーバー起動開始")
	//↓一つ目のハンドラ。これは「/」のパスに割り当てた静的ファイルを配信する部分
	// files := http.FileServer(http.Dir("public"))
	// http.Handle("/", files)
	hub := domain.NewHub()
	go hub.RunLoop()
	// 「http://localhost:8080/ws」と言うリクエストが来た際はhttp通信をsocketにアップグレードする
	http.HandleFunc("/ws", handler.NewWebsocketHandler(hub).Handle)

	http.HandleFunc("/home", home)
	http.HandleFunc("/signup", na.Signup)
	http.HandleFunc("/login", na.Login)
	http.HandleFunc("/logout", na.Logout)
	//以下のisAuthorizedというエンドポイントは「ログイン済みのユーザー」のみがアクセスできるようにしたいので
	//authというログインを確認するミドルウェアを噛ませている
	http.HandleFunc("/isAuthorized", auth(isAuthorized))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Panicln("Serve Error:", err)
	}
	fmt.Println("webサーバー起動終了")

}

//ログイン・認証関連の未実装機能
//現状、セッションの保存が「セッションID↔️メアド」という対応になっている
//メアドも一意な数字だからこれでもいいけど「セッションID↔️ユーザーID」の方が何かとその後の実装が便利だからそうする

//永続セッション機能（remember me機能・ログアウトしない限りブラウザを閉じてもログインが保持される機能）の実装もする

//ユーザーの登録機能は作ったけど変更・削除機能・ユーザー情報の閲覧機能もまだ
//メアドによるアカウント有効化機能（受信ボックスのリンクをクリックしてもらう機能）もまだ

//上記によるパスワードの再設定機能もまだ
