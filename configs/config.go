package configs

import (
	"github.com/spf13/viper"
)

func Init() error {
	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}
