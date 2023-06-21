package config

import "os"

// env Структура для хранения переменных среды
type env struct {
	Port    string
	Host    string
	Address string
	ApiURL  string
}

// Env глобальная переменная для доступа к переменным среды
var Env env

func InitEnv() {
	Env = env{
		Port:    os.Getenv("PORT"),
		Host:    os.Getenv("HOST"),
		Address: os.Getenv("ADDRESS"),
		ApiURL:  os.Getenv("API_URL"),
	}
}
