package models

// SaveText модель для сохранения текста
type SaveText struct {
	Text string `json:"text"`
	Meta string `json:"meta"`
	Name string `json:"name_pair"`
}
