package main

import (
	"github.com/kamiazya/memomail/adapter/mailer"
	"github.com/spf13/viper"
)

type config struct {
	Mailer *mailer.Config
}

func getConfig() (*config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/memomail/")
	viper.AddConfigPath("$HOME/.memomail")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()
	c := new(config)
	err := viper.Unmarshal(&c)
	return c, err
}
