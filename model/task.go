package model

type Task struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}
