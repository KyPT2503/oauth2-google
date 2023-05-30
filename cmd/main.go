package main

import (
	"log"
	"oauth2-google/configs"
	"oauth2-google/services"
	"oauth2-google/util/logger"
)

func init() {
	if err := logger.Init(); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	logger.Log.Info("initializing config ...")
	if err := configs.Init(); err != nil {
		logger.Log.Error(err.Error())
	}

	logger.Log.Info("initializing services ...")
	if err := services.Init(); err != nil {
		logger.Log.Error(err.Error())
	}

}
