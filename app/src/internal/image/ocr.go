package image

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/shinya-ac/1Q1A/config"

	vision "cloud.google.com/go/vision/apiv1"
)

func DetectString(filename string) string {
	//TODO: リファクタリングをする
	ctx := context.Background()

	fmt.Printf("api key：%s\n", config.Config.GOOGLE_APPLICATION_CREDENTIALS)
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", config.Config.GOOGLE_APPLICATION_CREDENTIALS)
	files, _ := os.ReadDir(".")
	fmt.Printf("ディレクトリ直下のファイル一覧")
	for _, file := range files {
		fmt.Println(file.Name())
	}
	fmt.Printf("api key（環境変数）：%s\n", os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))

	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	defer file.Close()
	image, err := vision.NewImageFromReader(file)
	if err != nil {
		log.Fatalf("Failed to create image: %v", err)
	}

	texts, err := client.DetectTexts(ctx, image, nil, 10)
	if err != nil {
		log.Fatalf("Failed to detect labels: %v", err)
	}

	var textArr []string
	// for _, text := range texts {
	// 	textArr = append(textArr, text.Description)
	// 	//fmt.Println(text.Description)
	// }
	if len(texts) > 0 {
		textArr = []string{texts[0].Description}
	}
	fmt.Printf("TextArrの値：%v\n", textArr)
	result := strings.Join(textArr, " ")
	fmt.Printf("resultの値：%v\n", result)
	return result
}
