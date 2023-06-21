package cmd

import (
	"client/internal/services"
	"fmt"

	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Регистрирация пользователя",
	Long:  `Регистрация пользователя по переданным login + password`,
	Run: func(cmd *cobra.Command, args []string) {
		var authService = services.NewAuthService()

		login, _ := cmd.Flags().GetString("login")
		password, _ := cmd.Flags().GetString("password")

		statusCode, err := authService.Registration(login, password)
		if err != nil {
			println(err.Error())
		}

		if statusCode == 200 {
			fmt.Println("Регистрация прошла успешно")
		} else if statusCode == 400 {
			fmt.Println("Пользователь уже существует")
		}
	},
}

func init() {
	registerCmd.Flags().StringP("login", "l", "", "Почта для регистрации")
	registerCmd.Flags().StringP("password", "p", "", "Пароль для регистрации")

	rootCmd.AddCommand(registerCmd)
}
