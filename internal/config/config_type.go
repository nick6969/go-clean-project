package config

import "fmt"

type Config struct {
	Server     ServerConfig
	MySQL      DatabaseConfig
	Redis      RedisConfig
	Logger     LoggerConfig
	Token      TokenConfig
	APIDocAuth APIDocAuth
}

type ServerConfig struct {
	Environment string
	Port        string
}

type DatabaseConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func (db *DatabaseConfig) DSN() string {
	return db.User + ":" + db.Password + "@tcp(" + db.Host + ":" + db.Port + ")/" + db.Database + "?parseTime=true"
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
}

func (r *RedisConfig) Address() string {
	return fmt.Sprintf("%s:%s", r.Host, r.Port)
}

type LoggerConfig struct {
	Level  string
	Format string
}

type TokenConfig struct {
	Secret string
}

type APIDocAuth struct {
	UserName string
	Password string
}
