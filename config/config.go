package config

import (
	"log"

	"github.com/daichi-sato-design/todoApp_go_sqlite/utils"
	"gopkg.in/ini.v1"
)

// SettingConfig 設定リストのモデル
type SettingConfig struct{
	Port string
	SQLDriver string
	DbName string
	LogFile string
}

// Config 設定リストのグローバル変数
var Config SettingConfig

func init(){
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

// LoadConfig iniファイルから設定を読み込み、グローバル変数に代入する
func LoadConfig(){
	cfg, err := ini.Load("config.ini")
	if err != nil{
		log.Fatalln(err.Error())
	}

	Config = SettingConfig{
		Port: cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName: cfg.Section("db").Key("name").String(),
		LogFile: cfg.Section("web").Key("logfile").String(),
	}
}