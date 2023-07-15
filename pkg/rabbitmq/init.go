package rabbitmq

import (
	"log"

	"github.com/joho/godotenv"
)

var (
	mqAddr string
	queName string
)

func init() {
	mqMap, err := godotenv.Read()
	if err != nil {
		panic(err)
	}

	mqAddr = mqMap["MQ_ADDR"]
	queName = mqMap["MQ_QUE_NAME"]
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s\n", msg, err.Error())
	}
}