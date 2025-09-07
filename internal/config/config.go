package config

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/spf13/viper"
)

func LoadConfig() (*Config, error) {
	v := viper.New()

	// 啟用自動讀取環境變數, 用於 Server 部署
	v.AutomaticEnv()

	// 設定讀取 .env 檔案，用於本地開發使用
	v.SetConfigName(".env")
	v.SetConfigType("env")
	v.AddConfigPath(".")

	err := v.ReadInConfig()
	if err != nil {
		// 如果錯誤是指定的 config 檔案不存在，則不處理
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	}

	var cfg Config

	cfg.Server.Environment = v.GetString("SERVER_ENVIRONMENT")
	cfg.Server.Port = v.GetString("SERVER_PORT")

	cfg.MySQL.User = v.GetString("MYSQL_USER")
	cfg.MySQL.Password = v.GetString("MYSQL_PASSWORD")
	cfg.MySQL.Host = v.GetString("MYSQL_HOST")
	cfg.MySQL.Port = v.GetString("MYSQL_PORT")
	cfg.MySQL.Database = v.GetString("MYSQL_DATABASE")

	cfg.Redis.Host = v.GetString("REDIS_HOST")
	cfg.Redis.Port = v.GetString("REDIS_PORT")
	cfg.Redis.Password = v.GetString("REDIS_PASSWORD")

	cfg.Logger.Level = v.GetString("LOGGER_LEVEL")
	cfg.Logger.Format = v.GetString("LOGGER_FORMAT")

	tokenSecret := v.GetString("TOKEN_SECRET")
	if tokenSecret == "" {
		return nil, fmt.Errorf("TOKEN_SECRET is required")
	}
	cfg.Token.Secret = replaceEscape(tokenSecret)

	cfg.APIDocAuth.UserName = v.GetString("API_DOC_AUTH_USERNAME")
	cfg.APIDocAuth.Password = v.GetString("API_DOC_AUTH_PASSWORD")

	return &cfg, nil
}

func replaceEscape(value string) string {
	escapeRegex := regexp.MustCompile(`\\.`)
	newValue := escapeRegex.ReplaceAllStringFunc(value, func(match string) string {
		c := strings.TrimPrefix(match, `\`)
		switch c {
		case "n":
			return "\n"
		case "r":
			return "\r"
		default:
			return match
		}
	})
	return newValue
}
