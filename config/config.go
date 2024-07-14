package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type wxConfig struct {
	AppID          string
	AppSecret      string
	Token          string
	EncodingAESKey string
	AccessToken    string
}

type appConfig struct {
	Port string
}

var (
	WxConfig  *wxConfig
	AppConfig *appConfig
)

func LoadEnv() error {

	WxConfig = &wxConfig{
		Token:          os.Getenv("TOKEN"),
		AppSecret:      os.Getenv("APP_SECRET"),
		AppID:          os.Getenv("APP_ID"),
		EncodingAESKey: os.Getenv("ENCODING_AES_KEY"),
		AccessToken: os.Getenv("ACCESS_TOKEN_TMP"),
	}

	AppConfig = &appConfig{
		Port: os.Getenv("PORT"),
	}

	return nil
}
