package main

import (
	log "github.com/sirupsen/logrus"
	config "main/config"
)

func main() {
	log.Info("Hello, world")

	var conf config.ApplicationConfig
	conf.LoadConfig("config.yml")

	log.Info(conf.Amqp.Url)
}
