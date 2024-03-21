package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	DbUser     string
	DbPassword string
	DbName     string
	DbHost     string
	DbPort     string
	ListenHost string
	ListenPort string
}

func LoadConfig() (*Config, error) {
	var conf = &Config{}
	viper.SetConfigFile("../config.yaml")

	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.SetEnvPrefix("SD")
	v.AutomaticEnv()

	v.SetDefault("DbUser", "postgres")
	v.SetDefault("DbPassword", "103159")
	v.SetDefault("DbName", "myDB")
	v.SetDefault("DbPort", "5432")
	v.SetDefault("DbHost", "127.0.0.1")
	v.SetDefault("ListenHost", "127.0.0.1")
	v.SetDefault("ListenPort", "8088")

	err := v.ReadInConfig() // Find and read the config file

	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
		return nil, err
	}

	conf.DbUser = strings.TrimSpace(v.GetString("DbUser"))
	conf.DbPassword = strings.TrimSpace(v.GetString("DbPassword"))
	conf.DbName = strings.TrimSpace(v.GetString("DbName"))
	conf.DbHost = strings.TrimSpace(v.GetString("DbHost"))
	conf.DbPort = strings.TrimSpace(v.GetString("DbPort"))
	conf.ListenHost = strings.TrimSpace(v.GetString("ListenHost"))
	conf.ListenPort = strings.TrimSpace(v.GetString("ListenPort"))

	return conf, nil
}
