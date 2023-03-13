package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/ini.v1"
	_ "gopkg.in/ini.v1"
)

//go get gopkg.in/ini.v1←コンフィグファイル

type ConfigList struct {
	GOOGLE_APPLICATION_CREDENTIALS string
	ApiSecret                      string
	OPEN_AI_API                    string
}

var Config ConfigList

func init() {
	fmt.Println("iniパッケージによるconfigの初期化開始")
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		log.Printf("fail to road file%v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		ApiSecret:                      cfg.Section("Slack").Key("fuga").String(),
		GOOGLE_APPLICATION_CREDENTIALS: cfg.Section("vision").Key("apiKey").String(),
		OPEN_AI_API:                    cfg.Section("openAI").Key("apiKey").String(),
	}

}
