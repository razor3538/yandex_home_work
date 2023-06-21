package main

import (
	"client/cmd"
	config "client/init"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")

	config.InitEnv()

	cmd.Execute()
}
