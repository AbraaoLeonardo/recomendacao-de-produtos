package utils

import (
	"log"

	"gopkg.in/ini.v1"
)

type Config struct {
	DbUser     string
	DbPassword string
	DbName     string
	DbHost     string
	DbPort     string
	ServerPort string
}

func LoadConfigFromFile(path string) *Config {
	cfgFile, err := ini.Load(path)
	if err != nil {
		log.Fatalf("Falha ao carregar o arquivo de configuração em %s: %v", path, err)
	}

	cfg := &Config{}
	dbSection := cfgFile.Section("database")
	cfg.DbUser = dbSection.Key("DB_USER").String()
	cfg.DbPassword = dbSection.Key("DB_PASSWORD").String()
	cfg.DbName = dbSection.Key("DB_NAME").String()
	cfg.DbHost = dbSection.Key("DB_HOST").String()

	cfg.DbPort = dbSection.Key("DB_PORT").MustString("5432")

	serverSection := cfgFile.Section("server")
	cfg.ServerPort = serverSection.Key("SERVER_PORT").MustString("8080")

	if cfg.DbUser == "" || cfg.DbPassword == "" {
		log.Fatal("As configurações de usuário e senha do banco de dados não foram encontradas no arquivo.")
	}

	return cfg
}
