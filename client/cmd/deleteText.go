package cmd

import (
	"client/internal/services"
	"fmt"

	"github.com/spf13/cobra"
)

// deleteTextCmd represents the deleteText command
var deleteTextCmd = &cobra.Command{
	Use:   "deleteText",
	Short: "Удаление текста",
	Long:  `Удаление текста по переданному имени по пользователю`,
	Run: func(cmd *cobra.Command, args []string) {
		var textService = services.NewTextService()

		name, _ := cmd.Flags().GetString("name")

		result, statusCode, err := textService.Delete(name)
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
	deleteTextCmd.Flags().StringP("name", "n", "", "Название текста для его удаления")

	rootCmd.AddCommand(deleteTextCmd)
}
