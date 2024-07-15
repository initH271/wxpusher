package config

import (
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/viper"
)

type wxConfig struct {
	AppID          string `json:"AppID"`
	AppSecret      string `json:"AppSecret"`
	Token          string `json:"Token"`
	EncodingAESKey string `json:"EncodingAESKey"`
	AccessToken    string `json:"AccessToken"`
}

type server struct {
	Port string
}

type appConfig struct {
	WxConfig wxConfig
	Server   server
}

var (
	AppConfig appConfig
)

func LoadEnv() error {
	viper.AddConfigPath(".")
	viper.SetConfigName("wxconfig")
	viper.SetConfigType("toml")
	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("read config failed: %v", err)
	}

	err := viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Printf("read config failed: %v", err)
	}

	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Printf("配置文件发生改变: %s\n", in.Name)
		if in.Op == fsnotify.Write && strings.Contains(in.Name, "wxconfig") {
			viper.Unmarshal(&AppConfig)
		}
	})
	log.Printf("读取配置 Config: %+v", AppConfig)

	viper.WatchConfig()
	return nil
}
