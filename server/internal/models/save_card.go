package models

// SaveCard модель для сохранения данных карты
type SaveCard struct {
	Name    string `json:"name_pair"`
	Number  string `json:"number"`
	DateEnd string `json:"date_end"`
	CVS     string `json:"cvs"`
	Bank    string `json:"bank"`
	Meta    string `json:"meta"`
}
