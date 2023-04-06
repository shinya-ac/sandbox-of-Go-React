package image

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func ConvertImageHandler(w http.ResponseWriter, r *http.Request) {
	//TODO: リファクタリングをする

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
		// multipart/form-dataをパースする
		err := r.ParseMultipartForm(32 << 20)
		if err != nil {
			log.Println(err)
			http.Error(w, "multipart formをパースできませんでした", http.StatusBadRequest)
			return
		}

		// バリデーション
		const MaxUploadSize = 1024 * 1024
		r.Body = http.MaxBytesReader(w, r.Body, MaxUploadSize)
		if err := r.ParseMultipartForm(MaxUploadSize); err != nil {
			http.Error(w, "アップロードされたファイルが大きすぎます。1MB以下のファイルを選択してください", http.StatusBadRequest)
		}

		// リクエストボディからファイルを取得
		file, _, err := r.FormFile("image")
		if err != nil {
			log.Printf("エラー：%v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			errMessage := err.Error()
			errResponse := []byte(`{"error": "` + errMessage + `"}`)
			w.Write(errResponse)
			json.NewEncoder(w).Encode(map[string]string{"message": "ファイルが存在しません"})
			fmt.Printf("クライアントからのファイル取得エラー: %s\n", err.Error())
			return
		}
		defer file.Close()

		// ファイルをバイト配列に変換
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			log.Println(err)
			http.Error(w, "ファイルの読み込みに失敗しました", http.StatusInternalServerError)
			return
		}

		// ローカルにファイルを保存する
		err = ioutil.WriteFile("image.jpg", fileBytes, os.ModePerm)
		if err != nil {
			http.Error(w, "Failed to save image file", http.StatusInternalServerError)
			return
		}

		// base64形式にエンコード
		//b64 := base64.StdEncoding.EncodeToString(fileBytes)

		// 画像を文字に変換
		// detectedStr := DetectString("image.jpg")　こっちが本番用
		// detectedStr := DetectString("img.png")　こっちはテスト用（どんな写真入れても決めうちで文字付きの写真が送信される）
		detectedStr := DetectString("image.jpg")

		if len(detectedStr) == 0 {
			log.Println("文字は空")
			http.Error(w, "文字は何も読み取れませんでした", http.StatusInternalServerError)
			return
		}

		// レスポンスを返却　以下は出力しない
		// w.Write([]byte(b64))
		// スライスをJSONに変換
		// jsonBytes, err := json.Marshal(detectedStr)
		// if err != nil {
		// 	log.Fatal(err)
		// 	return
		// }
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		//w.Write(jsonBytes)
		//fmt.Printf("detectedStr：%s", []byte(detectedStr))
		//fmt.Printf("[]byte(detectedStr)：%s", []byte(detectedStr))
		w.Write([]byte(detectedStr))
		// openAIに投げる
		// res := CallOpenAI(detectedStr)

		// w.Write([]byte(b64))
		// // スライスをJSONに変換
		// resJsonByte, err := json.Marshal(res)
		// if err != nil {
		// 	log.Fatal(err)
		// 	return
		// }

		// w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		// w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(http.StatusOK)
		// w.Write(resJsonByte)
	}

}
