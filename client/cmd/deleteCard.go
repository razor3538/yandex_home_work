/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"client/internal/services"
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCardCmd represents the deleteCard command
var deleteCardCmd = &cobra.Command{
	Use:   "deleteCard",
	Short: "Удаление данных карты",
	Long:  `Удаление данных карты по переданному имени по пользователю`,

	Run: func(cmd *cobra.Command, args []string) {
		var cardService = services.NewCardService()

		name, _ := cmd.Flags().GetString("name")

		result, statusCode, err := cardService.Delete(name)
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
	deleteCardCmd.Flags().StringP("name", "n", "", "Название пары для удаления логина и пароля")

	rootCmd.AddCommand(deleteCardCmd)
}
