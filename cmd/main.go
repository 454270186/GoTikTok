package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

var (
	host string
	port string
)

func init() {
	webEnv, err := godotenv.Read("../.env")
	if err != nil {
		panic("fail to read env: " + err.Error())
	}

	host = webEnv["HOST"]
	port = webEnv["PORT"]
}

func main() {
	router := NewRouter()

	ipAddr := fmt.Sprintf("%s:%s", host, port)

	router.Run(ipAddr)
}