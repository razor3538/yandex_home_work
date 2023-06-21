package main

import (
	"github.com/joho/godotenv"
	"server/config"
	migrate "server/init/db"
	"server/internal/app/routes"
)

func main() {
	_ = godotenv.Load(".env")

	config.InitEnv()
	config.InitDB()
	migrate.Migrate()

	r := routes.SetupRouter()

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
