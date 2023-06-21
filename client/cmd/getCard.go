package cmd

import (
	"client/internal/services"
	"fmt"

	"github.com/spf13/cobra"
)

// getCardCmd represents the getCard command
var getCardCmd = &cobra.Command{
	Use:   "getCard",
	Short: "Показать данные карты",
	Long:  `Возвращает данные карты по переданному имени`,
	Run: func(cmd *cobra.Command, args []string) {
		var cardService = services.NewCardService()

		name, _ := cmd.Flags().GetString("name")

		result, statusCode, err := cardService.Get(name)
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
	getCardCmd.Flags().StringP("name", "n", "", "Название карты для его получения")

	rootCmd.AddCommand(getCardCmd)
}
