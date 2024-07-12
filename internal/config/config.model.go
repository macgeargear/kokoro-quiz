package config

type AppConfig struct {
	Port        string
	Env         string
	ServiceName string
}

type DbConfig struct {
	Url string
}

type Config struct {
	App AppConfig
	Db  DbConfig
}
