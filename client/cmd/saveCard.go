package cmd

import (
	"client/internal/services"
	"fmt"

	"github.com/spf13/cobra"
)

// saveCardCmd represents the saveCard command
var saveCardCmd = &cobra.Command{
	Use:   "saveCard",
	Short: "Сохранение данных карты",
	Long:  `Сохранение данных карты по имени`,
	Run: func(cmd *cobra.Command, args []string) {
		var cardService = services.NewCardService()

		number, _ := cmd.Flags().GetString("number")
		cvs, _ := cmd.Flags().GetString("cvs")
		dateEnd, _ := cmd.Flags().GetString("date_end")
		bank, _ := cmd.Flags().GetString("bank")
		meta, _ := cmd.Flags().GetString("meta")
		name, _ := cmd.Flags().GetString("name")

		statusCode, err := cardService.Save(number, cvs, dateEnd, bank, name, meta)
		if err != nil {
			println(err.Error())
		}

		if statusCode == 201 {
			fmt.Println("Сохранение данных карты прошло успешно")
		} else if statusCode == 400 {
			fmt.Println("Что то пошло не так. Проверьте правильность ввода данных")
		} else if statusCode == 401 {
			fmt.Println("Ошибка авторизации. Попробуйте войти еще раз")
		}
	},
}

func init() {
	saveCardCmd.Flags().StringP("name", "n", "", "Название пары")
	saveCardCmd.Flags().StringP("meta", "m", "", "Мета информация")
	saveCardCmd.Flags().StringP("number", "u", "", "Номер карты для сохранения")
	saveCardCmd.Flags().StringP("cvs", "c", "", "cvs для сохранения")
	saveCardCmd.Flags().StringP("date_end", "d", "", "Дата истечения годности для сохранения")
	saveCardCmd.Flags().StringP("bank", "b", "", "Название банка для сохранения")

	rootCmd.AddCommand(saveCardCmd)
}
