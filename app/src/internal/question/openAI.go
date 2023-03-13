package question

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/shinya-ac/1Q1A/config"
)

var apiKey = config.Config.OPEN_AI_API
var model = "text-davinci-003"

func CallOpenAI(texts string) ([]string, []string) {
	//joinedTexts := strings.Join(texts, " ")
	prompt := texts + "上記の文章から一問一答を５つ作成してください。なお、質問は「Q.」から書き始めてください。解答は「A.」から始めてください。"
	temperature := 0.0
	maxTokens := 2060
	requestBody, err := json.Marshal(map[string]interface{}{
		"model":       model,
		"prompt":      prompt,
		"temperature": temperature,
		"max_tokens":  maxTokens,
	})
	if err != nil {
		panic(err)
	}
	request, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", bytes.NewBuffer(requestBody))
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+apiKey)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	var data map[string]interface{}
	err = json.Unmarshal(responseBody, &data)
	if err != nil {
		panic(err)
	}
	text := data["choices"].([]interface{})[0].(map[string]interface{})["text"].(string)
	fmt.Printf("作成されたQAです：%s", text)

	// AIから返されるQAのデータが「Q.」から質問部分が始まって、改行して「A.」から解答部分が始まって改行されているという前提
	re := regexp.MustCompile(`Q\.(.*?)\nA\.(.*?)\n`)
	matches := re.FindAllStringSubmatch(text, -1)

	// 文章から正規表現パターンにマッチする部分を抽出する
	var questionContent []string
	var answerContent []string

	for _, match := range matches {
		questionContent = append(questionContent, strings.TrimSpace(match[1]))
		answerContent = append(answerContent, strings.TrimSpace(match[2]))
	}

	fmt.Println("Questions:")
	for _, question := range questionContent {
		fmt.Println(question)
	}

	fmt.Println("\nAnswers:")
	for _, answer := range answerContent {
		fmt.Println(answer)
	}

	// 文章から正規表現パターンにマッチする部分を抽出する

	return questionContent, answerContent
}
