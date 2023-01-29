package signup

import (
	"net/http"
	"time"

	"github.com/google/uuid"
)

func createSession(w http.ResponseWriter, r *http.Request, email string) {
	// セッションIDを生成
	sessionID := uuid.New().String()
	// セッションIDをemailやuser_id(usersテーブルのid・主キー)にしてもいいけどuuid（ランダムなid）をセッションに保持させて
	// ユーザーとuuidを対応させたDBを作成し、そこに対応関係を登録しておき、後にクライアントから送信されるsessionのuuidとDB内のuuidが一致するかどうかを
	// 検証する方式の方が安全（つまりセッションハイジャック攻撃対策）

	// セッションIDをCookieに設定（セッション保持期限は24時間）
	http.SetCookie(w, &http.Cookie{
		Name:    "session_id",
		Value:   sessionID,
		Path:    "/",
		Expires: time.Now().Add(time.Hour * 24),
	})
	// セッションIDとemailをデータベースに保存
	saveSessionToDB(sessionID, email)
}
