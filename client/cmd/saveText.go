package cmd

import (
	"client/internal/services"
	"fmt"

	"github.com/spf13/cobra"
)

// saveTextCmd represents the saveText command
var saveTextCmd = &cobra.Command{
	Use:   "saveText",
	Short: "Сохранение текста",
	Long:  `Сохранение текста по имени`,
	Run: func(cmd *cobra.Command, args []string) {
		var textService = services.NewTextService()

		text, _ := cmd.Flags().GetString("text")
		meta, _ := cmd.Flags().GetString("meta")
		name, _ := cmd.Flags().GetString("name")

		statusCode, err := textService.Save(text, name, meta)
		if err != nil {
			println(err.Error())
		}

		if statusCode == 201 {
			fmt.Println("Сохранение текста прошло успешно")
		} else if statusCode == 400 {
			fmt.Println("Что то пошло не так. Проверьте правильность ввода данных")
		} else if statusCode == 401 {
			fmt.Println("Ошибка авторизации. Попробуйте войти еще раз")
		}
	},
}

func init() {
	saveTextCmd.Flags().StringP("name", "n", "", "Название пары")
	saveTextCmd.Flags().StringP("meta", "m", "", "Мета информация")
	saveTextCmd.Flags().StringP("text", "t", "", "Текст для сохранения")

	rootCmd.AddCommand(saveTextCmd)
}
