package cmd

import (
	"client/internal/services"
	"fmt"

	"github.com/spf13/cobra"
)

// getTextCmd represents the getText command
var getTextCmd = &cobra.Command{
	Use:   "getText",
	Short: "Показать текст",
	Long:  `Возвращает текст по переданному имени`,
	Run: func(cmd *cobra.Command, args []string) {
		var textService = services.NewTextService()

		name, _ := cmd.Flags().GetString("name")

		result, statusCode, err := textService.Get(name)
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
	getTextCmd.Flags().StringP("name", "n", "", "Название текста для его получения")

	rootCmd.AddCommand(getTextCmd)
}
