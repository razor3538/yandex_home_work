package cmd

import (
	"client/internal/services"
	"fmt"

	"github.com/spf13/cobra"
)

// getPasswordCmd represents the getPassword command
var getPasswordCmd = &cobra.Command{
	Use:   "getPassword",
	Short: "Узнать пароль",
	Long:  `Возвращает пару логин + пароль по переданному имени пары`,
	Run: func(cmd *cobra.Command, args []string) {
		var passService = services.NewPassService()

		name, _ := cmd.Flags().GetString("name")

		result, statusCode, err := passService.Get(name)
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
	getPasswordCmd.Flags().StringP("name", "n", "", "Название пары для получения логина и пароля")

	rootCmd.AddCommand(getPasswordCmd)
}
