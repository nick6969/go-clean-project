package config

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

type RedisConfig struct {
	Host     string
	Port     string
	Password string
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
