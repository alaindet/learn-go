package main

import "github.com/spf13/viper"

type config struct {
	PORT                   string
	VERSION                string
	APP_NAME               string
	DATABASE_ROOT_PASSWORD string
	DATABASE_NAME          string
	DATABASE_USER          string
	DATABASE_PASSWORD      string
}

func loadConfig() config {
	viper.SetConfigFile(".env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		panic("Cannot load config file")
	}

	var c config
	viper.Unmarshal(&c)

	return c
}
