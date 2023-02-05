package signup

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func createSession(tx *sql.Tx, w http.ResponseWriter, r *http.Request, user_id int64) bool {
	// セッションIDを生成
	sessionID := uuid.New().String()

	// セッションIDをCookieに設定（セッション保持期限は24時間）
	http.SetCookie(w, &http.Cookie{
		Name:    "session_id",
		Value:   sessionID,
		Path:    "/",
		Expires: time.Now().Add(time.Hour * 24),
	})
	// セッションIDとemailをデータベースに保存
	if saveSessionToDB(tx, sessionID, user_id) {
		log.Println("セッション保存処理に成功")
		return true
	} else {
		log.Println("セッション保存処理でエラー")
		return false
	}
}
