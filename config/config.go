package config

import (
	"log"

	"github.com/spf13/viper"
)

type Settings struct {
	URL string
}

var Config Settings

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Read config error: %v", err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		log.Fatalf("Unmarshal error: %v", err)
	}

}
