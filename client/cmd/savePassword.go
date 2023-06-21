package cmd

import (
	"client/internal/services"
	"fmt"

	"github.com/spf13/cobra"
)

// savePasswordCmd represents the savePassword command
var savePasswordCmd = &cobra.Command{
	Use:   "savePassword",
	Short: "Сохранение пары логин пароль",
	Long:  `Сохранение пары логин + пароль по имени`,
	Run: func(cmd *cobra.Command, args []string) {
		var passService = services.NewPassService()

		login, _ := cmd.Flags().GetString("login")
		password, _ := cmd.Flags().GetString("password")
		meta, _ := cmd.Flags().GetString("meta")
		name, _ := cmd.Flags().GetString("name")

		statusCode, err := passService.Save(login, password, name, meta)
		if err != nil {
			println(err.Error())
		}

		if statusCode == 201 {
			fmt.Println("Сохранение пары логин - пароль прошло успешно")
		} else if statusCode == 400 {
			fmt.Println("Что то пошло не так. Проверьте правильность ввода данных")
		} else if statusCode == 401 {
			fmt.Println("Ошибка авторизации. Попробуйте войти еще раз")
		}
	},
}

func init() {
	savePasswordCmd.Flags().StringP("name", "n", "", "Название пары")
	savePasswordCmd.Flags().StringP("meta", "m", "", "Мета информация")
	savePasswordCmd.Flags().StringP("login", "l", "", "Логин для сохранения")
	savePasswordCmd.Flags().StringP("password", "p", "", "Пароль для сохранения")

	rootCmd.AddCommand(savePasswordCmd)
}
