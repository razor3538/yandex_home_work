package cmd

import (
	"client/internal/services"
	"fmt"

	"github.com/spf13/cobra"
)

// deletePasswordCmd represents the deletePassword command
var deletePasswordCmd = &cobra.Command{
	Use:   "deletePassword",
	Short: "Удаление пары логин + пароль",
	Long:  `Удаление пары логин + пароль по переданному имени по пользователю`,
	Run: func(cmd *cobra.Command, args []string) {
		var passService = services.NewPassService()

		name, _ := cmd.Flags().GetString("name")

		result, statusCode, err := passService.Delete(name)
		if err != nil {
			println(err.Error())
		}

		if statusCode == 200 {
			println(result)
		} else if statusCode == 400 {
			fmt.Println("Что то пошло не так. Проверьте правильность ввода данных")
		} else if statusCode == 401 {
			fmt.Println("Ошибка авторизации. Попробуйте войти еще раз")
		}
	},
}

func init() {
	deletePasswordCmd.Flags().StringP("name", "n", "", "Название пары для удаления логина и пароля")

	rootCmd.AddCommand(deletePasswordCmd)
}
