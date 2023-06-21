package config

import "os"

// env Структура для хранения переменных среды
type env struct {
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
	Port       string
	Host       string
	Address    string
}

// Env глобальная переменная для доступа к переменным среды
var Env env

func InitEnv() {
	Env = env{
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbName:     os.Getenv("DB_NAME"),
		Port:       os.Getenv("PORT"),
		Host:       os.Getenv("HOST"),
		Address:    os.Getenv("ADDRESS"),
	}
}
