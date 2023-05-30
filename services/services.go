package services

import "github.com/spf13/viper"

func Init() error {
	oAuthGoogleCfg.ClientID = viper.GetString("oAuthStateSetting.google.clientID")
	oAuthGoogleCfg.ClientSecret = viper.GetString("oAuthStateSetting.google.clientSecret")
	oAuthGoogleState = viper.GetString("oAuthGoogleState")

	return nil
}
