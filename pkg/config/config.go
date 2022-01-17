package config

import (
	"github.com/spf13/viper"
	"log"
)

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigType("json")
	viper.SetConfigName("app.config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}
}