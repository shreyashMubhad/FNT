package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var (
	conf         *viper.Viper
	UccWhitelist = make(map[string]bool)
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	LogDir     string
	SSLMode    string
}

func Load(env string, configPaths ...string) {
	var err error
	conf = viper.New()
	conf.SetConfigType("yaml")
	conf.SetConfigName(env)
	conf.AddConfigPath("../config")
	conf.AddConfigPath("config/")
	conf.AddConfigPath("../../app/")
	conf.AddConfigPath(".")
	if len(configPaths) != 0 {
		for _, path := range configPaths {
			conf.AddConfigPath(path)
		}
	}
	err = conf.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file", err)
		return
	}

}

func LoadConfigFromFile() (*Config, error) {
	viper.AutomaticEnv()
	viper.SetConfigFile("../pkg/config/conf.env")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if viper.ConfigFileUsed() == "" {
		return nil, fmt.Errorf("configuration file not found")
	}

	return &Config{
		DBHost:     viper.GetString("DB_HOST"),
		DBPort:     viper.GetString("DB_PORT"),
		DBUser:     viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBName:     viper.GetString("DB_NAME"),
		LogDir:     viper.GetString("LOG_DIR"),
		SSLMode:    viper.GetString("SSLMode"),
	}, nil
}

func GetConfig() *viper.Viper {
	return conf
}
