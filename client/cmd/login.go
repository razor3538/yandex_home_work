package cmd

import (
	"client/internal/services"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Авторизация",
	Long:  `Авторизация в системе по переданным login и password`,
	Run: func(cmd *cobra.Command, args []string) {
		var authService = services.NewAuthService()

		login, _ := cmd.Flags().GetString("login")
		password, _ := cmd.Flags().GetString("password")

		token, statusCode, err := authService.Login(login, password)
		if err != nil {
			log.Println(err.Error())
		}

		if statusCode == 200 {
			fmt.Println("Авторизация прошла успешно")

			f, err := os.Create("cred.txt")

			if err != nil {
				log.Println(err)
			}

			defer f.Close()

			_, err2 := f.WriteString(token)

			if err2 != nil {
				log.Println(err2)
			}
		} else if statusCode == 401 {
			fmt.Println("Не верный логин или пароль")
		}
	},
}

func init() {
	loginCmd.Flags().StringP("login", "l", "", "Почта для авторизации")
	loginCmd.Flags().StringP("password", "p", "", "Пароль для авторизации")

	rootCmd.AddCommand(loginCmd)
}
