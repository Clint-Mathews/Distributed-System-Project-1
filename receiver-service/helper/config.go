package helper

import (
	"log"

	"github.com/spf13/viper"
)

func readConfig(key string) string {
	viper.SetConfigFile("../../../.env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error readig configurations. Error: %+v", err)
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	return value
}
